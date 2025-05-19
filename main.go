package main

// Farhadul Islam CSE-56

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt() int {
	input, _ := reader.ReadString('\n')
	val, _ := strconv.Atoi(strings.TrimSpace(input))
	return val
}

func readInts() []int {
	input, _ := reader.ReadString('\n')
	str := strings.Fields(input)
	nums := make([]int, len(str))
	for i, p := range str {
		nums[i], _ = strconv.Atoi(p)
	}
	return nums
}

func print(a ...interface{}) {
	fmt.Fprint(writer, a...)
}

func println(a ...interface{}) {
	fmt.Fprintln(writer, a...)
}

func main() {
	defer writer.Flush()

	t := readInt()
	for i := 0; i < t; i++ {

	}
}
