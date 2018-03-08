package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	rollNo    int
	firstName string
	lastName  string
}

func (p Person) fullName() string {
	return p.firstName + p.lastName
}

func (p Person) greeting() {
	fmt.Println("Greeting a Person")
}

type Licenses struct {
	Person
	canKill bool
}

func (l Licenses) greeting() {
	fmt.Println("Greets - Calling from Licenses")
}

type newPerson struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := Person{1, "Ankur", "Raina"}
	p2 := Person{2, "Ankit", "Raina"}
	fmt.Println(p1.firstName, p1.lastName, p1.rollNo)
	fmt.Println(p2.firstName, p2.lastName, p2.rollNo)
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	p3 := Licenses{
		Person: Person{
			rollNo:    3,
			lastName:  "Raina",
			firstName: "Ankur",
		},
		canKill: true,
	}
	fmt.Println(p3.fullName())
	fmt.Println(p3.canKill, p3.Person)
	p3.greeting()
	p3.Person.greeting()
	p4 := newPerson{"Hakka", "Hulla", 20}
	bs, _ := json.Marshal(p4)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
	var p5 newPerson
	println(p5.FirstName, "---", p5.LastName, "---", p5.Age)

	bs1 := []byte(`{"FirstName":"Will", "LastName":"Smith", "Age":50}`)
	json.Unmarshal(bs1, &p5)
	println(p5.FirstName, "---", p5.LastName, "---", p5.Age)
}
