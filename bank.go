package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do?")
	fmt.Print("1. Check Balance\n2. Deposit Money\n3. Withdraw Money\n4. Exit\n")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance )
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Your withdrawal: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		accountBalance -= withdrawalAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}