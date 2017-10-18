package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func print(i int) {
	fmt.Println(i)
}

func main() {
	nextInt := intSeq()

	print(nextInt())
	print(nextInt())
	print(nextInt())

	newInts := intSeq()

	print(newInts())
}
