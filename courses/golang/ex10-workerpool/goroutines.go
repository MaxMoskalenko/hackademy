package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func Run(poolSize int) {
	jobs := make(chan string, poolSize)
	var wg sync.WaitGroup
	id := 1
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		jobs <- scan.Text()
		if id <= poolSize {
			wg.Add(1)
			go worker(id, jobs, &wg)
			id++
		}
	}
	close(jobs)
	wg.Wait()
}

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	if len(jobs) > 0 {
		fmt.Printf("worker:%d spawning\n", id)
		for j := range jobs {
			fmt.Printf("worker:%d sleep:%s\n", id, j)
			t, err := time.ParseDuration(j + "s")
			handleError(err)
			time.Sleep(t)
		}
		fmt.Printf("worker:%d stopping\n", id)
	}
	wg.Done()
}

func handleError(err error) {
	if err == nil {
		return
	}
	fmt.Println("ERROR: ", err.Error())
}
