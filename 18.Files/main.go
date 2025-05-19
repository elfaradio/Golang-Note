package main

import (
	"fmt"
	"io"
	"os"
)

//Farhadul Islam CSE-56 IIUC

func main() {

	txt := "helllo"
	file, err := os.Create("./mygreetings.txt") // creating

	checkNilerr(err) // check error

	io.WriteString(file, txt) // file e ja likhbo

	defer file.Close() // must

	readFile("./mygreetings.txt") //file theke read korbo

}

func readFile(fileName string) {

	//data, err := ioutil.ReadFile(fileName) // file read kora  old version

	data, err := os.ReadFile(fileName) // new version

	checkNilerr(err)

	fmt.Println(data) // 104,102,......

	fmt.Println(string(data)) // main data

}
func checkNilerr(err error) {

	if err != nil {
		panic(err)
	}
}
