package main

import "fmt"
import "time"

func print(str string) {
	fmt.Println(str)
}

func main() {
	i := 3
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		print("one")
	case 2:
		print("two")
	case 3:
		print("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		print("It's the weekend")
	default:
		print("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		print("It's before noon")
	default:
		print("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			print("I'm a bool")
		case int:
			print("I'm an int")
		default:
			fmt.Printf("Don't knowk type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
