package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Pallinder/go-randomdata"
)

var accountBalance, _ = getBalance()
var choice int

func main() {

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Contect us 24*7 at :", randomdata.PhoneNumber())
	for {
		userGreet()
		choice = userChoise()

		if choice == 4 {
			break
		}
		userResponse(choice)
	}

}
func userGreet() {

	fmt.Println("What do you want to do?")
	fmt.Println("1. Ckeck money")
	fmt.Println("2. Depodit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
func userChoise() int {
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	fmt.Println("Your choice: ", choice)
	return choice
}

func userResponse(choice int) {
	switch choice {
	case 1:
		fmt.Println("Your balance is:", accountBalance)
	case 2:
		fmt.Print("Your deposit amount :")
		var depositAmmount float64
		fmt.Scan(&depositAmmount)
		if depositAmmount <= 0 {
			fmt.Println("Invalid amount. Amount must be grater then 0")
			return
		}
		accountBalance += depositAmmount //accountBalance + depositAmmount
		writeBalance(accountBalance)
		fmt.Println("Total Balance: ", accountBalance)
	case 3:

		fmt.Print("How much amount You withdraw ? :")
		var withdrawBalance float64
		fmt.Scan(&withdrawBalance)
		if withdrawBalance <= 0 {
			fmt.Println("Invalid amount. Amount must be grater then 0")
			return
		}
		if withdrawBalance > accountBalance {
			fmt.Println("You can't withdraw more then you have.")
			return
		}
		accountBalance -= withdrawBalance //accountBalance - depositAmmount
		writeBalance(accountBalance)
		fmt.Println("Total Remaining Balance: ", accountBalance)
	case 4:
		fmt.Println("Good By!")
		return
	default:
		fmt.Println(" Please enter valid input.")
		return
	}

}
func writeBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
func getBalance() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("error while getin balance")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("error while converting balance in float64")
	}
	return balance, nil
}
