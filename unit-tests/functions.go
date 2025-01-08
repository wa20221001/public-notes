package main

import (
	"errors"
	"strings"
)

// Function 1: Return the sum of two integers
func add(a, b int) int {
	return a + b
}

// Function 2: Check if a string contains another substring
func containsSubstring(str, substr string) bool {
	return strings.Contains(str, substr)
}

// Function 3: Calculate the factorial of a non-negative integer
func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial of negative number is undefined")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// Function 4: Reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Function 5: Find the minimum value in a slice of integers
func minInSlice(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("slice is empty")
	}
	min := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min, nil
}

// Function 6: Count the number of words in a string
func countWords(s string) int {
	return len(strings.Fields(s))
}

// Function 7: Check if two slices are equal
func slicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

// Function 8: Check if a number is a palindrome
func isPalindromeNumber(n int) bool {
	if n < 0 {
		return false
	}
	original, reversed := n, 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}

// Function 9: Convert a string to title case (capitalize each word)
func toTitleCase(s string) string {
	return strings.Title(s)
}

// Function 10: Find the intersection of two slices
func intersection(slice1, slice2 []int) []int {
	m := make(map[int]bool)
	result := []int{}
	for _, val := range slice1 {
		m[val] = true
	}
	for _, val := range slice2 {
		if m[val] {
			result = append(result, val)
		}
	}
	return result
}
