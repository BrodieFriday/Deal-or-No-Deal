package main

import (
	"fmt"
	"math/rand"
)

var initialCaseNumber int
var initialCaseValue int
var offerFromBank int
var responseToBankOffer string
var selectionRound int
var arrCases = [20]int{
	10, 15, 20, 25, 30, 40, 45, 55, 75, 100, 125, 150, 200, 250, 275, 325, 400, 500, 750, 1000,
}

func main() {
	fmt.Println("Welcome to Deal or No Deal")
	fmt.Println()
	fmt.Println("Please select a case to get started")
	fmt.Println()

	// fmt.Println(arrCases)
	// Randomize the cases
	for i := range arrCases {
		j := rand.Intn(i + 1)
		arrCases[i], arrCases[j] = arrCases[j], arrCases[i]
	}

	printCases()

	selectInitialCase()

	caseSelection()

	for responseToBankOffer != "deal" {
		caseSelection()
	}

}

func printCases() {
	fmt.Println()

	// Print the cases
	for i := 0; i < len(arrCases); i++ {
		// Display the emoji.
		if arrCases[i] == -1 {
			fmt.Print("case ", i, " ", "❌", " ")
		} else if arrCases[i] == -2 {
			fmt.Print("case ", i, " ", "✅", " ")
		} else {
			fmt.Print("case ", i, " ", "💼", " ")
		}
		if i == 5 || i == 10 || i == 15 || i == len(arrCases) {
			fmt.Println()
		}
	}
	fmt.Println()
}

func selectInitialCase() {
	fmt.Println()
	fmt.Print("What case number did you select as your initial case? ")
	fmt.Scan(&initialCaseNumber)
	fmt.Println("Great! We will find out what is in case ", initialCaseNumber, " later.")
	initialCaseValue = arrCases[initialCaseNumber]
	arrCases[initialCaseNumber] = -2
	selectionRound = 1

}

func caseSelection() {
	fmt.Println()
	fmt.Println()
	if selectionRound == 1 {
		fmt.Println("Lets start by opening up five cases")
	} else {
		fmt.Println("Lets start by opening up five more cases")
	}
	var lastOpenedCase int
	for i := 1; i < 6; i++ {
		printCases()
		fmt.Print("What is your selection for case ", i, " ?")
		fmt.Scan(&lastOpenedCase)
		fmt.Println("Case ", lastOpenedCase, " contained ", arrCases[lastOpenedCase])
		arrCases[lastOpenedCase] = -1
		fmt.Println()
	}

	offerFromBanker()
}

func offerFromBanker() {
	max := arrCases[0]
	for i := 1; i < len(arrCases); i++ {
		if max < arrCases[i] {
			max = arrCases[i]
		}
	}

	offerFromBank = max + rand.Intn(100)
	fmt.Println("The bank is offering you ", offerFromBank)
	fmt.Println("DEAL OR NO DEAL ?")
	fmt.Scanln(&responseToBankOffer)

	if responseToBankOffer == "deal" || responseToBankOffer == "DEAL" {
		fmt.Println("CONGRADULATIONS! You won $", offerFromBank)
		fmt.Println("Your case value was $", initialCaseValue)
	} else {
		selectionRound = 2
	}

}
