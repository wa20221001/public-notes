package main

import (
	"fmt"
	"reflect"
)

// Exercise 1: Convert a string to uppercase
func toUpperCase(s string) string {
	// TODO: Implement converting the string s to uppercase
	return ""
}

func testToUpperCase() {
	input := "golang"
	expected := "GOLANG"
	result := toUpperCase(input)
	if result == expected {
		fmt.Println("testToUpperCase passed")
	} else {
		fmt.Printf("testToUpperCase failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 2: Check if a string is a palindrome
func isPalindrome(s string) bool {
	// TODO: Implement checking if the string s is a palindrome
	return false
}

func testIsPalindrome() {
	input := "madam"
	expected := true
	result := isPalindrome(input)
	if result == expected {
		fmt.Println("testIsPalindrome passed")
	} else {
		fmt.Printf("testIsPalindrome failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 3: Replace spaces with dashes in a string
func replaceSpaces(s string) string {
	// TODO: Implement replacing spaces with dashes in the string
	return ""
}

func testReplaceSpaces() {
	input := "Go is great"
	expected := "Go-is-great"
	result := replaceSpaces(input)
	if result == expected {
		fmt.Println("testReplaceSpaces passed")
	} else {
		fmt.Printf("testReplaceSpaces failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 4: Find the frequency of each rune in a string
func runeFrequency(s string) map[rune]int {
	// TODO: Implement finding the frequency of each rune in the string
	return nil
}

func testRuneFrequency() {
	input := "banana"
	expected := map[rune]int{'b': 1, 'a': 3, 'n': 2}
	result := runeFrequency(input)
	if reflect.DeepEqual(result, expected) {
		fmt.Println("testRuneFrequency passed")
	} else {
		fmt.Printf("testRuneFrequency failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 5: Find the index of a target element in a slice
func indexOf(slice []int, target int) int {
	// TODO: Implement finding the index of the target element
	return -1
}

func testIndexOf() {
	input := []int{10, 20, 30, 40}
	target := 30
	expected := 2
	result := indexOf(input, target)
	if result == expected {
		fmt.Println("testIndexOf passed")
	} else {
		fmt.Printf("testIndexOf failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 6: Create a slice of even numbers from 1 to n
func evenNumbers(n int) []int {
	// TODO: Implement generating a slice of even numbers from 1 to n
	return nil
}

func testEvenNumbers() {
	input := 10
	expected := []int{2, 4, 6, 8, 10}
	result := evenNumbers(input)
	if reflect.DeepEqual(result, expected) {
		fmt.Println("testEvenNumbers passed")
	} else {
		fmt.Printf("testEvenNumbers failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 7: Check if a map is empty
func isMapEmpty(m map[string]int) bool {
	// TODO: Implement checking if the map is empty
	return false
}

func testIsMapEmpty() {
	input := map[string]int{}
	expected := true
	result := isMapEmpty(input)
	if result == expected {
		fmt.Println("testIsMapEmpty passed")
	} else {
		fmt.Printf("testIsMapEmpty failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 8: Count vowels in a string
func countVowels(s string) int {
	// TODO: Implement counting vowels in the string
	return 0
}

func testCountVowels() {
	input := "golang"
	expected := 2
	result := countVowels(input)
	if result == expected {
		fmt.Println("testCountVowels passed")
	} else {
		fmt.Printf("testCountVowels failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 9: Remove all spaces from a string
func removeSpaces(s string) string {
	// TODO: Implement removing all spaces from the string
	return ""
}

func testRemoveSpaces() {
	input := "Go is fun"
	expected := "Goisfun"
	result := removeSpaces(input)
	if result == expected {
		fmt.Println("testRemoveSpaces passed")
	} else {
		fmt.Printf("testRemoveSpaces failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 10: Merge two slices
func mergeSlices(s1, s2 []int) []int {
	// TODO: Implement merging two slices
	return nil
}

func testMergeSlices() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6}
	result := mergeSlices(s1, s2)
	if reflect.DeepEqual(result, expected) {
		fmt.Println("testMergeSlices passed")
	} else {
		fmt.Printf("testMergeSlices failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 11: Find the first non-repeating character in a string
func firstNonRepeatingChar(s string) rune {
	// TODO: Implement finding the first non-repeating character
	return 0
}

func testFirstNonRepeatingChar() {
	input := "swiss"
	expected := 'w'
	result := firstNonRepeatingChar(input)
	if result == expected {
		fmt.Println("testFirstNonRepeatingChar passed")
	} else {
		fmt.Printf("testFirstNonRepeatingChar failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 12: Find the common elements of two slices
func commonElements(s1, s2 []int) []int {
	// TODO: Implement finding the common elements of two slices
	return nil
}

func testCommonElements() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{3, 4, 5, 6}
	expected := []int{3, 4}
	result := commonElements(s1, s2)
	if reflect.DeepEqual(result, expected) {
		fmt.Println("testCommonElements passed")
	} else {
		fmt.Printf("testCommonElements failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 13: Capitalize the first letter of each word in a string
func capitalizeWords(s string) string {
	// TODO: Implement capitalizing the first letter of each word
	return ""
}

func testCapitalizeWords() {
	input := "golang is great"
	expected := "Golang Is Great"
	result := capitalizeWords(input)
	if result == expected {
		fmt.Println("testCapitalizeWords passed")
	} else {
		fmt.Printf("testCapitalizeWords failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 14: Remove all digits from a string
func removeDigits(s string) string {
	// TODO: Implement removing all digits from the string
	return ""
}

func testRemoveDigits() {
	input := "Go123lang"
	expected := "Golang"
	result := removeDigits(input)
	if result == expected {
		fmt.Println("testRemoveDigits passed")
	} else {
		fmt.Printf("testRemoveDigits failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 15: Reverse the order of words in a string
func reverseWords(s string) string {
	// TODO: Implement reversing the order of words in the string
	return ""
}

func testReverseWords() {
	input := "Go is fun"
	expected := "fun is Go"
	result := reverseWords(input)
	if result == expected {
		fmt.Println("testReverseWords passed")
	} else {
		fmt.Printf("testReverseWords failed: expected %v, got %v\n", expected, result)
	}
}

// Run all tests
func runTests() {
	testToUpperCase()
	testIsPalindrome()
	testReplaceSpaces()
	testRuneFrequency()
	testIndexOf()
	testEvenNumbers()
	testIsMapEmpty()
	testCountVowels()
	testRemoveSpaces()
	testMergeSlices()
	testFirstNonRepeatingChar()
	testCommonElements()
	testCapitalizeWords()
	testRemoveDigits()
	testReverseWords()
}

func main() {
	runTests()
}
