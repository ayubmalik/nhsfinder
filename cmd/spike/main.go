package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday"}

	for _, day := range days {
		fmt.Printf("Day is %s\n", day)
	}
}
