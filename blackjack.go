package exercicios

// Function to calculate the numerical value of a card
func ParseCard(card string) int {
	switch card {
	case "ace", "ás":
		return 11
	case "two", "dois":
		return 2
	case "three", "três":
		return 3
	case "four", "quatro":
		return 4
	case "five", "cinco":
		return 5
	case "six", "seis":
		return 6
	case "seven", "sete":
		return 7
	case "eight", "oito":
		return 8
	case "nine", "nove":
		return 9
	case "ten", "dez":
		return 10
	case "jack", "queen", "king", "valete", "dama", "rei":
		return 10
	default:
		return 0
	}
}

// Function to determine the action to take in the first turn of Blackjack
func FirstTurn(card1, card2, dealerCard string) string {
	totalValue := ParseCard(card1) + ParseCard(card2)

	if card1 == "ace" || card1 == "ás" {
		if card2 == "ace" || card2 == "ás" {
			return "P"
		}
	} else if totalValue == 21 && (dealerCard != "ace" && dealerCard != "ten" && dealerCard != "ás" && dealerCard != "dez") {
		return "W"
	} else if (totalValue >= 17 && totalValue <= 20) || (totalValue == 21 && (card1 == "ace" || card1 == "ás" || card2 == "ace" || card2 == "ás")) {
		return "S"
	} else if (totalValue >= 12 && totalValue <= 16) || (totalValue == 21 && (dealerCard == "seven" || dealerCard == "sete" || dealerCard == "eight" || dealerCard == "oito" || dealerCard == "nine" || dealerCard == "nove" || dealerCard == "ten" || dealerCard == "dez" || dealerCard == "ace" || dealerCard == "ás") && (card1 == "ace" || card1 == "ás" || card2 == "ace" || card2 == "ás")) {
		return "H"
	} else {
		return "H" // Add a default return value here
	}
	return "ola"
}
