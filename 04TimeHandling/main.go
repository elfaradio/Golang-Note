package main

//Farhadul Islam CSE-56 IIUC

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // present time

	fmt.Println(t.Date()) // date y-m-d

	fmt.Println(t.Format("01-02-2006 15:04:05 Monday "))         // must be used d-m-y-time-day....
	nt := time.Date(2026, time.May, 10, 23, 25, 12, 1, time.UTC) // nijer moto kore date create
	fmt.Println(nt)

}
