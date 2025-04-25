package main

import "fmt"

func main() {
    ch := make(chan int)
    
    go func() {
        for i := 1; i <= 3; i++ {
            ch <- i
        }
		close(ch) // will deadlock if commented out
    }()
    
    // range automatically stops when channel is closed
    for value := range ch {
		fmt.Println("inside range: ")
        fmt.Println("Received:", value)
    }
    fmt.Println("Channel closed!")
}