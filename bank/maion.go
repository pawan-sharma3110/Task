package main

import "fmt"

var accountBalance float64 = 1000.0
var choice int

func main() {

	userGreet()
	choice = userChoise()
	userResponse(choice)

}
func userGreet() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you eant to do?")
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
	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit amount :")
		var depositAmmount float64
		fmt.Scan(&depositAmmount)
		if depositAmmount <= 0 {
			fmt.Println("Invalid amount. Amount must be grater then 0")
			return
		}
		accountBalance += depositAmmount //accountBalance + depositAmmount
		fmt.Println("Total Balance: ", accountBalance)
	} else if choice == 3 {
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
		fmt.Println("Total Remaining Balance: ", accountBalance)
	} else if choice == 4 {
		fmt.Println("Good By!")
		return
	} else {
		fmt.Println(" Please enter valid input.")
		return
	}
}
