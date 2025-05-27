package main

import "fmt"

//Farhadul Islam CSE-56 IIUC
// defer kono line or function ke sobar last e execute korte bole

// ekadik defer thakle LIFO style

func main() {

	//defer fmt.Print("Hello")

	//fmt.Println("Farhad")

	//Farhad Hello

	nydef()

	x := sum()
	fmt.Println(x)

	y := sum1()
	fmt.Println(y)

	z := sum2()
	fmt.Println(z)
	// Farhad then 43210 then hello bcz last in fast out if different defer ...so je last defer se age print hobe

}

func sum() int {

	ans := 0
	fmt.Println(ans)
	ans = 5
	fmt.Println(ans)
	return ans // return 5
}

func sum1() int {
	ans := 0
	show := func() {
		ans = 5

	}
	defer show()
	return ans // not returnting 5 returnting 0 but why?? See in Name_Return_Value file and wath defer in go with habib channel

}

func sum2() (s int) {

	show := func() {
		s = 5
	}
	defer show()

	return // returnning 5
}
func nydef() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}

}
