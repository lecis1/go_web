package gobasic

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	speciality string
}

func (self *Person) PersonInfo() {
	fmt.Printf("The person's name is %s\n", self.Name)
	fmt.Printf("%s's age is %d", self.Name, self.Age)
}

func StudentInfo() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His Speciality is ", mark.speciality)

	fmt.Println("Mark become old")
	mark.Human.age = 46
	fmt.Println("His age is ", mark.age)

	fmt.Println("Mark is not an athlet anymore")
	mark.Human.weight += 60
	fmt.Println("His weight is", mark.weight)
}
