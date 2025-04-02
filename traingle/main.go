package main

import "fmt"

func main() {
	var rows int

	fmt.Print("Enter the number of Rows: ")
	fmt.Scanf("%d", &rows)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}