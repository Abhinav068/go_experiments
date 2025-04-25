package main

import (
    "fmt"
    "sync"
)

func worker(id int, done chan bool, wg *sync.WaitGroup) {
    defer wg.Done()
    <-done  // Wait for signal
    fmt.Printf("Worker %d stopping\n", id)
}

func main1() {
    done := make(chan bool)
    var wg sync.WaitGroup
    
    // Start multiple workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, done, &wg)
    }
    
    // Signal all workers to stop by closing channel
    close(done)
    wg.Wait()
    fmt.Println("All workers stopped")
}