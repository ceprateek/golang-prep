package hash

import (
	"crypto/sha256"
	"fmt"
)

type Person struct {
	Name    string
	Surname string
	Age     int
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func RunSample() {
	person1 := Person{
		Name:    "John",
		Surname: "Smith",
		Age:     18,
	}
	fmt.Println("------------- Person1's hash")
	fmt.Println(asSha256(person1))

	person2 := Person{
		Name:    "John",
		Surname: "Smith",
		Age:     22,
	}
	fmt.Println("------------- Person2's hash")
	fmt.Println(asSha256(person2))

	person3 := Person{
		Name:    "John",
		Surname: "Smith",
		Age:     18,
	}
	fmt.Println("------------- person3's hash")
	fmt.Println(asSha256(person3))

	person4 := Person{
		Name:    "John",
		Surname: "Smith",
		Age:     22,
	}
	fmt.Println("------------- person4's hash")
	fmt.Println(asSha256(person4))
}
