package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Println("worker-", id, "started job")
		// 执行任务
		time.Sleep(time.Second)
		log.Println("worker-", id, "finished job")

		// collect results
		results <- j * 2
	}
}

func workerEfficient(id int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			log.Println("worker-", id, "started job")
			// 执行任务
			time.Sleep(time.Second)
			log.Println("worker-", id, "finished job")

			// collect results
			results <- job * 2
		}(j)
	}
	wg.Wait()
}

const NumJobs = 8

func main() {
	jobs := make(chan int, NumJobs)
	results := make(chan int, NumJobs)
	for w := 1; w <= 3; w++ {
		go workerEfficient(w, jobs, results)
	}

	// 2. send the work
	// other goroutine sends the work to the channels

	// in this example, the `main` goroutine sends the work to the channel `jobs`
	for j := 1; j <= NumJobs; j++ {
		jobs <- j
	}
	close(jobs)
	fmt.Println("Closed job")
	for a := 1; a <= NumJobs; a++ {
		<-results
	}
	close(results)

}
