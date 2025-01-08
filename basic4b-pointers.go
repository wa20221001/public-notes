package main

import "fmt"

// Exercise 1: Efficiently pass a large struct by pointer
// Operation: Pass a large struct by pointer to avoid copying large data
type LargeStruct struct {
	Data [1000]int
}

func modifyLargeStruct(ls *LargeStruct) {
	// TODO: Modify the first element of the Data array to 99
}

func testModifyLargeStruct() {
	ls := LargeStruct{}
	modifyLargeStruct(&ls)
	expected := 99
	if ls.Data[0] == expected {
		fmt.Println("testModifyLargeStruct passed")
	} else {
		fmt.Printf("testModifyLargeStruct failed: expected %v, got %v\n", expected, ls.Data[0])
	}
}

// Exercise 2: Use a function to reset a variable to its zero value using a pointer
// Operation: Reset a variable by passing a pointer
func resetToZero(ptr *int) {
	// TODO: Set the value pointed to by ptr to 0
}

func testResetToZero() {
	x := 25
	resetToZero(&x)
	expected := 0
	if x == expected {
		fmt.Println("testResetToZero passed")
	} else {
		fmt.Printf("testResetToZero failed: expected %v, got %v\n", expected, x)
	}
}

// Exercise 3: Use a function to return the address of a variable
// Operation: Return a pointer from a function
func getPointerToInt(value int) *int {
	// TODO: Return a pointer to the given value
	return nil
}

func testGetPointerToInt() {
	value := 42
	ptr := getPointerToInt(value)
	if ptr != nil && *ptr == value {
		fmt.Println("testGetPointerToInt passed")
	} else {
		fmt.Printf("testGetPointerToInt failed: expected %v, got %v\n", value, ptr)
	}
}

// Exercise 4: Use a function to compare two integers via pointers
// Operation: Compare two integers using pointers
func compareInts(a, b *int) string {
	// TODO: Return "equal" if the values are the same, otherwise "not equal"
	return ""
}

func testCompareInts() {
	x, y := 10, 10
	expected := "equal"
	result := compareInts(&x, &y)
	if result == expected {
		fmt.Println("testCompareInts passed")
	} else {
		fmt.Printf("testCompareInts failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 5: Create a function to increment a value by a given amount using a pointer
// Operation: Increment a value using a pointer
func incrementBy(ptr *int, amount int) {
	// TODO: Increment the value pointed to by ptr by amount
}

func testIncrementBy() {
	x := 5
	incrementBy(&x, 10)
	expected := 15
	if x == expected {
		fmt.Println("testIncrementBy passed")
	} else {
		fmt.Printf("testIncrementBy failed: expected %v, got %v\n", expected, x)
	}
}

// Exercise 6: Pass a pointer to a slice to append an element
// Operation: Modify a slice via a pointer
func appendToSlice(slice *[]int, value int) {
	// TODO: Append value to the slice pointed to by slice
}

func testAppendToSlice() {
	slice := []int{1, 2, 3}
	appendToSlice(&slice, 4)
	expected := []int{1, 2, 3, 4}
	if fmt.Sprintf("%v", slice) == fmt.Sprintf("%v", expected) {
		fmt.Println("testAppendToSlice passed")
	} else {
		fmt.Printf("testAppendToSlice failed: expected %v, got %v\n", expected, slice)
	}
}

// Exercise 7: Use two functions to calculate the square and cube of a number using pointers
// Operation: Use multiple functions and pointers to calculate results
func square(ptr *int) {
	// TODO: Square the value pointed to by ptr
}

func cube(ptr *int) {
	// TODO: Cube the value pointed to by ptr
}

func testSquareAndCube() {
	x := 2
	square(&x)
	expectedSquare := 4
	if x != expectedSquare {
		fmt.Printf("testSquare failed: expected %v, got %v\n", expectedSquare, x)
		return
	}
	cube(&x)
	expectedCube := 64
	if x == expectedCube {
		fmt.Println("testSquareAndCube passed")
	} else {
		fmt.Printf("testCube failed: expected %v, got %v\n", expectedCube, x)
	}
}

// Exercise 8: Return multiple values by using pointers
// Operation: Return two results via pointers
func divideAndRemainder(a, b int, quotient, remainder *int) {
	// TODO: Compute quotient and remainder of a / b and store in the respective pointers
}

func testDivideAndRemainder() {
	a, b := 10, 3
	var quotient, remainder int
	divideAndRemainder(a, b, &quotient, &remainder)
	expectedQuotient, expectedRemainder := 3, 1
	if quotient == expectedQuotient && remainder == expectedRemainder {
		fmt.Println("testDivideAndRemainder passed")
	} else {
		fmt.Printf("testDivideAndRemainder failed: expected quotient=%v, remainder=%v; got quotient=%v, remainder=%v\n",
			expectedQuotient, expectedRemainder, quotient, remainder)
	}
}

// Exercise 9: Create a function to find the minimum and maximum of an array using pointers
// Operation: Use pointers to return multiple results
func minMax(arr []int, min, max *int) {
	// TODO: Find the minimum and maximum values in the array and store them in min and max
}

func testMinMax() {
	arr := []int{5, 1, 9, 3, 7}
	var min, max int
	minMax(arr, &min, &max)
	expectedMin, expectedMax := 1, 9
	if min == expectedMin && max == expectedMax {
		fmt.Println("testMinMax passed")
	} else {
		fmt.Printf("testMinMax failed: expected min=%v, max=%v; got min=%v, max=%v\n",
			expectedMin, expectedMax, min, max)
	}
}

// Exercise 10: Create a function to reverse a string using pointers
// Operation: Reverse a string by modifying it via a pointer
func reverseString(ptr *string) {
	// TODO: Reverse the string pointed to by ptr
}

func testReverseString() {
	str := "hello"
	reverseString(&str)
	expected := "olleh"
	if str == expected {
		fmt.Println("testReverseString passed")
	} else {
		fmt.Printf("testReverseString failed: expected %v, got %v\n", expected, str)
	}
}

// Run all tests
func runTests() {
	testModifyLargeStruct()
	testResetToZero()
	testGetPointerToInt()
	testCompareInts()
	testIncrementBy()
	testAppendToSlice()
	testSquareAndCube()
	testDivideAndRemainder()
	testMinMax()
	testReverseString()
}

func main() {
	runTests()
}
