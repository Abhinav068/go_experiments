package main

import (
	"fmt"
	"reflect"
)

// Example 1: Basic Reflection to get Value and Type
type Person struct {
	Name string
	Age  int
}

func reflectExample1() {
	p := Person{Name: "John", Age: 25}

	// Get Value and Type
	val := reflect.ValueOf(p) // Creates Value instance
	typ := val.Type()         // Gets Type

	fmt.Printf("Number of fields: %d and type: %s\n", val.NumField(), typ) // Output: 2

	// Iterate through fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)     // this gives ith field's value of the value (struct in this case)
		typeField := typ.Field(i) // this gives ith field's name and its type and some other values.
		// fmt.Printf("Field: %s, Value: %v\n", typeField.Name, field.Interface())
		fmt.Printf("field: %v,| typeField: %v\n", field, typeField)
	}
}

// Example 2: Modifying struct fields using reflection
func reflectExample2() {
	p := &Person{Name: "John", Age: 25}

	// Get pointer's underlying value
	val := reflect.ValueOf(p).Elem()

	// Modify fields
	nameField := val.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Alice")
	}

	ageField := val.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(30)
	}

	fmt.Printf("Modified person: %+v\n", p) // Output: {Name:Alice Age:30}
}

func main() {
	reflectExample1()
	// reflectExample2()
}



// {Age  int  16 [1] false}
//      ↑   ↑   ↑   ↑    ↑
//      |   |   |   |    not an embedded field
//      |   |   |   second field in struct
//      |   |   starts at byte offset 16
//      |   field type is int
//      no struct tags