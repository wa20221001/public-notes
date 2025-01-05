
package main

import (
	"fmt"
)

func main() {
	runTests()
}

// Exercise 1: Reverse a String
// Problem: Write a function that reverses a string.
// Example: Given "hello", the function should return "olleh".

func reverseString(s string) string {
	// Your code here
	return ""
}

func testReverseString() {
	input := "hello"
	expected := "olleh"
	result := reverseString(input)
	if result != expected {
		fmt.Println("Test failed: reverseString")
	} else {
		fmt.Println("Test passed: reverseString")
	}
}

// Exercise 3: Check for Anagram
// Problem: Write a function that checks if two strings are anagrams of each other.
// Example: Given "listen" and "silent", the function should return true.

func isAnagram(s1, s2 string) bool {
	// Your code here
	return false
}

func testIsAnagram() {
	input1, input2 := "listen", "silent"
	expected := true
	result := isAnagram(input1, input2)
	if result != expected {
		fmt.Println("Test failed: isAnagram")
	} else {
		fmt.Println("Test passed: isAnagram")
	}
}

func reverseArray(arr []int) []int {
	// Your code here
	return nil
}

// Exercise 4: Remove Duplicates from Slice
// Problem: Write a function that removes duplicates from an integer slice.
// Example: Given [1, 2, 2, 3, 4, 4, 5], the function should return [1, 2, 3, 4, 5].

func removeDuplicates(arr []int) []int {
	// Your code here
	return nil
}

func testRemoveDuplicates() {
	input := []int{1, 2, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	result := removeDuplicates(input)
	if !equalSlices(result, expected) {
		fmt.Println("Test failed: removeDuplicates")
	} else {
		fmt.Println("Test passed: removeDuplicates")
	}
}

// Exercise 5: First Unique Character in String
// Problem: Write a function that returns the index of the first unique character in a string.
// Example: Given "swiss", the function should return 0.

func firstUniqueChar(s string) int {
	// Your code here
	return -1
}

func testFirstUniqueChar() {
	input := "swiss"
	expected := 0
	result := firstUniqueChar(input)
	if result != expected {
		fmt.Println("Test failed: firstUniqueChar")
	} else {
		fmt.Println("Test passed: firstUniqueChar")
	}
}

// Exercise 6: Group Anagrams
// Problem: Write a function that groups anagrams together from a list of strings.
// Example: Given ["eat", "tea", "tan", "ate", "nat", "bat"], the function should return [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]].

func groupAnagrams(strs []string) [][]string {
	// Your code here
	return nil
}

func testGroupAnagrams() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	result := groupAnagrams(input)
	if !equalNestedSlices(result, expected) {
		fmt.Println("Test failed: groupAnagrams")
	} else {
		fmt.Println("Test passed: groupAnagrams")
	}
}

