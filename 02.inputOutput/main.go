package main

//Farhadul Islam CSE-56 IIUC

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// lexer is the grammar

	inputOutput() // comma ok syntax,bufio

}

func inputOutput() {

	reader := bufio.NewReader(os.Stdin) // read
	// Comma ok syntax cz no try catch

	fmt.Println("enter val")

	input, _ := reader.ReadString('\n') // read str until '\n' // (_) means error// input-try,err->catch  if err!=nil then work

	fmt.Println("hello", input) // type string %T

}
