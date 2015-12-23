package main

import "fmt"
import "strconv"

func readInt(a *int) {
	_, err := fmt.Scanf("%d", a)
	if err != nil {
		*a = 0
	}
}

func main() {
	a, b := 0, 0
	fmt.Print("Give me first numb:")
	readInt(&a)
	fmt.Print("Give me second numb:")
	readInt(&b)
	fmt.Println("" + strconv.Itoa(a) + "%" + strconv.Itoa(b) + "=" + strconv.Itoa(a%b))

	for i := 1; i < 1000; i += 1 {
		if i%2 == 0 {
			fmt.Print(strconv.Itoa(i) + ",")
		}
	}
	fmt.Println("")

	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("")

	sum, itr := 0, 0
	for itr < 1000 {
		if itr%3 == 0 || itr%5 == 0 {
			sum += itr
		}
		itr++
	}
	fmt.Println(sum)
}
