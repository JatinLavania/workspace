// this is go file
package main

import "fmt"

func main() {

	var sum int

	num1 := 10
	num2 := 25

	sum = num1 + num2

	fmt.Println("Hello jatin 1", sum)
	var num11, num22 int
	fmt.Print("Enter the First Number = ")
	fmt.Scanln(&num11)

	fmt.Print("Enter the Second Number = ")
	fmt.Scanln(&num22)

	sum += num11 + num22

	fmt.Println("Hello jatin 2", sum)
}
