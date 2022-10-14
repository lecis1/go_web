package gobasic

import (
	"fmt"
	"strconv"
)

type Human2 struct {
	Name  string
	Age   int
	Phone string
}

type Student2 struct {
	Human2 //匿名字段
	School string
	Loan   float32
}

type Employee struct {
	Human2  //匿名字段
	Company string
	Money   float32
}

type Element interface{}
type List []Element

type Person2 struct {
	Name string
	Age  int
}

func (h *Human2) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}

func (h *Human2) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func (h *Human2) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, Call me on %s\n", e.Name, e.Company, e.Phone)
}

func (s *Student2) BorrowMoney(amount float32) {
	s.Loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
	e.Money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	// Guzzle(beerStein string)
}

// type YoungChap interface {
// 	SayHi()
// 	Sing(song string)
// 	BorrowMoney(amount float32)
// }

// type ElderlyGent interface {
// 	SayHi()
// 	Sing(song string)
// 	SpendSalary(amount float32)
// }

func InterfaceTest() {
	mike := Student2{Human2{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student2{Human2{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human2{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human2{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	var i Men = &mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = &tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = &paul, &sam, &mike
	for _, value := range x {
		value.SayHi()
	}

}

func (h *Human2) String() string {
	return "(" + h.Name + " - " + strconv.Itoa(h.Age) + " years - ✆ " + h.Phone + ")"
}

func StringTest() {
	Bob := &Human2{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}

func (p *Person2) String() string {
	return "(name: " + p.Name + " - age: " + strconv.Itoa(p.Age) + " years)"
}

func TypeElement() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person2{"Dennis", 70}

	// for index, element := range list {
	// 	if value, ok := element.(int); ok {
	// 		fmt.Printf("list[%d] is a int and its value is %d\n", index, value)
	// 	} else if value, ok := element.(string); ok {
	// 		fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	// 	} else if value, ok := element.(Person2); ok {
	// 		fmt.Printf("list[%d] is a Person2 and its value is %s\n", index, &value)
	// 	} else {
	// 		fmt.Printf("list[%d] is of a defferent type\n", index)
	// 	}
	// }

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is a int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person2:
			fmt.Printf("list[%d] is a Person2 and its value is %s\n", index, &value)
		default:
			fmt.Printf("list[%d] is of a defferent type\n", index)
		}
	}
}
