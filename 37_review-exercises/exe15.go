package main


import (
	"fmt"
	"strconv"
)

type thing interface{}
type Person struct {
	Age  int
	Name string
}

func main() {
	fmt.Printf("343 Binary: %b\n", 343)
	fmt.Printf("343 Hex   : %x\n", 343)
	h343 := fmt.Sprintf("%x", 343)
	fmt.Println("343 Hex string: ", h343)
	fmt.Printf("Type of h343: %T\n", h343)

	var a int
	fmt.Print("Enter an int: ")
	fmt.Scanf("%i", a)
	fmt.Println(a)

	var r rune
	r = 'a'
	fmt.Printf("Type of rune: %T\n", r)
	fmt.Println("Runes are int because of ASCII")

	fmt.Println("")
	for i := rune('!'); i < rune('z'); i += 1 {
		fmt.Printf("%c", i)
	}
	fmt.Println("")

	var stringVariableOfStaticKind string = "these are characters"

	fmt.Printf("Bytes of string: % x\n", stringVariableOfStaticKind)
	fmt.Printf("Length: %d\n", len(stringVariableOfStaticKind))
	fmt.Printf("First Rune: %c\n", stringVariableOfStaticKind[0])
	fmt.Printf("Slice of string: %s\n", stringVariableOfStaticKind[0:5])
	fmt.Printf("Concat of string: %s\n", stringVariableOfStaticKind[0:6]+"Concats are here.")

	convertedATOI, _ := strconv.Atoi("100")
	fmt.Printf("Atoi: %d\n", convertedATOI)
	convertedITOA := strconv.Itoa(100)
	fmt.Printf("Itoa: %s\n", convertedITOA)
	fmt.Printf("Int,float,back: %d\n", int(float64(10)))

	fmt.Println(string([]byte{'b', 'o', 'r', 'i', 'n', 'g', ' ', 's', 't', 'r', 'i', 'n', 'g'}))
	fmt.Println([]byte("Hello"))

	var thingVar thing
	thingVar = "Hello basic"
	fmt.Println(thingVar)

	sliceOfInts := []int{1, 2, 3, 4, 5}
	for _, i := range sliceOfInts {
		fmt.Println(i)
	}

	var dictionary map[string]string = make(map[string]string)
	dictionary["Hello"] = "Greeting used in the english language."
	delete(dictionary, "Hello")

	billy := Person{20, "Billy"}
	fmt.Printf("This person is %s who is age %d.\n", billy.Name, billy.Age)
}
