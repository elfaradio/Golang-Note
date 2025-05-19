package main

import "fmt"

//Farhadul Islam CSE-56 IIUC
func main() {

	for i := 0; i < 10; i += 1 { // 1 to 10
		fmt.Print(i+1, "\n")
	}
	arr := []int{1, 2, 3, 3}

	for i, j := range arr { // i=idx,j=value
		fmt.Print(i, j, "\n")
	}

	for i := 0; i < len(arr); i++ { // another way
		fmt.Print(arr[i])

	}
	cnt := 1
	for cnt < 10 { // another way or while
		if cnt == 2 {
			goto bro
		}
		if cnt == 5 {
			cnt++ // must be increase .... if not then infinity loop

			continue
		} else if cnt == 8 {
			break
		}
		fmt.Print(cnt, "\n")
		cnt++
	}
bro:
	fmt.Print("Bro", "\n")

}
