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

func Google(query string) []Result {
	res := make([]Result, 0)
	res = append(res, Web(query))
	res = append(res, Video(query))
	res = append(res, Image(query))
	return res
}

// Serial execution
// code source:
// https://talks.golang.org/2012/concurrency.slide#45
func main() {
	rand.Seed(time.Now().UnixNano())
	sTime := time.Now()
	results := Google("golang")
	fmt.Println(results)
	useTime := time.Since(sTime)
	fmt.Printf("usetime is %s\n", useTime.String())
}
