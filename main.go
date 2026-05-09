package main

import "fmt"

type Expense struct {
	Title  string
	Amount float64
}

func main() {
	// Slice to store expenses
	expenses := []Expense{}

	for {
		fmt.Println("\n==== Expense Tracker ====")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Total Expenses")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Choose option: ")
		fmt.Scanln(&choice)

		// Handle menu options
		if choice == 1 {

			var title string
			var amount float64

			fmt.Print("Enter title: ")
			fmt.Scanln(&title)

			fmt.Print("Enter amount: ")
			fmt.Scanln(&amount)

			
			expense := Expense{
				Title:  title,
				Amount: amount,
			}

			
			expenses = append(expenses, expense)

			fmt.Println("Expense added successfully!")

		} else if choice == 2 {

			if len(expenses) == 0 {
				fmt.Println("No expenses found")
				continue
			}

			fmt.Println("\n==== Expenses ====")

			for i, expense := range expenses {
				fmt.Printf("%d. %s - ₦%.2f\n",
					i+1,
					expense.Title,
					expense.Amount,
				)
			}

		} else if choice == 3 {

			var total float64

			for _, expense := range expenses {
				total += expense.Amount
			}

			fmt.Printf("Total Expenses: ₦%.2f\n", total)

		} else if choice == 4 {

			fmt.Println("Goodbye!")
			break

		} else {

			fmt.Println("Invalid choice")

		}
	}
}