package main

import "fmt"

// Creating addition funciton
func add(a int, b int) {
	fmt.Println("Addition of 2 numbers ", a+b)
}

// Creating Subtraction function
func subtract(a int, b int) {
	fmt.Println("Subtraction of 2 numbers ", a-b)
}

// Creating Multiplication function
func multiply(a int, b int) {
	fmt.Println("Multiplication of 2 numbers ", a*b)
}

// Creating Divison function
func divide(a int, b int) {
	fmt.Println("Division of 2 numbers ", a/b)
}

//Creating mod(remainder) function
func mod(a int, b int) {
	fmt.Println("Mod of 2 numbers ", a%b)
}

func main() {
	var a int
	var b int
	var c int

	fmt.Printf("Please enter first number: ")
	fmt.Scanln(&a)

	fmt.Printf("Please enter Second number: ")
	fmt.Scanln(&b)

	fmt.Printf("Please select the operation : \n 1 add \n 2 subtract \n 3 multiply \n 4 divide \n 5 mod \n")
	fmt.Scanln(&c)

	if c == 1 {

		add(a, b)
	} else if c == 2 {

		subtract(a, b)
	} else if c == 3 {

		multiply(a, b)
	} else if c == 4 {

		divide(a, b)
	} else if c == 5 {

		mod(a, b)
	} else {
		fmt.Println("Invalid Input")
	}
}
