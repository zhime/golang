package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup
var lock sync.Mutex

func Hello() {
	lock.Lock()
	fmt.Println("Hello")
	lock.Unlock()
	wait.Done()
}

func main() {
	wait.Add(1)
	go Hello()
	wait.Wait()
}
