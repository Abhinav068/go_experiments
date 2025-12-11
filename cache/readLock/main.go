package main

import (
	"fmt"
	"sync"
)

var a = 0
var mtx = sync.RWMutex{}
var wg = sync.WaitGroup{}

func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go count()
	}
	for range "4" {
		fmt.Println("4 times")
	}
	wg.Wait()
	fmt.Print("final_value", a)
}

func count() {
	defer wg.Done()
	defer mtx.Unlock()
	mtx.Lock()
	a = a + 1
}

func readVal() {

}
