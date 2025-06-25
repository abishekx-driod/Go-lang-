package main

import (
	"fmt"
	"sync"
)

func workers(id int, job <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for jobs := range job {
		fmt.Printf("Worker %d processing job %d\n", id, jobs)
		res := jobs * jobs
		result <- res
	}
}

func main() {
	job := make(chan int, 10)
	res := make(chan int, 10) 
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workers(i, job, res, &wg)
	}
	for i := 1; i <= 10; i++ {
		job <- i
	}
	close(job) 

	
	wg.Wait()

	for ress := range res {
		fmt.Println("Result received:", ress)
	}
}
