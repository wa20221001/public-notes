package main

import "fmt"

// Exercise 1: Define a struct for a Book with fields Title, Author, and Pages
// Operation: Define and use a struct
type Book struct {
	// TODO: Add fields Title (string), Author (string), and Pages (int)
}

func describeBook(b Book) string {
	// TODO: Return a description of the book in the format "Title by Author, Pages pages"
	return ""
}

func testDescribeBook() {
	b := Book{Title: "1984", Author: "George Orwell", Pages: 328}
	expected := "1984 by George Orwell, 328 pages"
	result := describeBook(b)
	if result == expected {
		fmt.Println("testDescribeBook passed")
	} else {
		fmt.Printf("testDescribeBook failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 2: Define a struct for a Point with fields X and Y
// Operation: Struct with basic arithmetic
type Point struct {
	// TODO: Add fields X and Y (int)
}

func addPoints(p1, p2 Point) Point {
	// TODO: Return a new Point where X and Y are the sums of p1 and p2's X and Y
	return Point{}
}

func testAddPoints() {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 3, Y: 4}
	expected := Point{X: 4, Y: 6}
	result := addPoints(p1, p2)
	if result == expected {
		fmt.Println("testAddPoints passed")
	} else {
		fmt.Printf("testAddPoints failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 3: Create a method to move a point by given dx and dy
// Operation: Methods on structs
func (p *Point) move(dx, dy int) {
	// TODO: Update the Point by adding dx to X and dy to Y
}

func testMovePoint() {
	p := Point{X: 5, Y: 5}
	p.move(3, 4)
	expected := Point{X: 8, Y: 9}
	if p == expected {
		fmt.Println("testMovePoint passed")
	} else {
		fmt.Printf("testMovePoint failed: expected %v, got %v\n", expected, p)
	}
}

// Exercise 4: Define a struct for a Car with fields Brand, Model, and Speed
// Operation: Struct with methods and fields
type Car struct {
	// TODO: Add fields Brand (string), Model (string), and Speed (int)
}

// Method to increase the speed of the car by a given amount
func (c *Car) accelerate(amount int) {
	// TODO: Increase the Speed by amount
}

// Method to decrease the speed of the car by a given amount
func (c *Car) brake(amount int) {
	// TODO: Decrease the Speed by amount, ensuring it doesn't go below 0
}

func testCar() {
	c := Car{Brand: "Tesla", Model: "Model 3", Speed: 0}
	c.accelerate(50)
	c.brake(20)
	expectedSpeed := 30
	if c.Speed == expectedSpeed {
		fmt.Println("testCar passed")
	} else {
		fmt.Printf("testCar failed: expected speed %v, got %v\n", expectedSpeed, c.Speed)
	}
}

// Exercise 5: Define an interface for a Drawable shape with a method draw()
// Operation: Define and use an interface
type Drawable interface {
	draw() string
}

// Implement Drawable for a Circle struct
type Circle struct {
	Radius int
}

func (c Circle) draw() string {
	// TODO: Return "Drawing a circle with radius <Radius>"
	return ""
}

func testDrawCircle() {
	c := Circle{Radius: 5}
	expected := "Drawing a circle with radius 5"
	result := c.draw()
	if result == expected {
		fmt.Println("testDrawCircle passed")
	} else {
		fmt.Printf("testDrawCircle failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 6: Implement Drawable for a Rectangle struct
// Operation: Interface implementation for a different struct
type Rectangle struct {
	Width, Height int
}

func (r Rectangle) draw() string {
	// TODO: Return "Drawing a rectangle with width <Width> and height <Height>"
	return ""
}

func testDrawRectangle() {
	r := Rectangle{Width: 4, Height: 7}
	expected := "Drawing a rectangle with width 4 and height 7"
	result := r.draw()
	if result == expected {
		fmt.Println("testDrawRectangle passed")
	} else {
		fmt.Printf("testDrawRectangle failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 7: Create a function to print drawings of multiple Drawable shapes
// Operation: Polymorphism with interfaces
func printDrawings(shapes []Drawable) {
	// TODO: Iterate over the shapes and print the result of calling draw() on each
}

func testPrintDrawings() {
	shapes := []Drawable{
		Circle{Radius: 3},
		Rectangle{Width: 5, Height: 6},
	}
	fmt.Println("Expected:")
	fmt.Println("Drawing a circle with radius 3")
	fmt.Println("Drawing a rectangle with width 5 and height 6")
	fmt.Println("Result:")
	printDrawings(shapes)
}

// Exercise 8: Define a struct for a Student with fields Name and Grades (slice of int)
// Operation: Struct with a slice field
type Student struct {
	// TODO: Add fields Name (string) and Grades ([]int)
}

func (s Student) averageGrade() float64 {
	// TODO: Calculate and return the average of the grades
	return 0
}

func testAverageGrade() {
	s := Student{Name: "Alice", Grades: []int{90, 80, 70, 100}}
	expected := 85.0
	result := s.averageGrade()
	if result == expected {
		fmt.Println("testAverageGrade passed")
	} else {
		fmt.Printf("testAverageGrade failed: expected %v, got %v\n", expected, result)
	}
}

// Exercise 9: Define a struct for a BankAccount with fields AccountNumber and Balance
// Operation: Methods for deposit and withdraw
type BankAccount struct {
	// TODO: Add fields AccountNumber (string) and Balance (float64)
}

func (a *BankAccount) deposit(amount float64) {
	// TODO: Increase the Balance by amount
}

func (a *BankAccount) withdraw(amount float64) {
	// TODO: Decrease the Balance by amount, ensuring it doesn't go below 0
}

func testBankAccount() {
	account := BankAccount{AccountNumber: "12345", Balance: 1000.0}
	account.deposit(500)
	account.withdraw(300)
	expectedBalance := 1200.0
	if account.Balance == expectedBalance {
		fmt.Println("testBankAccount passed")
	} else {
		fmt.Printf("testBankAccount failed: expected balance %v, got %v\n", expectedBalance, account.Balance)
	}
}

// Exercise 10: Define a struct for an Employee with fields Name, Salary, and Position
// Operation: Basic struct usage and method
type Employee struct {
	// TODO: Add fields Name (string), Salary (float64), and Position (string)
}

func (e Employee) describe() string {
	// TODO: Return a description of the employee in the format "Name works as Position and earns Salary"
	return ""
}

func testEmployee() {
	e := Employee{Name: "Bob", Salary: 50000, Position: "Engineer"}
	expected := "Bob works as Engineer and earns 50000"
	result := e.describe()
	if result == expected {
		fmt.Println("testEmployee passed")
	} else {
		fmt.Printf("testEmployee failed: expected %v, got %v\n", expected, result)
	}
}

// Run all tests
func runTests() {
	testDescribeBook()
	testAddPoints()
	testMovePoint()
	testCar()
	testDrawCircle()
	testDrawRectangle()
	testPrintDrawings()
	testAverageGrade()
	testBankAccount()
	testEmployee()
}

func main() {
	runTests()
}
