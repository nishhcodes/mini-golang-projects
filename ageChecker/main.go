package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInfo() (string, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	var age int
	fmt.Print("What's your Age: ")
	fmt.Scanf("%d", &age)

	return name, age
}	

func isAdult(age int, name string) {
	if age >= 18 {
		fmt.Println(name, "You are an Adult!")
	} else {
		fmt.Println(name, "You are not an Adult!")
	}
}

func main() {
	name, age := getInfo()
	isAdult(age, name)
}