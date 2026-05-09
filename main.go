package main

import "fmt"

type Expense struct {

	Title string
	Amount float64
}

func main() {
	fmt.Println("Expense tracker")

	expenses := []Expense{}

	for {

		fmt.Println("==== Expense Tracker ====")
		fmt.Println("Add Expense")
		fmt.Println("View Expense")
		fmt.Println("Total Expense")
		fmt.Println("Exit")

		var choice int
		fmt.Print("Choose option:")
		fmt.Scanln(&choice)

		fmt.Println(choice)

		if choice == 1 {
			fmt.Println("Add Expense")
		} else if choice == 2 {
			fmt.Println("View Expenses")
		} else if choice == 3 {
			fmt.Println("Total Expense")
		} else if choice == 4 {
			fmt.Println("Bye,Bye")
		} else {
			fmt.Println("invalid")
		}

	}
}