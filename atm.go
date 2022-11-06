package main

import (
	"fmt"
	"os"
)

//welcome page function
//get inputs from the console
// Change Pin (you can create a default pin in your code and then change it in the function)
// Account balance
// Withdraw funds
// Deposit funds
// Cancel/Exit (selecting this option should exit the program)

type UserData struct {
	userName string
	pin      int
	balance  int
}

var user1 = UserData{
	userName: "sam",
	pin:      0000,
	balance:  60000,
}

func main() {

	for {
		welcomePage()

		fmt.Println("")
		fmt.Println("Please enter your four digit pin:")

		_, err := fmt.Scan(&user1.pin)
		if err != nil {
			fmt.Println("Invalid Pin")
		}

		if user1.pin == 0000 {
			fmt.Println("")
			fmt.Println("Change the default pin from '000' to a more secure four digit pin")
			menuPage()
		} else {
			fmt.Println("Incorrect pin. Enter correct pin to continue.")
			menuPage()
		}
	}

}

func welcomePage() {
	fmt.Println("Welcome to our your mobile ATM App")
}

func menuPage() {
	var input int

	fmt.Println("---------------------------------")
	fmt.Println("Press 1 to change your pin")
	fmt.Println("Press 2 to check your account balance")
	fmt.Println("Press 3 to withdraw from your account")
	fmt.Println("Press 4 to deposit into your account")
	fmt.Println("Press 0 to exit this app")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid Input")
	}

	switch input {
	case 1:
		changePin()
	case 2:
		accountBalance()
	case 3:
		withdrawMoney()
	case 4:
		depositMoney()
	case 0:
		exitApp()
	default:
		fmt.Println("Enter a valid menu number")
		menuPage()
	}
}

func changePin() {
	fmt.Println("")
	fmt.Println("Please enter your new pin: ")
	var updatedPin int
	_, err := fmt.Scan(&updatedPin)
	if err != nil {
		fmt.Println("Error: Invalid input")
	}
	user1.pin = updatedPin
	fmt.Println("Pin changed successfully.")
	doAnotherTransaction()
}

// check account balance
func accountBalance() {
	fmt.Println("")
	fmt.Println("Your account balance is: ", user1.balance)
	doAnotherTransaction()
}

// withdraw money
func withdrawMoney() {
	fmt.Println("")
	fmt.Println("Kindly enter the amount you want to withdraw: ")
	var amount int
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Error: Invalid input")
	}
	if amount > user1.balance {
		fmt.Println("Error: Insufficient funds")
	} else {
		user1.balance -= amount
		fmt.Println("Withdrawal performed successfully")
	}
	doAnotherTransaction()
}

// DepositMoney deposit money
func depositMoney() {
	fmt.Println("")
	fmt.Println("Kindly enter the amount you want to deposit: ")
	var amount int
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Error: Invalid input")
	}
	user1.balance += amount
	fmt.Println("Deposit performed successfully")
	doAnotherTransaction()
}

func doAnotherTransaction() {
	fmt.Println("Do you want to perform another transaction? (Y/N)")
	var response string
	_, err := fmt.Scan(&response)
	if err != nil {
		fmt.Println("Error: Invalid input")
	}
	if response == "Y" {
		menuPage()
	} else if response == "N" {
		exitApp()
	} else {
		fmt.Println("Error: Invalid input")
	}
}

// exit program
func exitApp() {
	fmt.Println("Thanks your for using our service.")
	os.Exit(0)
}
