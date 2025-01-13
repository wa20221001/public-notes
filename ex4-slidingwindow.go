package main

import (
	"fmt"
)

// Problem 1: Maximum Sum of Subarray of Size K
// Given an array of integers and a number K, find the maximum sum of any contiguous subarray of size K.
func maxSumSubarray(arr []int, k int) int {
	// TODO: Implement the function using the sliding window technique
	return 0
}

// Test for "Maximum Sum of Subarray of Size K"
func testMaxSumSubarray() {
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3
	expected := 9
	result := maxSumSubarray(arr, k)
	if result != expected {
		fmt.Printf("Test failed: maxSumSubarray(%v, %d), expected %d, got %d\n", arr, k, expected, result)
	} else {
		fmt.Println("Test passed: maxSumSubarray")
	}
}

// Problem 2: Smallest Subarray with a Given Sum
// Given an array of positive integers and a positive integer S, find the length of the smallest contiguous subarray
// whose sum is greater than or equal to S. Return 0 if no such subarray exists.
func minSubarrayLen(s int, arr []int) int {
	// TODO: Implement the function using the sliding window technique
	return 0
}

// Test for "Smallest Subarray with a Given Sum"
func testMinSubarrayLen() {
	s := 7
	arr := []int{2, 3, 1, 2, 4, 3}
	expected := 2
	result := minSubarrayLen(s, arr)
	if result != expected {
		fmt.Printf("Test failed: minSubarrayLen(%d, %v), expected %d, got %d\n", s, arr, expected, result)
	} else {
		fmt.Println("Test passed: minSubarrayLen")
	}
}

// Problem 3: Longest Substring with K Distinct Characters
// Given a string, find the length of the longest substring that contains K distinct characters.
func longestSubstringKDistinct(s string, k int) int {
	// TODO: Implement the function using the sliding window technique
	return 0
}

// Test for "Longest Substring with K Distinct Characters"
func testLongestSubstringKDistinct() {
	s := "araaci"
	k := 2
	expected := 4
	result := longestSubstringKDistinct(s, k)
	if result != expected {
		fmt.Printf("Test failed: longestSubstringKDistinct(%s, %d), expected %d, got %d\n", s, k, expected, result)
	} else {
		fmt.Println("Test passed: longestSubstringKDistinct")
	}
}

// Problem 4: Maximum Number of Vowels in a Substring of Given Length
// Given a string and an integer K, find the maximum number of vowels in any contiguous substring of length K.
func maxVowels(s string, k int) int {
	// TODO: Implement the function using the sliding window technique
	return 0
}

// Test for "Maximum Number of Vowels in a Substring of Given Length"
func testMaxVowels() {
	s := "abciiidef"
	k := 3
	expected := 3
	result := maxVowels(s, k)
	if result != expected {
		fmt.Printf("Test failed: maxVowels(%s, %d), expected %d, got %d\n", s, k, expected, result)
	} else {
		fmt.Println("Test passed: maxVowels")
	}
}

// Problem 5: Longest Repeating Character Replacement
// Given a string s and an integer k, return the length of the longest substring
// containing the same letter you can get after performing at most k character replacements.
func characterReplacement(s string, k int) int {
	// TODO: Implement the function using the sliding window technique
	return 0
}

// Test for "Longest Repeating Character Replacement"
func testCharacterReplacement() {
	s := "AABABBA"
	k := 1
	expected := 4
	result := characterReplacement(s, k)
	if result != expected {
		fmt.Printf("Test failed: characterReplacement(%s, %d), expected %d, got %d\n", s, k, expected, result)
	} else {
		fmt.Println("Test passed: characterReplacement")
	}
}

// // Problem 6: Longest Substring Without Repeating Characters
// // Given a string, find the length of the longest substring without repeating characters.
// func lengthOfLongestSubstring(s string) int {
// 	// TODO: Implement the function using the sliding window technique
// 	return 0
// }

