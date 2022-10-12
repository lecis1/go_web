package main

import (
	"fmt"
	"go_web_programming/src/gobasic"
	"os"
)

func main() {

	// a := 2

	// fmt.Printf("a = %d\n", a)

	// a1 := gobasic.Transferparameters(a)
	// a1 := gobasic.Transferpoint(&a)

	// fmt.Printf("a1 = %d\n", a1)
	// fmt.Printf("a = %d\n", a)

	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := gobasic.Filter(slice, gobasic.IsOdd)
	fmt.Println("odd elements of slice are: ", odd)
	even := gobasic.Filter(slice, gobasic.IsEven)
	fmt.Println("even elements of slice are: ", even)

	var user = os.Getenv("USER")

	fmt.Println("user = ", user)

	person := &gobasic.Person{
		Name: "Alex",
		Age:  26,
	}

	person.PersonInfo()

	gobasic.StudentInfo()

	gobasic.CalcBox()

}
