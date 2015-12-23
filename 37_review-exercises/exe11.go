package main

import "fmt"
import "strconv"

func ageMe(a *int) {
	*a += 1
}

func main() {
	age := 23

	fmt.Println("An age that exists is " + strconv.Itoa(age))
	fmt.Println(&age)

	var input string
	fmt.Print("Type me something reasonable: ")
	fmt.Scan(&input)
	fmt.Println("Hello " + input)

	ageMe(&age)
	fmt.Println("I changed my mind, your actually " + strconv.Itoa(age) + " years old.")
}
