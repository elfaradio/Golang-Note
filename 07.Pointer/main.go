package main

//Farhadul Islam CSE-56 IIUC

import "fmt"

func main() {

	// all time value passing copy noy actual value
	// using pointer we can pass the real value

	var ptr *int // i am val pointer

	fmt.Println(ptr) // default nil

	val := 23

	fmt.Print(&val, val, "\n") // same to aptr &-address,val-val

	var aptr = &val // aptr=val er adress

	fmt.Print(aptr, "\n", *aptr, "\n") // same to address of val and value of val

	*aptr = *aptr * 2 // multiply by 2 which will change the value of val and aptr cz both share the same address

	fmt.Print(val) // 23*2

	// pointer operate in actual value not in copy value

}