// Helper function to compare nested slices
func equalNestedSlices(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalSlicesStr(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalSlicesStr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Exercise 7: Longest Common Prefix
// Problem: Write a function that finds the longest common prefix among an array of strings.
// Example: Given ["flower", "flow", "flight"], the function should return "fl".

func longestCommonPrefix(strs []string) string {
	// Your code here
	return ""
}

func testLongestCommonPrefix() {
	input := []string{"flower", "flow", "flight"}
	expected := "fl"
	result := longestCommonPrefix(input)
	if result != expected {
		fmt.Println("Test failed: longestCommonPrefix")
	} else {
		fmt.Println("Test passed: longestCommonPrefix")
	}
}

// Exercise 8: Check Palindrome
// Problem: Write a function that checks if a string is a palindrome.
// Example: Given "racecar", the function should return true.

func isPalindrome(s string) bool {
	// Your code here
	return false
}

func testIsPalindrome() {
	input := "racecar"
	expected := true
	result := isPalindrome(input)
	if result != expected {
		fmt.Println("Test failed: isPalindrome")
	} else {
		fmt.Println("Test passed: isPalindrome")
	}
}

// Exercise 9: Maximum Subarray Sum
// Problem: Write a function that finds the maximum sum of a contiguous subarray.
// Example: Given [-2, 1, -3, 4, -1, 2, 1, -5, 4], the function should return 6.

func maxSubArray(nums []int) int {
	// Your code here
	return 0
}

func testMaxSubArray() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	result := maxSubArray(input)
	if result != expected {
		fmt.Println("Test failed: maxSubArray")
	} else {
		fmt.Println("Test passed: maxSubArray")
	}
}

// Exercise 10: Find All Duplicates in Array
// Problem: Write a function that returns all elements that appear twice in an array.
// Example: Given [4, 3, 2, 7, 8, 2, 3, 1], the function should return [2, 3].

func findDuplicates(nums []int) []int {
	// Your code here
	return nil
}

func testFindDuplicates() {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{2, 3}
	result := findDuplicates(input)
	if !equalSlices(result, expected) {
		fmt.Println("Test failed: findDuplicates")
	} else {
		fmt.Println("Test passed: findDuplicates")
	}
}

// Unit Test for Exercise 1
func testReverseArray() {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	result := reverseArray(input)
	if !equalSlices(result, expected) {
		fmt.Println("Test failed: reverseArray")
	} else {
		fmt.Println("Test passed: reverseArray")
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Exercise 2: Count Vowels in a String
// Problem: Write a function that takes a string and returns the number of vowels (a, e, i, o, u) in it.
// Example: Given "golang", the function should return 2.

func countVowels(s string) int {
	// Your code here
	return 0
}

func testCountVowels() {
	input := "golang"
	expected := 2
	result := countVowels(input)
	if result != expected {
		fmt.Println("Test failed: countVowels")
	} else {
		fmt.Println("Test passed: countVowels")
	}
}

// Exercise 4: Convert String to Runes and Reverse
// Problem: Write a function that converts a string to a slice of runes, reverses it, and returns the reversed string.
// Example: Given "hello", the function should return "olleh".

func reverseStringWithRunes(s string) string {
	// Your code here
	return ""
}

func testReverseStringWithRunes() {
	input := "hello"
	expected := "olleh"
	result := reverseStringWithRunes(input)
	if result != expected {
		fmt.Println("Test failed: reverseStringWithRunes")
	} else {
		fmt.Println("Test passed: reverseStringWithRunes")
	}
}

// Exercise 5: Find Maximum Element in an Array
// Problem: Write a function that returns the maximum element in an array of integers.
// Example: Given [1, 3, 7, 2, 5], the function should return 7.

func findMax(arr []int) int {
	// Your code here
	return 0
}

func testFindMax() {
	input := []int{1, 3, 7, 2, 5}
	expected := 7
	result := findMax(input)
	if result != expected {
		fmt.Println("Test failed: findMax")
	} else {
		fmt.Println("Test passed: findMax")
	}
}

// Exercise 6: Merge Two Sorted Arrays
// Problem: Write a function that merges two sorted arrays into one sorted array.
// Example: Given [1, 3, 5] and [2, 4, 6], the function should return [1, 2, 3, 4, 5, 6].

func mergeSortedArrays(arr1, arr2 []int) []int {
	// Your code here
	return nil
}

func testMergeSortedArrays() {
	input1 := []int{1, 3, 5}
	input2 := []int{2, 4, 6}
	expected := []int{1, 2, 3, 4, 5, 6}
	result := mergeSortedArrays(input1, input2)
	if !equalSlices(result, expected) {
		fmt.Println("Test failed: mergeSortedArrays")
	} else {
		fmt.Println("Test passed: mergeSortedArrays")
	}
}

// Exercise 7: Find the First Non-Repeating Character
// Problem: Write a function that finds the first non-repeating character in a string.
// Example: Given "swiss", the function should return 'w'.

func firstNonRepeatingChar(s string) rune {
	// Your code here
	return ' '
}

func testFirstNonRepeatingChar() {
	input := "swiss"
	expected := 'w'
	result := firstNonRepeatingChar(input)
	if result != expected {
		fmt.Println("Test failed: firstNonRepeatingChar")
	} else {
		fmt.Println("Test passed: firstNonRepeatingChar")
	}
}

// Exercise 8: Rotate Array
// Problem: Write a function that rotates an array to the right by k steps.
// Example: Given [1, 2, 3, 4, 5] and k = 2, the function should return [4, 5, 1, 2, 3].
// Hint: Use modulus to wrap around the array length.

func rotateArray(arr []int, k int) []int {
	// Your code here
	return nil
}

func testRotateArray() {
	input := []int{1, 2, 3, 4, 5}
	k := 2
	expected := []int{4, 5, 1, 2, 3}
	result := rotateArray(input, k)
	if !equalSlices(result, expected) {
		fmt.Println("Test failed: rotateArray")
	} else {
		fmt.Println("Test passed: rotateArray")
	}
}

// Exercise 9: Find Missing Number in Array
// Problem: Write a function to find the missing number in an array containing numbers from 1 to n.
// Example: Given [1, 2, 4, 5, 6], the function should return 3.

func findMissingNumber(arr []int) int {
	// Your code here
	return 0
}

func testFindMissingNumber() {
	input := []int{1, 2, 4, 5, 6}
	expected := 3
	result := findMissingNumber(input)
	if result != expected {
		fmt.Println("Test failed: findMissingNumber")
	} else {
		fmt.Println("Test passed: findMissingNumber")
	}
}

// Exercise 11: Check if Two Strings Are Anagrams
// Problem: Write a function that checks if two strings are anagrams.
// Example: Given "listen" and "silent", the function should return true.

func areAnagrams(s1, s2 string) bool {
	// Your code here
	return false
}

func testAreAnagrams() {
	input1 := "listen"
	input2 := "silent"
	expected := true
	result := areAnagrams(input1, input2)
	if result != expected {
		fmt.Println("Test failed: areAnagrams")
	} else {
		fmt.Println("Test passed: areAnagrams")
	}
}

// Exercise 12: Sum of Even Numbers in a Map
// Problem: Write a function that takes a map with integer keys and values and returns the sum of all even values.
// Example: Given map[int]int{1: 2, 2: 3, 3: 4}, the function should return 6.

func sumEvenValues(m map[int]int) int {
	// Your code here
	return 0
}

func testSumEvenValues() {
	input := map[int]int{1: 2, 2: 3, 3: 4}
	expected := 6
	result := sumEvenValues(input)
	if result != expected {
		fmt.Println("Test failed: sumEvenValues")
	} else {
		fmt.Println("Test passed: sumEvenValues")
	}
}

// Exercise 13: Group Elements by Frequency
// Problem: Write a function that groups elements in an array by their frequency and returns a map where the keys are the elements and the values are their frequencies.
// Example: Given [1, 2, 2, 3, 3, 3], the function should return map[int]int{1: 1, 2: 2, 3: 3}.

func groupByFrequency(arr []int) map[int]int {
	// Your code here
	return nil
}

func testGroupByFrequency() {
	input := []int{1, 2, 2, 3, 3, 3}
	expected := map[int]int{1: 1, 2: 2, 3: 3}
	result := groupByFrequency(input)
	if !equalMaps(result, expected) {
		fmt.Println("Test failed: groupByFrequency")
	} else {
		fmt.Println("Test passed: groupByFrequency")
	}
}

func equalMaps(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// Exercise 14: Find the Most Frequent Element
// Problem: Write a function that finds the most frequent element in an array.
// Example: Given [1, 3, 2, 2, 4, 3, 3], the function should return 3.

func mostFrequentElement(arr []int) int {
	// Your code here
	return 0
}

func testMostFrequentElement() {
	input := []int{1, 3, 2, 2, 4, 3, 3}
	expected := 3
	result := mostFrequentElement(input)
	if result != expected {
		fmt.Println("Test failed: mostFrequentElement")
	} else {
		fmt.Println("Test passed: mostFrequentElement")
	}
}

// Exercise 1: Longest Substring Without Repeating Characters
// Problem: Write a function that finds the length of the longest substring without repeating characters.
// Example: Given "abcabcbb", the function should return 3.

func lengthOfLongestSubstring(s string) int {
	// Your code here
	return 0
}

func testLengthOfLongestSubstring() {
	input := "abcabcbb"
	expected := 3
	result := lengthOfLongestSubstring(input)
	if result != expected {
		fmt.Println("Test failed: lengthOfLongestSubstring")
	} else {
		fmt.Println("Test passed: lengthOfLongestSubstring")
	}
}

// Exercise 3: Check if String is a Valid Parenthesis
// Problem: Write a function that checks if a given string has valid parentheses.
// Example: Given "()", the function should return true. Given "(]", it should return false.

func isValidParenthesis(s string) bool {
	// Your code here
	return false
}

func testIsValidParenthesis() {
	input := "()"
	expected := true
	result := isValidParenthesis(input)
	if result != expected {
		fmt.Println("Test failed: isValidParenthesis")
	} else {
		fmt.Println("Test passed: isValidParenthesis")
	}
}

// Exercise 4: Find Intersection of Two Arrays
// Problem: Write a function that returns the intersection of two integer arrays.
// Example: Given [1, 2, 2, 1] and [2, 2], the function should return [2].

func intersectArrays(nums1, nums2 []int) []int {
	// Your code here
	return nil
}

func testIntersectArrays() {
	input1 := []int{1, 2, 2, 1}
	input2 := []int{2, 2}
	expected := []int{2}
	result := intersectArrays(input1, input2)
	if !equalSlices(result, expected) {
		fmt.Println("Test failed: intersectArrays")
	} else {
		fmt.Println("Test passed: intersectArrays")
	}
}

// Exercise 5: Merge Overlapping Intervals
// Problem: Write a function that merges overlapping intervals in a slice.
// Example: Given [[1,3],[2,6],[8,10],[15,18]], the function should return [[1,6],[8,10],[15,18]].

func mergeIntervals(intervals [][]int) [][]int {
	// Your code here
	return nil
}

func testMergeIntervals() {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	result := mergeIntervals(input)
	if !equalNestedSlicesInt(result, expected) {
		fmt.Println("Test failed: mergeIntervals")
	} else {
		fmt.Println("Test passed: mergeIntervals")
	}
}

func equalNestedSlicesInt(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalSlices(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Exercise 6: Find Top K Frequent Elements
// Problem: Write a function that returns the k most frequent elements in an array.
// Example: Given [1, 1, 1, 2, 2, 3] and k = 2, the function should return [1, 2].

func topKFrequent(nums []int, k int) []int {
	// Your code here
	return nil
}

func testTopKFrequent() {
	input := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}
	result := topKFrequent(input, k)
	if !equalSlices(result, expected) {
		fmt.Println("Test failed: topKFrequent")
	} else {
		fmt.Println("Test passed: topKFrequent")
	}
}

// Exercise 7: Product of Array Except Self
// Problem: Write a function that returns an array where each element is the product of all other elements in the array except itself.
// Example: Given [1, 2, 3, 4], the function should return [24, 12, 8, 6].

func productExceptSelf(nums []int) []int {
	// Your code here
	return nil
}

func testProductExceptSelf() {
	input := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}
	result := productExceptSelf(input)
	if !equalSlices(result, expected) {
		fmt.Println("Test failed: productExceptSelf")
	} else {
		fmt.Println("Test passed: productExceptSelf")
	}
}

// Exercise 8: Find All Unique Triplets that Sum to Zero
// Problem: Write a function that returns all unique triplets in the array that sum to zero.
// Example: Given [-1, 0, 1, 2, -1, -4], the function should return [[-1, -1, 2], [-1, 0, 1]].

func threeSum(nums []int) [][]int {
	// Your code here
	return nil
}

func testThreeSum() {
	input := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	result := threeSum(input)
	if !equalNestedSlicesInt(result, expected) {
		fmt.Println("Test failed: threeSum")
	} else {
		fmt.Println("Test passed: threeSum")
	}
}

// Exercise 10: Length of Longest Consecutive Sequence
// Problem: Write a function that returns the length of the longest consecutive sequence in an unsorted array.
// Example: Given [100, 4, 200, 1, 3, 2], the function should return 4.

func longestConsecutive(nums []int) int {
	// Your code here
	return 0
}

func testLongestConsecutive() {
	input := []int{100, 4, 200, 1, 3, 2}
	expected := 4
	result := longestConsecutive(input)
	if result != expected {
		fmt.Println("Test failed: longestConsecutive")
	} else {
		fmt.Println("Test passed: longestConsecutive")
	}
}

// Run all tests
func runTests() {
	testReverseString()
	testIsAnagram()
	testReverseArray()
	testCountVowels()
	testIsPalindrome()
	testReverseStringWithRunes()
	testFindMax()
	testMergeSortedArrays()
	testFirstNonRepeatingChar()
	testRotateArray()
	testFindMissingNumber()
	testFindDuplicates()
	testAreAnagrams()
	testSumEvenValues()
	testGroupByFrequency()
	testMostFrequentElement()
	testMostFrequentElement()
	testReverseString()
	testCountVowels()
	testRemoveDuplicates()
	testFirstUniqueChar()
	testGroupAnagrams()
	testLongestCommonPrefix()
	testIsPalindrome()
	testMaxSubArray()
	testFindDuplicates()
	testLengthOfLongestSubstring()
	testFindMissingNumber()
	testIsValidParenthesis()
	testIntersectArrays()
	testMergeIntervals()
	testTopKFrequent()
	testProductExceptSelf()
	testThreeSum()
	testRotateArray()
	testLongestConsecutive()
}
