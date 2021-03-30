package deposit

func Calculate(amount int, currency string) (minAmount int, maxAmount int)  {
	minPercent, maxPercent := PercentFor(currency)

	minAmount = ((amount / 100) * minPercent) / 12
	maxAmount = ((amount / 100) * maxPercent) / 12

	return
}

func PercentFor(currency string) (int, int)  {
	switch currency {
	case "TJS":
		return 4, 6
	case "USD":
		return 1, 2
	default:
		return 0, 0
	}
}