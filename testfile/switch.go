package main

import (
	"fmt"
)

func main() {
	a := 1
	switch {
	case a > 2:
		fmt.Println("66")
	default:
		fmt.Println(a)
	}
	switch a {
	case 1:
		fmt.Println("只因")
	}
}
