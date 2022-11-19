package main

import (
	"fmt"
)

func main() {

	for eonum := 0; eonum <= 10; eonum++ {
		if eonum%2 == 0 {
			fmt.Println(eonum, "is even")
		} else {
			fmt.Println(eonum, "is odd")
		}
	}
}
