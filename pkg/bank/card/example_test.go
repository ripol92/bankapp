package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output: 0
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)

	// Output: 2000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 990_000_00)
	fmt.Println(result.Balance)

	// Output: 2000000
}

func ExampleDeposit_success() {
	card := types.Card{Balance: 10_000, Active: true}
	Deposit(&card, 10_000)
	fmt.Println(card.Balance)

	// Output: 20000
}

func ExampleDeposit_notActive() {
	card := types.Card{Balance: 10_000, Active: false}
	Deposit(&card, 10_000)
	fmt.Println(card.Balance)

	// Output: 10000
}

func ExampleDeposit_aboveLimit() {
	card := types.Card{Balance: 10_000, Active: true}
	Deposit(&card, 100_000_000)
	fmt.Println(card.Balance)

	// Output: 10000
}

func ExampleAddBonus_success() {
	card := types.Card{Balance: 10_000, Active: true, MinBalance: 10_000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: 10024
}

func ExampleAddBonus_notActive() {
	card := types.Card{Balance: 10_000, Active: false, MinBalance: 10_000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: 10000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -10_000, Active: true, MinBalance: 10_000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: -10000
}

func ExampleAddBonus_aboveLimit() {
	card := types.Card{Balance: 10_000_00_00, Active: true, MinBalance: 10_000_00_00}
	AddBonus(&card, 10, 30, 365)
	fmt.Println(card.Balance)

	// Output: 100000000
}

func ExampleTotal() {
	cards := []types.Card{
		{Balance: 10_000_00, Active: true},
		{Balance: 10_000_00, Active: true},
		{Balance: 10_000_00, Active: true},
	}
	total := Total(cards)
	fmt.Println(total)

	//Output: 3000000
}

func ExampleTotal_oneNotActiveCard() {
	cards := []types.Card{
		{Balance: 10_000_00, Active: true},
		{Balance: 10_000_00, Active: true},
		{Balance: 10_000_00, Active: false},
	}
	total := Total(cards)
	fmt.Println(total)

	//Output: 2000000
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{Balance: 0, Active: true, PAN: "5058 xxxx xxxx 8888"},
		{Balance: 10_000_00, Active: true, PAN: "5058 xxxx xxxx 8889"},
		{Balance: 10_000_00, Active: true, PAN: "5058 xxxx xxxx 88899"},
		{Balance: 10_000_00, Active: false, PAN: "5058 xxxx xxxx 8810"},
	}

	paymentSources := PaymentSources(cards)
	for _, source := range paymentSources {
		fmt.Println(source.Number)
	}

	// Output: 5058 xxxx xxxx 8889
	// 5058 xxxx xxxx 88899
}

