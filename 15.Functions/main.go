package main

import "fmt"

//Farhadul Islam CSE-56 IIUC
func main() {

	// cant be declared inside a function

	// func a{  wrong

	// }

	x := sum(1, 2)

	y, _ := arrsum(1, 2, 3, 3)
	fmt.Print(x, "\n")

	fmt.Print(y, "\n")

}

func arrsum(a ...int) (int, string) { /// three or more parameters or slice array and two returns

	s := 0
	for _, i := range a {
		s += i
	}
	return s, "hello" // return 2 value

}
func sum(a, b int) int {

	fmt.Print("sum\n")
	return a + b
}

func init() {

	fmt.Print("inti\n") // no matter what it will be called firts
}
