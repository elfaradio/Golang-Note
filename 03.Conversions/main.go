package main

//Farhadul Islam CSE-56 IIUC

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	convert()

}

func convert() {
	reader := bufio.NewReader(os.Stdin) // reading or input

	input, _ := reader.ReadString('\n') // read until '\n'  string input nibe

	fmt.Println("enter ")

	//ans := input + 1 // err cz no str to int,so we have to convert

	//ans, err := strconv.ParseFloat(input, 64) // wrong bcz error  -->strconv.ParseFloat: parsing"4\n" extra '\n' ,so (str+int)->>wrong
	ans, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //right TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.

	if err != nil {
		fmt.Println(err, ans) // an error
	} else {
		ans++
	}

	fmt.Println(ans, input)
}
