package main

import (
	"bufio"
	"fmt"
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

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run permutate.go <file1> <file2>")
		os.Exit(1)
	}

	filename1 := os.Args[1]
	file1, err := os.Open(filename1)
	if err != nil {
		fmt.Println("Error: Failed to open the file1.")
		os.Exit(1)
	}
	defer file1.Close()

	filename2 := os.Args[2]
	file2, err := os.Open(filename2)
	if err != nil {
		fmt.Println("Error: Failed to open the file2.")
		os.Exit(1)
	}
	defer file2.Close()

	permutationList := make([]string, 0)
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		permutationList = append(permutationList, scanner.Text())
	}

	workers := 16
	domains := make(chan string, workers)
	output := make(chan string, workers)
	wg := &sync.WaitGroup{}

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

	scanner = bufio.NewScanner(file2)

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
