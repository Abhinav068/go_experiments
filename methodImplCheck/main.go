package main

import "fmt"

// Our base interface - like ResponseWriter
type Logger interface {
	Log(message string)
}

// Imagine this is an optional capability some Logger implementations might have
type Disableable interface {
	DisableLogging()
}

// A basic implementation that only implements Logger
type SimpleLogger struct {
	enabled bool
}

func (s *SimpleLogger) Log(message string) {
	if s.enabled {
		fmt.Println("[Simple]:", message)
	}
}

// An advanced implementation that implements both interfaces
type AdvancedLogger struct {
	enabled bool
}

func (a *AdvancedLogger) Log(message string) {
	if a.enabled {
		fmt.Println("[Advanced]:", message)
	}
}

func (a *AdvancedLogger) DisableLogging() {
	a.enabled = false
	fmt.Println("Logging has been disabled")
}
func main() {
	// Create our two different loggers
	simple := &SimpleLogger{enabled: true}
	advanced := &AdvancedLogger{enabled: true}

	// Both can log
	simple.Log("Hello from simple")
	advanced.Log("Hello from advanced")

	// Now try to disable each logger
	fmt.Println("\nTrying to disable simple logger:")
	tryToDisable(simple)

	fmt.Println("\nTrying to disable advanced logger:")
	tryToDisable(advanced)

	// Let's see if they're still working
	fmt.Println("\nTrying to log after attempted disable:")
	simple.Log("Simple after disable attempt")
	advanced.Log("Advanced after disable attempt")
}

// This function takes any Logger but tries to disable it if possible
func tryToDisable(logger Logger) {
	// Here's our type assertion, similar to your example
	if disableable, ok := logger.(interface{ DisableLogging() }); ok {
		// This branch executes only if the concrete type behind logger
		// implements the DisableLogging method
		fmt.Println("This logger can be disabled!")
		disableable.DisableLogging()
	} else {
		fmt.Println("This logger cannot be disabled")
	}
}
