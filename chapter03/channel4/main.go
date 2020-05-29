package main

import "fmt"

func main() {
	valueStream := make(chan interface{})
	close(valueStream)
	v, ok := <-valueStream
	fmt.Printf("%v: %v", ok, v)
}
