package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("Image")
	Video = fakeSearch("video")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func GoogleV2(query string) []Result {
	ch := make(chan Result)
	go func() {
		ch <- Web(query)
	}()
	go func() {
		ch <- Video(query)
	}()
	go func() {
		ch <- Image(query)
	}()
	res := make([]Result, 0)

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		// add res timeout control
		select {
		case result := <-ch:
			res = append(res, result)
		case <-timeout:
			fmt.Println("timeout, 80ms exit")
			return res
		}
	}
	return res
}

// add timeout control
// code source:
// https://talks.golang.org/2012/concurrency.slide#47
func main() {
	rand.Seed(time.Now().UnixNano())
	sTime := time.Now()
	results := GoogleV2("golang")
	fmt.Println(results)
	useTime := time.Since(sTime)
	fmt.Printf("usetime is %s\n", useTime.String())
}
