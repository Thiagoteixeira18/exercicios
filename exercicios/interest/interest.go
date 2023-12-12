package exercicios

func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

func Interest(balance float64) float64 {
	interestRate := float64(InterestRate(balance)) 
	return balance * interestRate / 100
}

func AnnualBalanceUpdate(balance float64) float64 {
	interest := Interest(balance)
	return balance + interest
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0 
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
		if years > 1000 {
			break
		}
	}
	return years
}