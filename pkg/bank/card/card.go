package card

import "bank/pkg/bank/types"

const (
	Limit = 20_000
	DepositLimit = 50_000
	BonusLimit = 5_000
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card {
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}

	return card
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active {
		return card
	}

	if amount > card.Balance {
		return card
	}

	if amount < 0 {
		return card
	}

	if amount > (Limit * 100) {
		return card
	}

	card.Balance = card.Balance - amount
	return card
}

func Deposit(card *types.Card, amount types.Money)  {
	if !card.Active {
		return
	}

	if amount < 0 {
		return
	}

	if amount > DepositLimit * 100 {
		return
	}

	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int)  {
	if !card.Active {
		return
	}

	if card.Balance <= 0 {
		return
	}

	bonus := types.Money((int(card.MinBalance) * percent / 100) * daysInMonth / daysInYear)

	if bonus > BonusLimit * 100 {
		return
	}

	card.Balance += bonus
}

func Total(cards []types.Card) types.Money  {
	total := types.Money(0)
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			total += card.Balance
		}
	}

	return total
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSources []types.PaymentSource
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			paymentSource := types.PaymentSource{Type: "card", Number: string(card.PAN), Balance: card.Balance}
			paymentSources = append(paymentSources, paymentSource)
		}
	}

	return paymentSources
}
