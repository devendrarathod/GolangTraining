package main

import "fmt"
import "strconv"
import "math"

func readFloat(a *float64) {
	_, err := fmt.Scanf("%f", a)
	if err != nil {
		*a = 0
	}
}

func main() {
	fmt.Print("\tHello World!\n")
	s := "Strings are character arrays."
	fmt.Println(string(s[2]), len(s))
	sb := "words."
	fmt.Println(s[0:11], sb)

	ab, _ := strconv.Atoi("11")
	ba := strconv.Itoa(12)
	fmt.Println(ab, ba)

	ia := 20
	var ib float64 = 40.443
	fmt.Println(float64(ia), int(ib))

	fmt.Println(string([]byte{'b', 'o', 'r', 'i', 'n', 'g', ' ', 's', 't', 'r', 'i', 'n', 'g'}))
	fmt.Println([]byte("Hello"))

	ib = 0
	fmt.Println("Input float:")
	readFloat(&ib)

	fmt.Println(math.Ceil(ib))
}
