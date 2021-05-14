package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b) + "\n"
}

func main() {

	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := 0; i < 1000; i++ {
		data := RandStringBytes(6)

		if _, err = f.WriteString(data); err != nil {
			log.Println(err)
		}
	}
	fmt.Println("Done.")
}
