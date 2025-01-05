package main

import (
	"fmt"
)

func main() {
	runTests()
}

func runTests() {
	// Struct and Interface Basic Usage
	testGreet()
	testIsAdult()
	testDescribeStudent()
	// testShapes()
	// testPrintArea()

	// MORE Q
	// testParkingSystem()
	// testBasicElevator()
	// testLogger()
	// testMergeKLists()
	// testFileSystem()
	// testPlaylistManager()
	// testCalculator()
}

// Create a Person struct with fields Name (string) and Age (int). Write a function greet that prints "Hello, my name is {Name} and I am {Age} years old."

type Person struct {
	Name string
	Age  int
}

func greet(p Person) {
	// Your code here
}
func testGreet() {
	p := Person{Name: "John", Age: 25}
	greet(p) // Expected output: "Hello, my name is John and I am 25 years old."
}

// Add a method IsAdult to the Person struct that returns true if the person's age is 18 or above.
func (p Person) IsAdult() bool {
	// Your code here
	return false
}
func testIsAdult() {
	p := Person{Name: "Alice", Age: 20}
	result := p.IsAdult() // Expected: true
	if result {
		fmt.Println("Test passed: IsAdult")
	} else {
		fmt.Println("Test failed: IsAdult")
	}
}

// Create a Student struct that embeds the Person struct and adds a field Grade (int). Write a function describeStudent that prints the student's name, age, and grade.
type Student struct {
	Person
	Grade int
}

func describeStudent(s Student) {
	// Your code here
}

func testDescribeStudent() {
	s := Student{Person: Person{Name: "Mike", Age: 21}, Grade: 90}
	describeStudent(s) // Expected output: "Name: Mike, Age: 21, Grade: 90"
}

// Define an interface Shape with a method Area() float64. Create two structs Circle and Rectangle, and implement the Area method for both.
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// Implement Area method for Circle

// Implement Area method for Rectangle

// Implement Area method for Circle and Rectangle
func testShapes() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	fmt.Printf("Circle area: %.2f\n", c.Area())    // Expected: 78.54 (π * r^2)
	fmt.Printf("Rectangle area: %.2f\n", r.Area()) // Expected: 24.00
}

// Write a function printArea that takes a Shape interface as input and prints the area of the shape.
func printArea(s Shape) {
	// Your code here
}
func testPrintArea() {
	c := Circle{Radius: 7}
	r := Rectangle{Width: 5, Height: 3}

	printArea(c) // Expected: Circle area: 153.94
	printArea(r) // Expected: Rectangle area: 15.00
}

// Problem:
// Design a parking system for a parking lot with three types of parking spaces: big, medium, and small.
// Implement a struct ParkingSystem with the following methods:
// - Constructor(big int, medium int, small int): Initializes the number of available slots for each type.
// - AddCar(carType int) bool: Checks if a car of the given type can be parked.
//    carType can be 1 (big), 2 (medium), or 3 (small).
//    Return true if the car can be parked; otherwise, return false.

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big: big, medium: medium, small: small}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if p.big > 0 {
			p.big--
			return true
		}
	case 2:
		if p.medium > 0 {
			p.medium--
			return true
		}
	case 3:
		if p.small > 0 {
			p.small--
			return true
		}
	}
	return false
}

// MORE Q

// // Unit Test
// func testParkingSystem() {
// 	ps := Constructor(1, 1, 0)
// 	if ps.AddCar(1) != true {
// 		fmt.Println("Test failed: AddCar(1)")
// 	}
// 	if ps.AddCar(2) != true {
// 		fmt.Println("Test failed: AddCar(2)")
// 	}
// 	if ps.AddCar(3) != false {
// 		fmt.Println("Test failed: AddCar(3)")
// 	}
// 	fmt.Println("All tests passed for ParkingSystem")
// }

// // Problem:
// // Design a simple elevator system with the following functionalities:
// // - AddPerson(weight int) bool: Adds a person to the elevator if the total weight doesn't exceed the max weight.
// // - RemovePerson(weight int) bool: Removes a person from the elevator.
// // - CurrentFloor() int: Returns the current floor of the elevator.
// //
// // Implement a struct BasicElevator with a maximum weight limit and floor tracking.

// type Elevator interface {
// 	AddPerson(weight int) bool
// 	RemovePerson(weight int) bool
// 	CurrentFloor() int
// }

// type BasicElevator struct {
// 	maxWeight   int
// 	currentLoad int
// 	floor       int
// }

// func (e *BasicElevator) AddPerson(weight int) bool {
// 	if e.currentLoad+weight <= e.maxWeight {
// 		e.currentLoad += weight
// 		return true
// 	}
// 	return false
// }

// func (e *BasicElevator) RemovePerson(weight int) bool {
// 	if e.currentLoad >= weight {
// 		e.currentLoad -= weight
// 		return true
// 	}
// 	return false
// }

// func (e *BasicElevator) CurrentFloor() int {
// 	return e.floor
// }

