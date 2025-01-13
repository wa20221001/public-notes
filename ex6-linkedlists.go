package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Problem 1: Reverse Linked List
// Given the head of a singly linked list, reverse the list and return its head.
func reverseList(head *ListNode) *ListNode {
	// TODO: Implement the function
	return nil
}

// Test for "Reverse Linked List"
func testReverseList() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	expected := []int{5, 4, 3, 2, 1}

	result := reverseList(head)
	for _, v := range expected {
		if result == nil || result.Val != v {
			fmt.Printf("Test failed: reverseList(), expected %v, got %v\n", expected, listToSlice(result))
			return
		}
		result = result.Next
	}
	fmt.Println("Test passed: reverseList")
}

// Problem 2: Merge Two Sorted Lists
// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// TODO: Implement the function
	return nil
}

// Test for "Merge Two Sorted Lists"
func testMergeTwoLists() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	expected := []int{1, 1, 2, 3, 4, 4}

	result := mergeTwoLists(list1, list2)
	if listToSlice(result) == nil || !slicesEqual(listToSlice(result), expected) {
		fmt.Printf("Test failed: mergeTwoLists(), expected %v, got %v\n", expected, listToSlice(result))
	} else {
		fmt.Println("Test passed: mergeTwoLists")
	}
}

// Problem 3: Linked List Cycle
// Given the head of a linked list, determine if the list has a cycle in it.
func hasCycle(head *ListNode) bool {
	// TODO: Implement the function
	return false
}

// Test for "Linked List Cycle"
func testHasCycle() {
	head := &ListNode{3, &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}}
	head.Next.Next.Next.Next = head.Next // Create cycle

	expected := true
	result := hasCycle(head)
	if result != expected {
		fmt.Printf("Test failed: hasCycle(), expected %v, got %v\n", expected, result)
	} else {
		fmt.Println("Test passed: hasCycle")
	}
}

// Problem 4: Reorder List
// Given a singly linked list, reorder it such that the nodes are arranged in a specific pattern:
// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
// You must do this in-place without altering the nodes' values.
func reorderList(head *ListNode) {
	// TODO: Implement the function
}

// Test for "Reorder List"
func testReorderList() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	expected := []int{1, 5, 2, 4, 3}

	reorderList(head)
	result := listToSlice(head)
	if !slicesEqual(result, expected) {
		fmt.Printf("Test failed: reorderList(), expected %v, got %v\n", expected, result)
	} else {
		fmt.Println("Test passed: reorderList")
	}
}

// Problem 5: Remove Nth Node From End of List
// Given the head of a linked list, remove the nth node from the end of the list and return its head.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// TODO: Implement the function
	return nil
}

// Test for "Remove Nth Node From End of List"
func testRemoveNthFromEnd() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n := 2
	expected := []int{1, 2, 3, 5}

	result := removeNthFromEnd(head, n)
	if !slicesEqual(listToSlice(result), expected) {
		fmt.Printf("Test failed: removeNthFromEnd(), expected %v, got %v\n", expected, listToSlice(result))
	} else {
		fmt.Println("Test passed: removeNthFromEnd")
	}
}

// Utility function: Convert linked list to slice for comparison
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Utility function: Check if two slices are equal
func slicesEqual(a, b []int) bool {
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

// // Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// Problem 6: Copy List with Random Pointer
// A linked list is given where each node contains an additional random pointer that could point to any node in the list or null.
// Return a deep copy of the list.
type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	// TODO: Implement the function
	return nil
}

// Test for "Copy List with Random Pointer"
func testCopyRandomList() {
	head := &RandomNode{7, nil, nil}
	head.Next = &RandomNode{13, nil, head}
	head.Next.Next = &RandomNode{11, nil, nil}

	result := copyRandomList(head)
	if result == nil || result.Val != head.Val || result.Next.Random != result {
		fmt.Println("Test failed: copyRandomList()")
	} else {
		fmt.Println("Test passed: copyRandomList")
	}
}

// Problem 7: Add Two Numbers
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// TODO: Implement the function
	return nil
}

