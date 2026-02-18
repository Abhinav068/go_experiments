package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type Log struct {
	Level   string
	Message string
	Success bool
}

func main() {
	success := 0
	falied := 0
	input := []string{
		"INFO 1 user logged in",
		"ERROR 2 database connection failed",
		"WARN 3 cache miss",
		"ERROR 4 database connection failed",
		"WARN 5 cache miss",
		"ERROR 6 database connection failed",
		"WARN 7 cache miss",
		"hehehe", // invalid case
	}
	// var logProcessed int
	var op = make(chan Log, len(input))
	var pool = make(chan string, len(input))

	for _, val := range input {
		pool <- val

	}
	close(pool)
	wg.Add(len(input))
	for range 3 {
		go func(inp chan string, oup chan Log) {
			for val := range inp {
				task(val, op)
			}

		}(pool, op)
		// time.Sleep(time.Millisecond)
	}
	wg.Wait()
	close(op)
	for val := range op {
		if val.Success {
			success++
		} else {
			falied++
		}
	}
	// time.Sleep(time.Duration(time.Second * 1))

	fmt.Printf("total logs processed: %d, failed: %d\n", success, falied)
}

func task(log string, op chan Log) {
	// fmt.Println("working on: ", log)
	defer wg.Done()
	var out Log
	if strings.Contains(log, "INFO") {
		out = Log{
			Level:   "INFO",
			Message: strings.ReplaceAll(log, "INFO", ""),
			Success: true,
		}
		op <- out
	} else if strings.Contains(log, "ERROR") {
		out = Log{
			Level:   "ERROR",
			Message: strings.ReplaceAll(log, "ERROR", ""),
			Success: true,
		}
		op <- out
	} else if strings.Contains(log, "WARN") {
		out = Log{
			Level:   "WARN",
			Message: strings.ReplaceAll(log, "WARN", ""),
			Success: true,
		}
		op <- out
	} else {
		out = Log{
			Level:   "",
			Message: "",
			Success: false,
		}
		op <- out
	}
	// fmt.Printf("for: %s, result: %v \n\n", out.Message, out.Success)
	// fmt.Printf("worker length: %v", len(worker))
	// time.Sleep(time.Duration(time.Second * 1))
}
