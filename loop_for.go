package main

import (
	"fmt"
	//"reflect"
)

func main() {

	for i := 0; i <= 10; i++ {
		if i == 10 {
			fmt.Print(i)
		} else {
			fmt.Print(i, " - ")
		}
	}

	fmt.Println()

	for i := 10; i >= 0; i-- {
		if i == 0 {
			fmt.Print(i)
		} else {
			fmt.Print(i, " - ")
		}
	}

}
