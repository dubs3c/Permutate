package main

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
)

func Permutate(out chan<- string, domain *string, permutationList *[]string) {
	d := *domain
	for _, v := range *permutationList {
		str := v + "." + string(d)
		out <- str
	}
}

func main() {

	permutationList := []string{"mail", "vpn", "admin", "www", "remote"}
	workers := 16
	domains := make(chan string, workers)
	output := make(chan string, workers)
	wg := &sync.WaitGroup{}

	file, err := os.OpenFile("data.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatal("Could not open domain file, error: ", err)
	}

	defer file.Close()

	startTime := time.Now()

	// Responsible for outputting permutated domains
	// One output worker per consumer worker
	go func() {
		for i := 0; i < workers; i++ {
			go func() {
				for d := range output {
					os.Stdout.Write([]byte(d + "\n"))
				}
			}()
		}
	}()

	// consumer workers
	wg.Add(workers)
	go func() {
		for i := 0; i < workers; i++ {
			go func(i int) {
				defer wg.Done()

				for d := range domains {
					Permutate(output, &d, &permutationList)
				}
			}(i)
		}
	}()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Producer
	for scanner.Scan() {
		domains <- scanner.Text()
	}

	close(domains)
	wg.Wait()
	close(output)

	duration := time.Since(startTime)

	log.Printf("Time elapsed: %s\n", duration)
}
