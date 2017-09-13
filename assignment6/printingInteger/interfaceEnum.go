package main

import "strconv"
import "fmt"

type newInt int

const (
	val1 newInt = 1
)

func (num newInt) String() string {
	return strconv.Itoa(int(num))
}

func main() {
	var i newInt = 14343432432
	fmt.Printf("%s", i)
}
