package main

import "fmt"

// Define an interface
type Animal interface {
    Speak() string
}

// Implement the interface with a struct
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    var animals []Animal
    animals = append(animals, Dog{})
    animals = append(animals, Cat{})

    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}