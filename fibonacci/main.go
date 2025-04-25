package main

import "fmt"

func fibonacci() func() int {
	i := 0
	j := 1
	return func() int {
		ans := i
		i = j
		j = j + ans
		return ans
	}
}

func main() {
	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
