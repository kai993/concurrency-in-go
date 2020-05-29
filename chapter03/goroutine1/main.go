package main

import "fmt"

func main() {
	// 関数
	go sayHello()
	// 他の処理...

	// 無名関数
	go func() {
		fmt.Println("hello2")
	}()
	// 他の処理...

	// 変数
	sayHello := func() {
		fmt.Println("hello3")
	}
	// 他の処理...
	go sayHello()
}

func sayHello() {
	fmt.Println("hello1")
}
