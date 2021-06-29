package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2")
	Web3   = fakeSearch("web3")
	Image1 = fakeSearch("Image1")
	Image2 = fakeSearch("Image2")
	Image3 = fakeSearch("Image3")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
	Video3 = fakeSearch("video3")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// how do we avoid discarding result from the slow server
// we duplicates to many instance, and perform parallel request
func first(query string, replicas ...Search) Result {
	ch := make(chan Result)
	for _, search := range replicas {
		go func(s Search) {
			ch <- s(query)
		}(search)
	}
	return <-ch
}

func GoogleV3(query string) []Result {
	ch := make(chan Result)
	go func() {
		ch <- first(query, Web1, Web2, Web3)
	}()
	go func() {
		ch <- first(query, Image1, Image2, Image3)
	}()
	go func() {
		ch <- first(query, Video1, Video2, Video3)
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

// return first result and timeout control
// code source:
// https://talks.golang.org/2012/concurrency.slide#50
func main() {
	rand.Seed(time.Now().UnixNano())
	sTime := time.Now()
	results := GoogleV3("golang")
	fmt.Println(results)
	useTime := time.Since(sTime)
	fmt.Printf("usetime is %s\n", useTime.String())
}
