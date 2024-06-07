package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a buffered reader to read input from standard input
	reader := bufio.NewReader(os.Stdin)
	
	// Read the number of test cases
	input, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(input))

	results := make([]int, N)

	for i := 0; i < N; i++ {
		// Read the length of the current array
		input, _ := reader.ReadString('\n')
		X, _ := strconv.Atoi(strings.TrimSpace(input))
		
		// Read the array elements
		input, _ = reader.ReadString('\n')
		elements := strings.Fields(input)
		
		sum := 0
		for _, elem := range elements {
			num, _ := strconv.Atoi(elem)
			if num >= 0 {
				sum += num * num
			}
		}
		results[i] = sum
	}

	// Print all results
	for _, result := range results {
		fmt.Println(result)
	}
}