// Test for "Add Two Numbers"
func testAddTwoNumbers() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	expected := []int{7, 0, 8}

	result := addTwoNumbers(l1, l2)
	if !slicesEqual(listToSlice(result), expected) {
		fmt.Printf("Test failed: addTwoNumbers(), expected %v, got %v\n", expected, listToSlice(result))
	} else {
		fmt.Println("Test passed: addTwoNumbers")
	}
}

// Problem 8: Find the Duplicate Number
// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
// There is only one duplicate number in nums. Return this duplicate number.
// You must solve the problem without modifying the array and using constant extra space.
func findDuplicate(nums []int) int {
	// TODO: Implement the function using Floyd's Tortoise and Hare (Cycle Detection)
	return 0
}

// Test for "Find the Duplicate Number"
func testFindDuplicate() {
	nums := []int{1, 3, 4, 2, 2}
	expected := 2
	result := findDuplicate(nums)
	if result != expected {
		fmt.Printf("Test failed: findDuplicate(%v), expected %d, got %d\n", nums, expected, result)
	} else {
		fmt.Println("Test passed: findDuplicate")
	}
}

// Problem 9: LRU Cache
// Design a data structure that follows the constraints of a Least Recently Used (LRU) Cache.
// Implement the LRUCache class with the following methods:
// - Constructor(capacity int): Initializes the LRU cache with a positive size capacity.
// - Get(key int) int: Returns the value of the key if the key exists, otherwise return -1.
// - Put(key int, value int): Updates the value of the key if the key exists. Otherwise, adds the key-value pair to the cache.
type LRUCache struct {
	// TODO: Add necessary fields
}

func Constructor(capacity int) LRUCache {
	// TODO: Implement the constructor
	return LRUCache{}
}

func (c *LRUCache) Get(key int) int {
	// TODO: Implement Get method
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// TODO: Implement Put method
}

// Test for "LRU Cache"
func testLRUCache() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if result := cache.Get(1); result != 1 {
		fmt.Printf("Test failed: LRUCache.Get(1), expected %d, got %d\n", 1, result)
		return
	}
	cache.Put(3, 3) // Evicts key 2
	if result := cache.Get(2); result != -1 {
		fmt.Printf("Test failed: LRUCache.Get(2), expected %d, got %d\n", -1, result)
		return
	}
	fmt.Println("Test passed: LRUCache")
}

// Problem 10: Merge K Sorted Lists
// You are given an array of k linked lists, each sorted in ascending order.
// Merge all the linked lists into one sorted linked list and return it.
func mergeKLists(lists []*ListNode) *ListNode {
	// TODO: Implement the function using a min-heap
	return nil
}

// Test for "Merge K Sorted Lists"
func testMergeKLists() {
	list1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 := &ListNode{2, &ListNode{6, nil}}
	lists := []*ListNode{list1, list2, list3}
	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}

	result := mergeKLists(lists)
	if !slicesEqual(listToSlice(result), expected) {
		fmt.Printf("Test failed: mergeKLists(), expected %v, got %v\n", expected, listToSlice(result))
	} else {
		fmt.Println("Test passed: mergeKLists")
	}
}

// Problem 11: Reverse Nodes in k-Group
// Given a linked list, reverse the nodes of the list k at a time, and return the modified list.
// k is a positive integer and is less than or equal to the length of the linked list.
func reverseKGroup(head *ListNode, k int) *ListNode {
	// TODO: Implement the function
	return nil
}

// Test for "Reverse Nodes in k-Group"
func testReverseKGroup() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	k := 2
	expected := []int{2, 1, 4, 3, 5}

	result := reverseKGroup(head, k)
	if !slicesEqual(listToSlice(result), expected) {
		fmt.Printf("Test failed: reverseKGroup(), expected %v, got %v\n", expected, listToSlice(result))
	} else {
		fmt.Println("Test passed: reverseKGroup")
	}
}

// Run all tests
func runTests() {
	testReverseList()
	testMergeTwoLists()
	testHasCycle()
	testReorderList()
	testRemoveNthFromEnd()
	testCopyRandomList()
	testAddTwoNumbers()
	testFindDuplicate()
	testLRUCache()
	testMergeKLists()
	testReverseKGroup()
}

func main() {
	runTests()
}
