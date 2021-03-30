package transfer

import "bank/pkg/bank/types"

func Total(amount int) int {
	bonusPercent := 0.5

	total := (float64(amount) / 100.00) * bonusPercent
	return int(total) + amount
}


func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSources []types.PaymentSource
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			paymentSource := types.PaymentSource{Type: "card", Number: card.Number, Balance: card.Balance}
			paymentSources = append(paymentSources, paymentSource)
		}
	}

	return paymentSources
}
