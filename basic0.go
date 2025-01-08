package main

import "fmt"

// Exercise 1: Print numbers from 1 to n using a for loop
// Operation: Basic for loop
func printNumbers(n int) {
	// TODO: Use a for loop to print numbers from 1 to n
}

func testPrintNumbers() {
	fmt.Println("Expected: 1 2 3 4 5")
	fmt.Print("Result:   ")
	printNumbers(5)
	fmt.Println()
}

// Exercise 2: Print even numbers from 1 to n using a for loop with a condition
// Operation: For loop with an if condition
func printEvenNumbers(n int) {
	// TODO: Use a for loop and if condition to print even numbers from 1 to n
}

func testPrintEvenNumbers() {
	fmt.Println("Expected: 2 4 6 8 10")
	fmt.Print("Result:   ")
	printEvenNumbers(10)
	fmt.Println()
}

// Exercise 3: Sum numbers from 1 to n using a for loop
// Operation: Accumulating sum in a for loop
func sumNumbers(n int) int {
	// TODO: Use a for loop to calculate the sum of numbers from 1 to n
	return 0
}

func testSumNumbers() {
	input := 5
	expected := 15
	result := sumNumbers(input)
	if result == expected {
		fmt.Println("testSumNumbers passed")
	} else {
		fmt.Printf("testSumNumbers failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 4: Find the factorial of a number
// Operation: For loop with multiplication
func factorial(n int) int {
	// TODO: Use a for loop to calculate the factorial of n
	return 0
}

func testFactorial() {
	input := 5
	expected := 120
	result := factorial(input)
	if result == expected {
		fmt.Println("testFactorial passed")
	} else {
		fmt.Printf("testFactorial failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 5: Check if a number is prime
// Operation: Nested loop with a break statement
func isPrime(n int) bool {
	// TODO: Use a for loop to check if n is prime
	return false
}

func testIsPrime() {
	input := 7
	expected := true
	result := isPrime(input)
	if result == expected {
		fmt.Println("testIsPrime passed")
	} else {
		fmt.Printf("testIsPrime failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 6: Print elements of a slice using range
// Operation: Range loop over a slice
func printSliceElements(slice []int) {
	// TODO: Use range to print each element in the slice
}

func testPrintSliceElements() {
	fmt.Println("Expected: 1 2 3 4 5")
	fmt.Print("Result:   ")
	printSliceElements([]int{1, 2, 3, 4, 5})
	fmt.Println()
}

// Exercise 7: Count occurrences of a specific element in a slice
// Operation: Range loop with an if condition
func countOccurrences(slice []int, target int) int {
	// TODO: Use range to count occurrences of target in slice
	return 0
}

func testCountOccurrences() {
	input := []int{1, 2, 2, 3, 4, 2}
	target := 2
	expected := 3
	result := countOccurrences(input, target)
	if result == expected {
		fmt.Println("testCountOccurrences passed")
	} else {
		fmt.Printf("testCountOccurrences failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 8: Check if a slice contains a specific element
// Operation: Range loop with a return statement
func sliceContains(slice []int, target int) bool {
	// TODO: Use range to check if target is in slice
	return false
}

func testSliceContains() {
	input := []int{1, 2, 3, 4, 5}
	target := 3
	expected := true
	result := sliceContains(input, target)
	if result == expected {
		fmt.Println("testSliceContains passed")
	} else {
		fmt.Printf("testSliceContains failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 9: Print keys and values of a map using range
// Operation: Range loop over a map
func printMapElements(m map[string]int) {
	// TODO: Use range to print keys and values of the map
}

func testPrintMapElements() {
	fmt.Println("Expected: a:1 b:2 c:3")
	fmt.Print("Result:   ")
	printMapElements(map[string]int{"a": 1, "b": 2, "c": 3})
	fmt.Println()
}

// Exercise 10: Find the length of the longest word in a slice
// Operation: Range loop with string length comparison
func longestWordLength(words []string) int {
	// TODO: Use range to find the length of the longest word in words
	return 0
}

func testLongestWordLength() {
	input := []string{"Go", "Golang", "is", "fun"}
	expected := 6
	result := longestWordLength(input)
	if result == expected {
		fmt.Println("testLongestWordLength passed")
	} else {
		fmt.Printf("testLongestWordLength failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 11: Count the number of positive numbers in a slice
// Operation: Range loop with if condition
func countPositiveNumbers(slice []int) int {
	// TODO: Use range to count positive numbers in the slice
	return 0
}

func testCountPositiveNumbers() {
	input := []int{-1, 2, -3, 4, 5}
	expected := 3
	result := countPositiveNumbers(input)
	if result == expected {
		fmt.Println("testCountPositiveNumbers passed")
	} else {
		fmt.Printf("testCountPositiveNumbers failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 12: Print a pattern of stars for a given number of rows
// Operation: Nested loops
func printStarPattern(rows int) {
	// TODO: Use nested loops to print a pattern of stars
}

func testPrintStarPattern() {
	fmt.Println("Expected:\n*\n**\n***\n****\n*****")
	fmt.Println("Result:")
	printStarPattern(5)
}

// Run all tests
func runTests() {
	testPrintNumbers()
	testPrintEvenNumbers()
	testSumNumbers()
	testFactorial()
	testIsPrime()
	testPrintSliceElements()
	testCountOccurrences()
	testSliceContains()
	testPrintMapElements()
	testLongestWordLength()
	testCountPositiveNumbers()
	testPrintStarPattern()
}

func main() {
	runTests()
}
