package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // 1つのゴルーチンが起動したことを表す
	go func() {
		defer wg.Done() // ゴルーチンのクロージャーが終了する前にWaitGroupに終了することを確実に伝える
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // 全てのゴルーチンが終了したと伝えるまでメインゴルーチンをロックする
	fmt.Println("All goroutine complete.")
}
