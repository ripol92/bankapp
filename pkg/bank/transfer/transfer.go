package transfer

func Total(amount int) int {
	bonusPercent := 0.5

	total := (float64(amount) / 100.00) * bonusPercent
	return int(total) + amount
}
