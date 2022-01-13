package main

import (
	"edenfarm-be-jiddyabd/lib"
	"fmt"
	"strings"
)

func main() {
	length := 0
	fmt.Println("Enter the length of the array")
	fmt.Scanln(&length)
	fmt.Println("Enter value")
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
	}

	fmt.Println("Array: " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ", "), "[]"))
	fmt.Println("Result: " + fmt.Sprint(lib.LargestValue(numbers)))
}
