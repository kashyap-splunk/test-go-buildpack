package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello world, go-buildpack")
		time.Sleep(600 * time.Second)
	}
}
