package main

import "fmt"

//Farhadul Islam CSE-56 IIUC
// defer kono line or function ke sobar last e execute korte bole

// ekadik defer thakle LIFO style

func main() {

	defer fmt.Print("Hello")

	fmt.Println("Farhad")

	//Farhad Hello

	nydef()

	// Farhad then 43210 then hello bcz last in fast out if different defer ...so je last defer se age print hobe

}

func nydef() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}

}
