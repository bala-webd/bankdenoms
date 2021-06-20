package main

import (
	"fmt"
)	

var twos, fives, ones = 1000, 100, 100

func main() {
	fmt.Println("Welcome Please select your action. 1. Type WD for Withdrawl 2. Type Feed for add denoms")
	var option string
	_, err := fmt.Scanln(&option)
	if err != nil {
		fmt.Println("Sorry Error Occured while processing!")
		return
	}
	switch option {
		case "Feed":
			feeder()
		case "WD":
			denominator()
		default:
			break		
	}
}

func denominator() {
	fmt.Println("Please type your amount to debit: ")
	var amount int
	fmt.Scanln(&amount)
	if amount <= 10000 && validateAmount(amount) {
		if amount > 2000 {
			twoDenoms,balance := splitter(amount,2000)
			fmt.Println("Number of 2000s:", twoDenoms)
			if balance >= 500 {
				fiveDenoms,fiveBalance := splitter(balance,500)
				fmt.Println("Number of 500s:", fiveDenoms)
				if fiveBalance >= 100 && fiveBalance < 500 {
					oneDenoms, _ := splitter(fiveBalance,100)
					fmt.Println("Number of 100s:",oneDenoms)
				}
			}
		} else if amount >= 500 {
			fiveDenoms, fiveBalance := splitter(amount,500)
			fmt.Println("Number of 500s:",fiveDenoms)
			if fiveBalance >= 100 {
				oneDenoms, _ := splitter(fiveDenoms,100)
				fmt.Println("Number of 100s:",oneDenoms)
			}
		}
	} else {
		fmt.Println("Invalid Amount entered, Please enter multiples of 100s and amount should not exceeds 10000")
	}
}

func splitter(amount,value int) (denom,x int) {
	denom = amount / value
	x = amount - (denom * value)
	return
}

func validateAmount(amount int) bool {
	return amount % 100 == 0
}

func feeder() {

}
