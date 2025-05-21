package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// fmt.Println(`Rules :
//  0 : 2					--> 1
//  1 : 2, 4				--> 2
//  2 : 2, 4, 6				--> 3
//  4 : 2, 4, 6, 8			--> 4
//  5 : 2, 4, 6, 8, 10		--> 6
//  6 : 2, 4, 6, 8, 10, 12	--> 8`)

func main() {
	// Set up channel to listen for interrupt (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		output := []int{1, 2, 3, 4, 6, 8}

		for i := 0; i < 6; i++ {
			nums := []string{}
			for j := 1; j <= i+1; j++ {
				nums = append(nums, fmt.Sprintf("%d", j*2))
			}

			// Join the numbers with commas
			numList := strings.Join(nums, ", ")

			// Format the output with padding
			fmt.Printf("%2d : %-24s --> %d\n", i, numList, output[i])
		}
		fmt.Println("Press Ctrl+C to quit.")

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Enter number of terms: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			nTerms, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("⚠️ Abbeee! Sahi integer daalo. Jante nahi ho!")
				fmt.Println("")
				continue
			}

			getDigits(nTerms)
		}

	}()
	<-stop
	fmt.Println("\n👋 Kya Soche Ho!")
}

func getDigits(nTerms int) {

	var input, a, d int = nTerms, 2, 2
	tn := a + input*d
	fmt.Println("Last terms is : ", tn)

	mult := 1
	digits := 0
	n := 2
	for n <= tn {
		if n < int(math.Pow(10, float64(mult))) {
			digits += mult
		} else {
			// else if n >= int(math.Pow(10, float64(mult)))
			mult++
			digits += mult
		}
		n += d
	}
	fmt.Println("Total number of digits: ", digits)
	fmt.Println("")
}
