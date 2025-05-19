package main

//Farhadul Islam CSE-56 IIUC

import "fmt"

const cn int = 10000 // cannot change

var val int = 200 // right global or public or can be changed

//val:=100 //wrong

func main() {
	variable()
}
func variable() {
	var s string = "Hello" // s:="hello"

	var ok bool = true

	var a int = 2 // int32,int64

	fmt.Printf("%T", s) //type

	var x uint = 256 // also many kinds

	var z bool // init false,0,empty str

	fmt.Print(s, ok, a, x, z, "\n")

}
