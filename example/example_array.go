package example

import (
	"fmt"
)

func Example_Arrays() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println("Set: ", a)
	fmt.Println("Get: ", a[4])
	fmt.Println("Length: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoDimensinal [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensinal[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoDimensinal)
}
