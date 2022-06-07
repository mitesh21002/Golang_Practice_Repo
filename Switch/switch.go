package main

import "fmt"

func main() {
	var ans int
	fmt.Println("Enter the choice")
	fmt.Scanln(&ans)

	// ans := 3
	// during switch give only one type of datatype in input case "m1", "m2"  or case 1,2

	switch ans {
	case 1:
		fmt.Println("One")
		fmt.Println("Oonnee")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println(" there is no match")
	}
}
