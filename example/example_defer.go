package example

import "fmt"

func Example_Defer() {
	defer e(f())

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
func e(_ int) {
	fmt.Println("world")
}

func f() int {
	fmt.Println("test")
	return 1
}
