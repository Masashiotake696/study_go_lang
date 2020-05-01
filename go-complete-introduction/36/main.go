package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	m.Lock()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("unlock 1")
		m.Unlock()
	}()
	m.Lock()
	m.Unlock()
	fmt.Println("unlock 2")
}
