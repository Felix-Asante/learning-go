package main

import (
	"fmt"

	"example.com/working-with-packages/utils"
)

func main() {

	accountBalance, error := utils.GetFloatFromFile(utils.BalanceFileName)
	endApp := false

	if error != nil {
		fmt.Println("ERROR !!", error)
		fmt.Println("================")
	}

	fmt.Println("Welcome to Go bank !")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			accountBalance = utils.DepositToAccount(accountBalance)
		case 3:
			accountBalance = utils.WithrawFromAmount(accountBalance)
		default:
			fmt.Println("Goodbye !")
			endApp = true
		}

		if endApp {
			break
		}
	}

	fmt.Println("Thanks for choosing our bank !")

}
