package main

import "fmt"
// Slices of interface{}
func main() {
    // Example: Slice of interface{} containing itself
    selfSlice := []interface{}{1, "hello", nil, 42.0}
    selfSlice[2] = selfSlice // The third element now references the entire slice
    
    fmt.Println("Self-referential slice:", selfSlice)
    
    // We can access elements normally
    fmt.Println("First element:", selfSlice[0])
    fmt.Println("Second element:", selfSlice[1])
    
    // We can access the self-reference
    nestedSlice := selfSlice[2].([]interface{})
    fmt.Println("First element of nested reference:", nestedSlice[0])
}

// Slices of other reference types (like maps with interface{} values)

func main2() {
    // Example: Slice of maps that contains itself through a map value
    mapSlice := []map[string]interface{}{
        {"name": "first", "data": 100},
        {"name": "second", "data": 200},
    }
    
    // Create the self-reference in the second map's "ref" field
    mapSlice[1]["ref"] = mapSlice
    
    fmt.Println("Map slice with self-reference:", mapSlice)
    
    // We can access the self-reference
    secondMap := mapSlice[1]
    selfRef := secondMap["ref"].([]map[string]interface{})
    fmt.Println("Referenced name:", selfRef[0]["name"])
    
    // Another example with nested maps
    nestedMaps := map[string]interface{}{
        "outer": map[string]interface{}{
            "inner": "value",
        },
    }
    
    // Make the inner map reference the outer map
    innerMap := nestedMaps["outer"].(map[string]interface{})
    innerMap["parent"] = nestedMaps
    
    fmt.Println("Nested maps with circular reference:", nestedMaps)
}



//..............................................................................................................//
// Structs with pointer fields that can reference themselves

// Example 1: Simple linked list with self-reference
type Node struct {
    Value int
    Next  *Node
}

// Example 2: Tree structure with self-references
type TreeNode struct {
    Value    string
    Children []*TreeNode
    Parent   *TreeNode
}

// Example 3: Complex structure with multiple reference types
type Department struct {
    Name      string
    Employees []*Employee
}

type Employee struct {
    Name       string
    Department *Department
    Manager    *Employee
    Team       []*Employee
}

func main3() {
    // Example 1: Create a circular linked list
    head := &Node{Value: 1}
    second := &Node{Value: 2}
    third := &Node{Value: 3}
    
    head.Next = second
    second.Next = third
    third.Next = head // Creates a cycle
    
    fmt.Println("Circular linked list:", head, head.Next, head.Next.Next, head.Next.Next.Next)
    
    // Example 2: Create a tree with parent references
    root := &TreeNode{Value: "Root"}
    child1 := &TreeNode{Value: "Child 1", Parent: root}
    child2 := &TreeNode{Value: "Child 2", Parent: root}
    
    root.Children = []*TreeNode{child1, child2}
    
    fmt.Println("Tree with parent references:", root)
    fmt.Println("Child's parent value:", root.Children[0].Parent.Value)
    
    // Example 3: Create a complex organizational structure
    engineering := &Department{Name: "Engineering"}
    
    cto := &Employee{Name: "Alice", Department: engineering}
    dev1 := &Employee{Name: "Bob", Department: engineering, Manager: cto}
    dev2 := &Employee{Name: "Charlie", Department: engineering, Manager: cto}
    
    cto.Team = []*Employee{dev1, dev2}
    engineering.Employees = []*Employee{cto, dev1, dev2}
    
    fmt.Println("Department:", engineering.Name)
    fmt.Println("CTO manages:", engineering.Employees[0].Team[0].Name)
    fmt.Println("Dev1's department:", dev1.Department.Name)
    fmt.Println("Dev1's manager:", dev1.Manager.Name)
}