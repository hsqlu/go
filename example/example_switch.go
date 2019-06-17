package example

import (
	"fmt"
	"time"
)

func Example_Switch() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		case int64:
			fmt.Println("I'm an int64")
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(int64(1039))
}