// // Unit Test
// func testBasicElevator() {
// 	elevator := &BasicElevator{maxWeight: 500}
// 	if !elevator.AddPerson(100) {
// 		fmt.Println("Test failed: AddPerson(100)")
// 	}
// 	if !elevator.AddPerson(400) {
// 		fmt.Println("Test failed: AddPerson(400)")
// 	}
// 	if elevator.AddPerson(10) {
// 		fmt.Println("Test failed: AddPerson(10) should exceed max weight")
// 	}
// 	if !elevator.RemovePerson(100) {
// 		fmt.Println("Test failed: RemovePerson(100)")
// 	}
// 	if elevator.currentLoad != 400 {
// 		fmt.Println("Test failed: currentLoad should be 400")
// 	}
// 	fmt.Println("All tests passed for BasicElevator")
// }

// // Problem:
// // Design a logger that receives a string message along with a timestamp in seconds.
// // - A message should only be printed if it hasn't been printed in the last 10 seconds.
// // - Implement the Logger struct with the following methods:
// //   * Constructor() Logger: Initializes the logger object.
// //   * ShouldPrintMessage(timestamp int, message string) bool: Returns true if the message should be printed.

// type Logger struct {
// 	messageTimestamps map[string]int
// }

// func Constructor() Logger {
// 	return Logger{messageTimestamps: make(map[string]int)}
// }

// func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
// 	if ts, exists := l.messageTimestamps[message]; !exists || timestamp-ts >= 10 {
// 		l.messageTimestamps[message] = timestamp
// 		return true
// 	}
// 	return false
// }

// // Unit Test
// func testLogger() {
// 	logger := Constructor()
// 	if !logger.ShouldPrintMessage(1, "hello") {
// 		fmt.Println("Test failed: ShouldPrintMessage(1, 'hello')")
// 	}
// 	if logger.ShouldPrintMessage(2, "hello") {
// 		fmt.Println("Test failed: ShouldPrintMessage(2, 'hello')")
// 	}
// 	if !logger.ShouldPrintMessage(11, "hello") {
// 		fmt.Println("Test failed: ShouldPrintMessage(11, 'hello')")
// 	}
// 	fmt.Println("All tests passed for Logger")
// }

// // Problem:
// // You are given an array of k linked lists, each sorted in ascending order.
// // Merge all the linked lists into one sorted linked list and return it.
// //
// // Implement a function mergeKLists(lists []*ListNode) *ListNode to solve the problem.

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	}

// 	// Use a priority queue (min-heap) to keep track of the smallest node
// 	h := &MinHeap{}
// 	heap.Init(h)

// 	// Add the head of each list to the heap
// 	for _, list := range lists {
// 		if list != nil {
// 			heap.Push(h, list)
// 		}
// 	}

// 	dummy := &ListNode{}
// 	current := dummy

// 	// Continuously pop the smallest node and add its next node to the heap
// 	for h.Len() > 0 {
// 		smallest := heap.Pop(h).(*ListNode)
// 		current.Next = smallest
// 		current = smallest
// 		if smallest.Next != nil {
// 			heap.Push(h, smallest.Next)
// 		}
// 	}

// 	return dummy.Next
// }

// // MinHeap is a custom implementation of a priority queue for ListNode
// type MinHeap []*ListNode

// func (h MinHeap) Len() int            { return len(h) }
// func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
// func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// // Unit Test
// func testMergeKLists() {
// 	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
// 	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
// 	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

// 	lists := []*ListNode{list1, list2, list3}
// 	mergedList := mergeKLists(lists)

// 	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}
// 	for i := 0; mergedList != nil; i++ {
// 		if mergedList.Val != expected[i] {
// 			fmt.Println("Test failed: mergeKLists")
// 			return
// 		}
// 		mergedList = mergedList.Next
// 	}
// 	fmt.Println("All tests passed for mergeKLists")
// }

// // Problem:
// // Design a simple in-memory file system with the following methods:
// // - CreateFile(path string, content string) bool: Creates a new file at the given path with the specified content.
// // - ReadFile(path string) string: Returns the content of the file at the given path.
// // - ListDirectory(path string) []string: Lists all files and directories under the given path.
// //
// // Implement a struct SimpleFileSystem to solve the problem.

// type FileSystem interface {
// 	CreateFile(path string, content string) bool
// 	ReadFile(path string) string
// 	ListDirectory(path string) []string
// }

// type SimpleFileSystem struct {
// 	files map[string]string
// }

// func NewFileSystem() *SimpleFileSystem {
// 	return &SimpleFileSystem{files: make(map[string]string)}
// }

// func (fs *SimpleFileSystem) CreateFile(path string, content string) bool {
// 	if _, exists := fs.files[path]; exists {
// 		return false // File already exists
// 	}
// 	fs.files[path] = content
// 	return true
// }

// func (fs *SimpleFileSystem) ReadFile(path string) string {
// 	if content, exists := fs.files[path]; exists {
// 		return content
// 	}
// 	return ""
// }

// func (fs *SimpleFileSystem) ListDirectory(path string) []string {
// 	var result []string
// 	for filePath := range fs.files {
// 		if strings.HasPrefix(filePath, path) {
// 			result = append(result, filePath)
// 		}
// 	}
// 	return result
// }

