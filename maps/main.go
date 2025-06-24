package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func countrunes() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// To interact with this countrunes Go code, you'll need to:

// Save the code to a file with a .go extension, for example charcount.go
// Run the program by opening a terminal and typing:
//? go run charcount.go

// Provide input in one of these ways:
// 1. Type text directly and press Ctrl+D (Unix/Linux/Mac) or Ctrl+Z followed by Enter (Windows) to signal the end of input
// 2. Pipe a file into the program:
//? 		cat myfile.txt | go run charcount.go

// 3. Redirect a file as input:
//? 		go run charcount.go < myfile.txt

func main() {
	// Adding edges to build a simple directed graph
	addEdge("A", "B")
	addEdge("A", "C")
	addEdge("B", "C")
	addEdge("B", "D")
	addEdge("C", "D")
	addEdge("D", "A") // Creates a cycle

	// Test edge existence
	fmt.Println("Edge from A to B:", hasEdge("A", "B")) // true
	fmt.Println("Edge from A to D:", hasEdge("A", "D")) // false
	fmt.Println("Edge from C to D:", hasEdge("C", "D")) // true
	fmt.Println("Edge from D to C:", hasEdge("D", "C")) // false

	// Print the entire graph structure
	fmt.Println("\nGraph structure:")
	for node, edges := range graph {
		fmt.Printf("Node %s connects to: ", node)
		for target := range edges {
			fmt.Printf("%s ", target)
		}
		fmt.Println()
	}
}
