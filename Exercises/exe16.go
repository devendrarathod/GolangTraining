package main

import (
	"fmt"
	"strconv"
)

func divMod2(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func maxList(n ...int) (m int) {
	for _, i := range n {
		if i > m {
			m = i
		}
	}
	return
}

func twoParams(name string, age int) string {
	astr := strconv.Itoa(age)
	return name + " is " + astr + " years old."
}

func twoReturns(age int) (int, bool) {
	return age * 7, age*7 > 25
}

func dogage(age int) (dogYears int) {
	dogYears = age * 7
	return
}

func maxListCallback(c func(int), n ...int) {
	var m int
	for _, i := range n {
		if i > m {
			m = i
		}
	}
	c(m)
	return
}
func printInt(a int) {
	fmt.Println(a)
}

func main() {
	fmt.Println(divMod2(5))
	fmt.Println(maxList(1, 2, 4, 5, 7, 3, 6, 22, 6, 4, 8, 1))
	fmt.Println("Bool:", ((true && false) || (false && true) || !(false && true)))
	fmt.Println("two params:", twoParams("John", 27))

	dage, dold := twoReturns(10)
	if dold {
		fmt.Println("John is ", dage, "in dog years and is old")
	} else {
		fmt.Println("John is ", dage, "in dog years and is not old")
	}

	fmt.Println(dogage(7))

	maxListCallback(printInt, 4, 7, 2, 6, 4, 8, 4, 1)
}
