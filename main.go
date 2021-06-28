package main

import (
	"fmt"
)	

type Currency struct {
	Twothousands int
	Fivehundreds int
	Hundreds int
}

var twos, fives, ones = 1000, 100, 100

func main() {
	availableCurrencies := &Currency{
		Twothousands : 10,
		Fivehundreds : 20,
		Hundreds : 30,
	}
	Start:  
	fmt.Println("Welcome Please select your action. \n 1. Press 1 for Withdrawl  \n 2. Press 2 for Add Currencies \n")
	var option string
	_, err := fmt.Scanln(&option)
	if err != nil {
		fmt.Println("Sorry Error Occured while processing!")
		return
	}
	switch option {
		case "2":
			Feeder(availableCurrencies)
			goto Start
		case "1":
			Denominator(availableCurrencies)
			goto Start
		default:
			break		
	}
}

func Denominator(currency *Currency) {
	fmt.Println("Please type your amount to debit: ")
	var amount int
	fmt.Scanln(&amount)
	if amount <= 10000 && validateAmount(amount) {
		if amount > 2000 {
			twoDenoms,balance := splitter(amount,2000,currency)
			if twoDenoms <= currency.Twothousands {
				currency.Twothousands = currency.Twothousands - twoDenoms
				fmt.Println("Number of 2000s:", twoDenoms)
			} else {
				balance = amount
			}
			if balance >= 500 {
				handleFiveHundreds(balance, currency)
			} else {
				handleHundreds(balance, currency)
			}
		} else if amount >= 500 {
			handleFiveHundreds(amount, currency)
			fmt.Println(currency)
		}
	} else {
		fmt.Println("Invalid Amount entered, Please enter multiples of 100s and amount should not exceeds 10000")
	}
}

func splitter(amount,value int,currency *Currency) (denom,x int) {
	denom = amount / value
	x = amount - (denom * value)
	return
}

func handleFiveHundreds(balance int, currency *Currency) {
	fiveDenoms,fiveBalance := splitter(balance,500,currency)
	if fiveDenoms <= currency.Fivehundreds {
		currency.Fivehundreds = currency.Fivehundreds - fiveDenoms
		fmt.Println("Number of 500s:", fiveDenoms)
	} else {
		fiveBalance = balance	
	}
	handleHundreds(fiveBalance, currency)
}

func handleHundreds(balance int, currency *Currency) {
	if balance >= 100 {
		oneDenoms, _ := splitter(balance,100,currency)
		if oneDenoms <= currency.Hundreds {
			currency.Hundreds = currency.Hundreds - oneDenoms
			fmt.Println("Number of 100s:",oneDenoms)
		} else {
			fmt.Println("Sorry! Running out of cash")
		}
	}
}

func validateAmount(amount int) bool {
	return amount % 100 == 0
}

func Feeder(currency *Currency) {
	fmt.Println("Enter 2000 currencies quantity: ");
	var twoThousands int
	fmt.Scanln(&twoThousands)
	currency.Twothousands += twoThousands
	
	fmt.Println("Enter 500 Currencies quantity: ");
	var fiveHundreds int
	fmt.Scanln(&fiveHundreds)
	currency.Fivehundreds += fiveHundreds
	
	fmt.Println("Enter 100 Currencies quantity: ");
	var hundreds int
	fmt.Scanln(&hundreds)
	currency.Hundreds += hundreds
}
