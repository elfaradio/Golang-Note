package main

import "fmt"

//Farhadul Islam CSE-56 IIUC

type User struct {
	Name   string
	age    int
	status bool
}

func main() {

	// no class ,inheritence,super,parent in golang

	a := User{"fg", 12, true}

	fmt.Println(a)

	//	another way

	// a.age
	// a.status
	// a.Name

}
