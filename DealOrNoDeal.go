package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const numCases = 20

type Game struct {
	cases            [numCases]int
	playerCaseIndex  int
	playerCaseValue  int
	round            int
	bankOffer        int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{}
	game.initialize()

	fmt.Println("Welcome to Deal or No Deal")
	fmt.Println()

	game.printCases()
	game.selectInitialCase()

	for {
		game.playRound()
		if game.acceptOffer() {
			break
		}
		game.round++
	}
}

func (g *Game) initialize() {
	values := [numCases]int{10, 15, 20, 25, 30, 40, 45, 55, 75, 100, 125, 150, 200, 250, 275, 325, 400, 500, 750, 1000}

	// Fisher-Yates shuffle
	for i := len(values) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		values[i], values[j] = values[j], values[i]
	}

	g.cases = values
	g.round = 1
}

func (g *Game) printCases() {
	fmt.Println()
	for i := 0; i < numCases; i++ {
		switch {
		case g.cases[i] == -1:
			fmt.Printf("Case %2d ❌ ", i)
		case g.cases[i] == -2:
			fmt.Printf("Case %2d ✅ ", i)
		default:
			fmt.Printf("Case %2d 💼 ", i)
		}

		if (i+1)%5 == 0 || i == numCases-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func (g *Game) selectInitialCase() {
	fmt.Print("What case number did you select as your initial case? ")
	fmt.Scan(&g.playerCaseIndex)

	if g.playerCaseIndex < 0 || g.playerCaseIndex >= numCases {
		fmt.Println("Invalid case number. Using case 0.")
		g.playerCaseIndex = 0
	}

	g.playerCaseValue = g.cases[g.playerCaseIndex]
	g.cases[g.playerCaseIndex] = -2

	fmt.Printf("Great! We'll reveal Case %d at the end.\n", g.playerCaseIndex)
}

func (g *Game) playRound() {
	casesToOpen := 5
	fmt.Printf("\nRound %d: Open %d cases\n", g.round, casesToOpen)

	for i := 0; i < casesToOpen; i++ {
		g.printCases()
		var selection int
		fmt.Printf("Select case %d: ", i+1)
		fmt.Scan(&selection)

		if selection < 0 || selection >= numCases || g.cases[selection] < 0 {
			fmt.Println("Invalid selection. Try again.")
			i--
			continue
		}

		fmt.Printf("Case %d contained $%d\n", selection, g.cases[selection])
		g.cases[selection] = -1
	}
}

func (g *Game) calculateBankOffer() int {
	max := 0
	for _, val := range g.cases {
		if val > max {
			max = val
		}
	}
	// Improved offer calculation
	return max + rand.Intn(max/3+50)
}

func (g *Game) acceptOffer() bool {
	g.bankOffer = g.calculateBankOffer()

	fmt.Printf("\nThe bank is offering: $%d\n", g.bankOffer)
	fmt.Print("DEAL or NO DEAL? (deal/no): ")

	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(strings.TrimSpace(response))

	if response == "deal" {
		fmt.Printf("\nCongratulations! You accepted the bank's offer of $%d\n", g.bankOffer)
		fmt.Printf("Your original case contained: $%d\n", g.playerCaseValue)
		return true
	}

	return false
}
