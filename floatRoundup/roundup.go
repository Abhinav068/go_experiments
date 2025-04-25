package main

import (
	"fmt"
	"math"
)

const (
	deadbeef = 0xdeadbeef        // untyped int with value 3735928559
	a        = uint32(deadbeef)  // uint32 with value  3735928559
	b        = float32(deadbeef) // float32 with value 3735928576 (rounded up)
	c        = float64(deadbeef) // float64 with value 3735928559 (exact)
	// d        = int32(deadbeef)   // compile error: constant overflows int32
	// e        = float64(1e309)    // compile error: constant overflows float64
	// f        = uint(-1)          // compile error: constant underflows uint
)

func main() {
	var f float64 = 212

	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0"; 5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	fmt.Printf("32bits %.f\n", float32(0xdeadbeef))           // 3735928832
	fmt.Printf("%x\n", math.Float32bits(float32(0xdeadbeef))) // 4eb1df00

	fmt.Printf("const b: %.f", b)
}
