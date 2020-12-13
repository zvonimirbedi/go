package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range slice {
		if slice[i]%2 == 0 {
			fmt.Println("Number", slice[i], "is even")
		} else {
			fmt.Println("Number", slice[i], "is odd")
		}
	}
}
