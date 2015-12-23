package main

import "fmt"

func sumInts(a ...int) int {
	sOut := 0
	for _, i := range a {
		sOut += i
	}
	return sOut
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	sum := sumInts(s...)
	fmt.Println(sum)

	fmt.Println("------------------")
	fmt.Print("Give me integer: ")
	var lnIn int
	fmt.Scanln(&lnIn)
	lnIn = lnIn * 2
	fmt.Println("Double ", lnIn)
}
