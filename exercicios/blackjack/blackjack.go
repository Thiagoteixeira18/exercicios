package exercicios

// Function to calculate the numerical value of a card
// Function to calculate the numerical value of a card
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// Function to determine the action to take in the first turn of Blackjack
func FirstTurn(card1, card2, dealerCard string) string {
	totalValue := ParseCard(card1) + ParseCard(card2)

	if (card1 == "ace" || card1 == "ás") && (card2 == "ace" || card2 == "ás") {
		return "P" // Split
	} else if totalValue == 21 {
		if dealerCard == "ace" || dealerCard == "ten" || dealerCard == "jack" || dealerCard == "queen" || dealerCard == "king" {
			return "S" // Stand
		} else {
			return "W" // Automatically win
		}
	} else if totalValue >= 17 && totalValue <= 20 {
		return "S" // Stand
	} else if (totalValue >= 12 && totalValue <= 16) {
		if ParseCard(dealerCard) >= 7 {
			return "H" // Hit
		} else {
			return "S" // Stand
		}
	} else {
		return "H" // Hit as default action
	}
}