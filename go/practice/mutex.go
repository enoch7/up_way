package main

import (
	"fmt"
	"sync"
	"time"
	"sync/atomic"
)

func main() {
	var num int = 0
	var mutex = &sync.Mutex{}

	var writeCount uint64 = 0
	for i := 0; i < 20; i++ {
		go func() {
			mutex.Lock()		
			num = num + 1
			mutex.Unlock()
			atomic.AddUint64(&writeCount, 1)
			time.Sleep(time.Millisecond)
		}()	
	}

	for i := 0; i < 20; i++ {
		go func() {
			num = num + 1
			// atomic.AddUint64(&writeCount, 1)
			time.Sleep(time.Millisecond)
		}()	
	}

	time.Sleep(time.Second)

	fmt.Println("num:",num)

	fmt.Println("writeCount:",writeCount)
	
}