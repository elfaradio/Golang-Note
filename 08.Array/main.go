package main

//Farhadul Islam CSE-56 IIUC

import "fmt"

func main() {

	// not much use of array in golang

	var arr [10]int // or var arr[4]int{1,2,3,4}

	//fmt.Printf("%T", arr)[10]int

	fmt.Print(arr, "\n") /// 0,0,0,0....

	arr[0] = 2

	fmt.Print(arr, "\n") // 2,0,0....

	fmt.Print(len(arr)) //10 not 1
}
