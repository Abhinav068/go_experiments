package v1

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type Log struct {
	Level   string
	Message string
}

func V1() {
	input := []string{
		"INFO user logged in",
		"ERROR database connection failed",
		"WARN cache miss",
		"ERROR database connection failed",
		"WARN cache miss",
		"ERROR database connection failed",
		"WARN cache miss",
	}
	// var logProcessed int
	var inp = make(chan string, len(input))
	var workerChan = make(chan string, 3)
	var op = make(chan Log, len(input))

	for _, val := range input {
		inp <- val
	}
	close(inp)

	wg.Add(len(input))
	go func() {
		for val := range workerChan { // this won't stop draining the workerChan. It won't wait for the task to get completed.
			go task(val, &op)
		}
	}()

	for val := range inp {
		workerChan <- val
	}
	wg.Wait()
	close(workerChan)
	close(op)

	// time.Sleep(time.Duration(time.Second * 1))

	fmt.Println("total logs processed: ", len(op))
}

func task(log string, op *chan Log) {
	defer wg.Done()
	// defer fmt.Println("task done")

	if strings.Contains(log, "INFO") {
		out := Log{
			Level:   "INFO",
			Message: strings.ReplaceAll(log, "INFO", ""),
		}
		*op <- out
	}

	if strings.Contains(log, "ERROR") {
		out := Log{
			Level:   "ERROR",
			Message: strings.ReplaceAll(log, "ERROR", ""),
		}
		*op <- out
	}

	if strings.Contains(log, "WARN") {
		out := Log{
			Level:   "WARN",
			Message: strings.ReplaceAll(log, "WARN", ""),
		}
		*op <- out
	}
}
