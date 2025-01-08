package main

import "fmt"

// Exercise 1: Basic pointer usage
// Operation: Create a pointer to an integer and update its value through the pointer
func updateValue(ptr *int) {
	// TODO: Update the value pointed to by ptr to 42
}

func testUpdateValue() {
	x := 10
	updateValue(&x)
	expected := 42
	if x == expected {
		fmt.Println("testUpdateValue passed")
	} else {
		fmt.Printf("testUpdateValue failed: expected %v, got %v\n", expected, x)
	}
}

// Exercise 2: Swap two integers using pointers
// Operation: Use pointers to swap the values of two integers
func swap(a, b *int) {
	// TODO: Swap the values pointed to by a and b
}

func testSwap() {
	x, y := 5, 10
	swap(&x, &y)
	if x == 10 && y == 5 {
		fmt.Println("testSwap passed")
	} else {
		fmt.Printf("testSwap failed: expected x=10, y=5; got x=%v, y=%v\n", x, y)
	}
}

// Exercise 3: Find the maximum of two integers using pointers
// Operation: Use pointers to return the maximum of two integers
func max(a, b *int) int {
	// TODO: Return the greater of the two values pointed to by a and b
	return 0
}

func testMax() {
	x, y := 7, 3
	expected := 7
	result := max(&x, &y)
	if result == expected {
		fmt.Println("testMax passed")
	} else {
		fmt.Printf("testMax failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 4: Create a function that doubles a number using a pointer
// Operation: Modify a value through a pointer
func doubleValue(ptr *int) {
	// TODO: Double the value pointed to by ptr
}

func testDoubleValue() {
	x := 4
	doubleValue(&x)
	expected := 8
	if x == expected {
		fmt.Println("testDoubleValue passed")
	} else {
		fmt.Printf("testDoubleValue failed: expected %v, got %v\n", expected, x)
	}
}

// Exercise 5: Initialize a pointer to nil and check if it's nil
// Operation: Pointer initialization and nil check
func isNilPointer(ptr *int) bool {
	// TODO: Return true if ptr is nil, otherwise false
	return false
}

func testIsNilPointer() {
	var ptr *int
	expected := true
	result := isNilPointer(ptr)
	if result == expected {
		fmt.Println("testIsNilPointer passed")
	} else {
		fmt.Printf("testIsNilPointer failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 6: Change a string through a pointer
// Operation: Modify a string value through a pointer
func changeString(ptr *string) {
	// TODO: Change the value pointed to by ptr to "Hello, Go!"
}

func testChangeString() {
	str := "Hello"
	changeString(&str)
	expected := "Hello, Go!"
	if str == expected {
		fmt.Println("testChangeString passed")
	} else {
		fmt.Printf("testChangeString failed: expected %v, got %v\n", expected, str)
	}
}

// Exercise 7: Create a pointer to a struct and modify its fields
// Operation: Modify struct fields through a pointer
type Person struct {
	Name string
	Age  int
}

func updatePerson(ptr *Person) {
	// TODO: Update the Name to "John" and Age to 30
}

func testUpdatePerson() {
	p := Person{Name: "Alice", Age: 25}
	updatePerson(&p)
	expectedName, expectedAge := "John", 30
	if p.Name == expectedName && p.Age == expectedAge {
		fmt.Println("testUpdatePerson passed")
	} else {
		fmt.Printf("testUpdatePerson failed: expected Name=%v, Age=%v; got Name=%v, Age=%v\n",
			expectedName, expectedAge, p.Name, p.Age)
	}
}

// Exercise 8: Create a function to reverse an array in place using pointers
// Operation: Array manipulation using pointers
func reverseArray(arr []int) {
	// TODO: Use pointers to reverse the array in place
}

func testReverseArray() {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	reverseArray(arr)
	if fmt.Sprintf("%v", arr) == fmt.Sprintf("%v", expected) {
		fmt.Println("testReverseArray passed")
	} else {
		fmt.Printf("testReverseArray failed: expected %v, got %v\n", expected, arr)
	}
}

// Exercise 9: Return a pointer to a newly created integer
// Operation: Return a pointer from a function
func newInt(value int) *int {
	// TODO: Create a new integer with the given value and return a pointer to it
	return nil
}

func testNewInt() {
	value := 10
	ptr := newInt(value)
	if ptr != nil && *ptr == value {
		fmt.Println("testNewInt passed")
	} else {
		fmt.Printf("testNewInt failed: expected %v, got %v\n", value, ptr)
	}
}

// Exercise 10: Create a function that swaps two elements in a slice using pointers
// Operation: Slice element manipulation using pointers
func swapElements(slice []int, i, j int) {
	// TODO: Swap elements at indices i and j using pointers
}

func testSwapElements() {
	slice := []int{1, 2, 3, 4, 5}
	swapElements(slice, 1, 3)
	expected := []int{1, 4, 3, 2, 5}
	if fmt.Sprintf("%v", slice) == fmt.Sprintf("%v", expected) {
		fmt.Println("testSwapElements passed")
	} else {
		fmt.Printf("testSwapElements failed: expected %v, got %v\n", expected, slice)
	}
}

// Run all tests
func runTests() {
	testUpdateValue()
	testSwap()
	testMax()
	testDoubleValue()
	testIsNilPointer()
	testChangeString()
	testUpdatePerson()
	testReverseArray()
	testNewInt()
	testSwapElements()
}

func main() {
	runTests()
}
