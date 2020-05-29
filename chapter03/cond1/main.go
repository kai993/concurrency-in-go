// Condは条件が起きるのを待っている間、
// ロックをずっと保持しているように見えるが実際は違う
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	printQueue := func() {
		fmt.Printf("%#v\n", queue)
	}

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		// Locker.Lock()でクリティカルセッションに入る
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		printQueue()
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
