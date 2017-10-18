package main

import (
	"fmt"
	"strconv"
)

func vals() (int, int) {
	return 3, 7
}

func print(str string) {
	fmt.Println(str)
}

func toString(i int) string {
	return strconv.Itoa(i)
}

func main() {
	a, b := vals()

	print(toString(a))
	print(toString(b))

	_, c := vals()
	print(toString(c))
}
