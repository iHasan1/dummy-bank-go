package main

import (
	"fmt"
	"os"
	"strconv" 
	"errors"
)

//strconv helps with string conversion

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFile) // using underscore simply means telling go u dont want to use that returned value right now

	if err != nil {
		return 1000.0, errors.New("Failed to find file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000.0, errors.New("Failed to parse stored value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644) //0644 is a file permissions notation which means owner can read write and others can only read

}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
	}

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Print("1. Check Balance\n2. Deposit Money\n3. Withdraw Money\n4. Exit\n")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance )
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount, Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile((accountBalance))
		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount < 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			if withdrawalAmount >= accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				// return
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using our Bank.")
			return
		}
	}
}