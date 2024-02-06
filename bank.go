package main

import (
	"fmt"
	"os"
	"strconv" 
	"errors"
)

//strconv helps with string conversion

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error){
	data, err := os.ReadFile(fileName) // using underscore simply means telling go u dont want to use that returned value right now

	if err != nil {
		return 1000.0, errors.New("failed to find file")
	}

	fileText := string(data)
	dataStr, err := strconv.ParseFloat(fileText, 64)

	if err != nil {
		return 1000.0, errors.New("failed to parse stored value")
	}

	return dataStr, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644) //0644 is a file permissions notation which means owner can read write and others can only read

}

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank")

	for {
		presentOptions()

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
			writeFloatToFile(accountBalance, accountBalanceFile)
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
			writeFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using our Bank.")
			return
		}
	}
}