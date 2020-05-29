package main

import "fmt"

func main() {
	stringSteram := make(chan string)
	go func() {
		stringSteram <- "Hello channels!"
	}()
	fmt.Println(<-stringSteram) // Hello channels!
}