// // Test for "Longest Substring Without Repeating Characters"
// func testLengthOfLongestSubstring() {
// 	s := "abcabcbb"
// 	expected := 3
// 	result := lengthOfLongestSubstring(s)
// 	if result != expected {
// 		fmt.Printf("Test failed: lengthOfLongestSubstring(%s), expected %d, got %d\n", s, expected, result)
// 	} else {
// 		fmt.Println("Test passed: lengthOfLongestSubstring")
// 	}
// }

// Problem 7: Permutation in String
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, otherwise return false.
func checkInclusion(s1 string, s2 string) bool {
	// TODO: Implement the function using the sliding window technique
	return false
}

// Test for "Permutation in String"
func testCheckInclusion() {
	s1 := "ab"
	s2 := "eidbaooo"
	expected := true
	result := checkInclusion(s1, s2)
	if result != expected {
		fmt.Printf("Test failed: checkInclusion(%s, %s), expected %v, got %v\n", s1, s2, expected, result)
	} else {
		fmt.Println("Test passed: checkInclusion")
	}
}

// Problem 8: Sliding Window Maximum
// Given an array of integers nums and a sliding window of size k, return the maximum values in the window.
func maxSlidingWindow(nums []int, k int) []int {
	// TODO: Implement the function using the sliding window technique
	return []int{}
}

// Test for "Sliding Window Maximum"
func testMaxSlidingWindow() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	expected := []int{3, 3, 5, 5, 6, 7}
	result := maxSlidingWindow(nums, k)
	if len(result) != len(expected) {
		fmt.Printf("Test failed: maxSlidingWindow(%v, %d), expected %v, got %v\n", nums, k, expected, result)
		return
	}
	for i := range result {
		if result[i] != expected[i] {
			fmt.Printf("Test failed: maxSlidingWindow(%v, %d), expected %v, got %v\n", nums, k, expected, result)
			return
		}
	}
	fmt.Println("Test passed: maxSlidingWindow")
}

// Problem 9: Minimum Window Substring
// Given two strings s and t of lengths m and n respectively, return the minimum window
// in s which will contain all the characters in t. If there is no such window, return an empty string.
func minWindow(s string, t string) string {
	// TODO: Implement the function using the sliding window technique
	return ""
}

// Test for "Minimum Window Substring"
func testMinWindow() {
	s := "ADOBECODEBANC"
	t := "ABC"
	expected := "BANC"
	result := minWindow(s, t)
	if result != expected {
		fmt.Printf("Test failed: minWindow(%s, %s), expected %s, got %s\n", s, t, expected, result)
	} else {
		fmt.Println("Test passed: minWindow")
	}
}

// Problem 10: Best Time to Buy and Sell a Stock
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
func maxProfit(prices []int) int {
	// TODO: Implement the function using the sliding window technique
	return 0
}

// Test for "Best Time to Buy and Sell a Stock"
func testMaxProfit() {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	result := maxProfit(prices)
	if result != expected {
		fmt.Printf("Test failed: maxProfit(%v), expected %d, got %d\n", prices, expected, result)
	} else {
		fmt.Println("Test passed: maxProfit")
	}
}

// Problem 11: Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	// TODO: Implement the function using the sliding window technique
	return 0
}

// Test for "Longest Substring Without Repeating Characters"
func testLengthOfLongestSubstring() {
	s := "abcabcbb"
	expected := 3
	result := lengthOfLongestSubstring(s)
	if result != expected {
		fmt.Printf("Test failed: lengthOfLongestSubstring(%s), expected %d, got %d\n", s, expected, result)
	} else {
		fmt.Println("Test passed: lengthOfLongestSubstring")
	}
}

// Run all tests
func runTests() {
	testMaxSumSubarray()
	testMinSubarrayLen()
	testLongestSubstringKDistinct()
	testMaxVowels()
	testCharacterReplacement()
	testLengthOfLongestSubstring()
	testCheckInclusion()
	testMaxSlidingWindow()
	testMinWindow()
	testMaxProfit()
}

func main() {
	runTests()
}
