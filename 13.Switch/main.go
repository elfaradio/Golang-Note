package main

import "fmt"

//Farhadul Islam CSE-56 IIUC
func main() {

	s := "Jan"
	switch s {
	case "Jan":
		fmt.Print("Jan", "\n")
		fallthrough // if we want to next case executed
	case "Feb":
		fmt.Print("11\n")
	default:
		fmt.Print(2121892)

	}

}