// // Unit Test
// func testFileSystem() {
// 	fs := NewFileSystem()
// 	if !fs.CreateFile("/a/b/c", "hello") {
// 		fmt.Println("Test failed: CreateFile(/a/b/c, 'hello')")
// 	}
// 	if fs.CreateFile("/a/b/c", "world") {
// 		fmt.Println("Test failed: CreateFile(/a/b/c, 'world') should not overwrite existing file")
// 	}
// 	if content := fs.ReadFile("/a/b/c"); content != "hello" {
// 		fmt.Printf("Test failed: ReadFile(/a/b/c), got: %s, expected: 'hello'\n", content)
// 	}
// 	fs.CreateFile("/a/b/d", "world")
// 	if dirs := fs.ListDirectory("/a/b"); len(dirs) != 2 {
// 		fmt.Println("Test failed: ListDirectory(/a/b), expected 2 entries")
// 	}
// 	fmt.Println("All tests passed for FileSystem")
// }

// // Problem:
// // Design a PlaylistManager with the following methods:
// // - AddSong(song string) bool: Adds a new song to the playlist. Returns false if the song already exists.
// // - RemoveSong(song string) bool: Removes a song from the playlist.
// // - GetNextSong() string: Returns the next song in the playlist in a circular manner.

// type PlaylistManager struct {
// 	songs []string
// 	index int
// }

// func NewPlaylistManager() *PlaylistManager {
// 	return &PlaylistManager{songs: []string{}, index: -1}
// }

// func (pm *PlaylistManager) AddSong(song string) bool {
// 	for _, s := range pm.songs {
// 		if s == song {
// 			return false // Song already exists
// 		}
// 	}
// 	pm.songs = append(pm.songs, song)
// 	if pm.index == -1 {
// 		pm.index = 0
// 	}
// 	return true
// }

// func (pm *PlaylistManager) RemoveSong(song string) bool {
// 	for i, s := range pm.songs {
// 		if s == song {
// 			pm.songs = append(pm.songs[:i], pm.songs[i+1:]...)
// 			if pm.index >= len(pm.songs) {
// 				pm.index = 0
// 			}
// 			return true
// 		}
// 	}
// 	return false
// }

// func (pm *PlaylistManager) GetNextSong() string {
// 	if len(pm.songs) == 0 {
// 		return ""
// 	}
// 	song := pm.songs[pm.index]
// 	pm.index = (pm.index + 1) % len(pm.songs)
// 	return song
// }

// // Unit Test
// func testPlaylistManager() {
// 	pm := NewPlaylistManager()
// 	if !pm.AddSong("Song A") {
// 		fmt.Println("Test failed: AddSong('Song A')")
// 	}
// 	if !pm.AddSong("Song B") {
// 		fmt.Println("Test failed: AddSong('Song B')")
// 	}
// 	if pm.AddSong("Song A") {
// 		fmt.Println("Test failed: AddSong('Song A') should return false")
// 	}
// 	if song := pm.GetNextSong(); song != "Song A" {
// 		fmt.Printf("Test failed: GetNextSong(), got: %s, expected: 'Song A'\n", song)
// 	}
// 	if !pm.RemoveSong("Song A") {
// 		fmt.Println("Test failed: RemoveSong('Song A')")
// 	}
// 	if song := pm.GetNextSong(); song != "Song B" {
// 		fmt.Printf("Test failed: GetNextSong(), got: %s, expected: 'Song B'\n", song)
// 	}
// 	fmt.Println("All tests passed for PlaylistManager")
// }

// // Problem:
// // Design a simple calculator that supports addition, subtraction, multiplication, and division.
// // Implement a struct Calculator with the following methods:
// // - Add(a, b int) int
// // - Subtract(a, b int) int
// // - Multiply(a, b int) int
// // - Divide(a, b int) (int, error): Returns an error if division by zero is attempted.

// type Calculator struct{}

// func (c *Calculator) Add(a, b int) int {
// 	return a + b
// }

// func (c *Calculator) Subtract(a, b int) int {
// 	return a - b
// }

// func (c *Calculator) Multiply(a, b int) int {
// 	return a * b
// }

// func (c *Calculator) Divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("division by zero")
// 	}
// 	return a / b, nil
// }

// // Unit Test
// func testCalculator() {
// 	calc := &Calculator{}
// 	if result := calc.Add(2, 3); result != 5 {
// 		fmt.Println("Test failed: Add(2, 3)")
// 	}
// 	if result := calc.Subtract(5, 3); result != 2 {
// 		fmt.Println("Test failed: Subtract(5, 3)")
// 	}
// 	if result := calc.Multiply(4, 3); result != 12 {
// 		fmt.Println("Test failed: Multiply(4, 3)")
// 	}
// 	if result, err := calc.Divide(10, 2); err != nil || result != 5 {
// 		fmt.Println("Test failed: Divide(10, 2)")
// 	}
// 	if _, err := calc.Divide(10, 0); err == nil {
// 		fmt.Println("Test failed: Divide(10, 0) should return error")
// 	}
// 	fmt.Println("All tests passed for Calculator")
// }
