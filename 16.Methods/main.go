package main

import "fmt"

//Farhadul Islam CSE-56 IIUC
type User struct {
	name string
	age  int
}

func main() {

	v := User{"farhad", 2}
	fmt.Print(v, "\n")

	v.Getage() // for printing age or getter

	v.SetName() // doesnt change the main value

	fmt.Print(v.name, "\n") //no change

}

func (u User) Getage() { // getter

	fmt.Print(u.age, "\n")

}

func (u User) SetName() { // pass by value not changing main value  ... pointer can change main value

	//u.name = "rakib"
}
