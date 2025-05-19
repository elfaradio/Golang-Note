package main

import (
	"fmt"
	"sort"
)

//Farhadul Islam CSE-56 IIUC

func main() {

	//  Popular ,under the array,powerful,more features

	var slc = []int{}

	fmt.Print(slc, "\n") //[]

	fmt.Printf("%T\n", slc) // int[]

	var slc1 = []int{1, 3, 5, 7}

	fmt.Print(slc1, "\n")

	fmt.Println(len(slc), len(slc1)) // 0,4 length of slice

	slc1 = append(slc1, 2) // 2 add also append(slc1,2,3,4..)

	fmt.Println(len(slc), len(slc1)) // 0,5 ... increase by 1 bcz of append

	//slc1 = append(slc1[1:3]) // give 1-2 pos and remove baki element

	fmt.Println(len(slc1)) // after removing 2

	slc1 = append(slc1[1:3], 2, 8, 9) //give 1-2 pos and remove baki element  and then add 2,8,9

	fmt.Print(slc1, "\n")

	fmt.Println(len(slc1)) // 5

	scl2 := make([]int, 6) // another way using make()

	scl2[5] = 5

	fmt.Println(scl2, len(scl2)) // 6 length

	scl2 = append(scl2, 2, 3, 4)

	fmt.Println(scl2, len(scl2)) // after append length increase even init was 6

	// func (a SortBy) Len() int           { return len(a) }
	// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	// func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

	//sort.Ints(scl2) // sort the array

	sort.IntsAreSorted(scl2) // true if sorted

	// how to remove a value from slice based on index

	dx := 7

	scl2 = append(scl2[:dx], scl2[dx+1:]...) // remove 7th index

	fmt.Println(scl2)

}
