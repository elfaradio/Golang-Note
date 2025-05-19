package main

import "fmt"

//Farhadul Islam CSE-56 IIUC
func main() {

	mp := make(map[string]int) // declare--- mp[key]val

	mp["Farhad"] = 2 // accesing

	fmt.Println(mp["Farhad"]) //2

	fmt.Println(mp) //run this

	//delete(mp, "Farhad") // deleting

	// iteration in map

	for i, j := range mp {
		fmt.Print(mp[i], j)

	}

}
