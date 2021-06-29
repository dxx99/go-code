package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	filepath "path/filepath"
	"sort"
	"sync"
)

const NumDigesters = 20

var wg sync.WaitGroup

type md5HashByte [md5.Size]byte

// walkFiles starts a goroutines to walk the directory tree ad root and send the path
// of each regular file on the string channel. It sends the result of the walk on the
// error channel. If done is closed, walkFiles abandons its work
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	err := make(chan error, 1)

	go func() {
		defer close(paths)

		err <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, err
}

// result is the product of reading and summing a file using Md5
type result struct {
	path string
	sum  md5HashByte
	err  error
}

// digester reads paths name from paths and sends digests of the corresponding files on c until
// either paths or done is closed
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)

		// 防止处理任务阻塞 select-case 模型
		select {
		case c <- result{
			path: path,
			sum:  md5.Sum(data),
			err:  err,
		}:
		case <-done:
			return
		}
	}
}

func Md5All(root string) (map[string]md5HashByte, error) {
	done := make(chan struct{})
	defer close(done)

	// 搜集任务
	paths, err := walkFiles(done, root)

	c := make(chan result)
	wg.Add(NumDigesters)
	for i := 0; i < NumDigesters; i++ {
		go func() {
			defer wg.Done()

			// 将任务放到多个goroutines处理
			digester(done, paths, c)
		}()
	}

	// 这里必须用协程，不然就启动不了
	go func() {
		wg.Wait()
		// end of pipeline.
		close(c)
	}()

	m := make(map[string]md5HashByte)
	for res := range c {
		if res.err != nil {
			return nil, res.err
		}
		m[res.path] = md5HashByte(res.sum)
	}

	// check whether the walk failed
	if e := <-err; e != nil {
		return nil, e
	}
	return m, nil
}

// 遍历指定目录的数据，然后将结果通过channel传递
func main() {
	m, err := Md5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	paths := make([]string, 0)
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x --> %s\n", m[path], path)
	}
}
