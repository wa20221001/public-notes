package main

import (
	"fmt"
	"reflect"
)

// Exercise 1: Reverse a string
func reverseString(s string) string {
	input := []rune(s)
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return string(input)
}

func testReverseString() {
	input := "hello"
	expected := "olleh"
	result := reverseString(input)
	if result == expected {
		fmt.Println("testReverseString passed")
	} else {
		fmt.Printf("testReverseString failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 2: Count runes in a string
func countRunes(s string) int {
	return len([]rune(s))
}

func testCountRunes() {
	input := "こんにちは"
	expected := 5
	result := countRunes(input)
	if result == expected {
		fmt.Println("testCountRunes passed")
	} else {
		fmt.Printf("testCountRunes failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 3: Find the maximum in an array
func maxInArray(arr []int) int {
	max := 0
	for _, j := range arr {
		if j > max {
			max = j
		}
	}
	return max
}

func testMaxInArray() {
	input := []int{1, 3, 2, 8, 5}
	expected := 8
	result := maxInArray(input)
	if result == expected {
		fmt.Println("testMaxInArray passed")
	} else {
		fmt.Printf("testMaxInArray failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 4: Remove duplicates from a slice
func removeDuplicates(slice []int) []int {
	// TODO: Implement removing duplicates from the slice
	return nil
}

func testRemoveDuplicates() {
	input := []int{1, 2, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	result := removeDuplicates(input)
	if reflect.DeepEqual(result, expected) {
		fmt.Println("testRemoveDuplicates passed")
	} else {
		fmt.Printf("testRemoveDuplicates failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 5: Concatenate a slice of strings
func concatenateStrings(slice []string) string {
	// TODO: Implement concatenating the slice of strings
	return ""
}

func testConcatenateStrings() {
	input := []string{"Go", "is", "awesome"}
	expected := "Go is awesome"
	result := concatenateStrings(input)
	if result == expected {
		fmt.Println("testConcatenateStrings passed")
	} else {
		fmt.Printf("testConcatenateStrings failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 6: Check if a map contains a key
func mapContainsKey(m map[string]int, key string) bool {
	// TODO: Implement checking if the map contains the key
	return false
}

func testMapContainsKey() {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	key := "b"
	expected := true
	result := mapContainsKey(input, key)
	if result == expected {
		fmt.Println("testMapContainsKey passed")
	} else {
		fmt.Printf("testMapContainsKey failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 7: Find the sum of an array
func sumArray(arr []int) int {
	// TODO: Implement summing the elements of the array
	return 0
}

func testSumArray() {
	input := []int{1, 2, 3, 4, 5}
	expected := 15
	result := sumArray(input)
	if result == expected {
		fmt.Println("testSumArray passed")
	} else {
		fmt.Printf("testSumArray failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 8: Find the longest word in a string
func longestWord(s string) string {
	// TODO: Implement finding the longest word in the string
	return ""
}

func testLongestWord() {
	input := "Golang is fun to learn"
	expected := "Golang"
	result := longestWord(input)
	if result == expected {
		fmt.Println("testLongestWord passed")
	} else {
		fmt.Printf("testLongestWord failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 9: Check if a slice is sorted
func isSorted(slice []int) bool {
	// TODO: Implement checking if the slice is sorted in ascending order
	return false
}

func testIsSorted() {
	input := []int{1, 2, 3, 4, 5}
	expected := true
	result := isSorted(input)
	if result == expected {
		fmt.Println("testIsSorted passed")
	} else {
		fmt.Printf("testIsSorted failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 10: Merge two maps
func mergeMaps(m1, m2 map[string]int) map[string]int {
	// TODO: Implement merging two maps
	return nil
}

func testMergeMaps() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"c": 3, "d": 4}
	expected := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	result := mergeMaps(m1, m2)
	if reflect.DeepEqual(result, expected) {
		fmt.Println("testMergeMaps passed")
	} else {
		fmt.Printf("testMergeMaps failed: expected %v, got %v\n", expected, result)
	}
}

// Run all tests
func runTests() {
	testReverseString()
	testCountRunes()
	testMaxInArray()
	testRemoveDuplicates()
	testConcatenateStrings()
	testMapContainsKey()
	testSumArray()
	testLongestWord()
	testIsSorted()
	testMergeMaps()
}

func main() {
	runTests()
}
