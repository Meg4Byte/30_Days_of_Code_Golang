package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if initialAge < 0 {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
	} else {
		p.age = initialAge
	}
	return p
}

func (p *person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if 13 <= p.age && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}
}

func (p *person) yearPasses() person {
	//Increment the age of the person in here
	p.age += 1
	return *p
}

func main() {
	// init struct use {age: 100 }
	p := person{age: 100}
	p = p.NewPerson(100)
	fmt.Println("You are", p.age, "years age")
	p.amIOld()
	p.yearPasses()
	fmt.Println("You are", p.age, "years age") //101
	p.yearPasses()
	fmt.Println("You are", p.age, "years age") //102
	p.yearPasses()
	fmt.Println("You are", p.age, "years age") //103

}
