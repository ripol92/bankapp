package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{ID: 200, Amount: types.Money(400)},
		{ID: 101, Amount: types.Money(200)},
		{ID: 102, Amount: types.Money(300)},
	}

	max := Max(payments)
	fmt.Println(max.ID)

	// Output: 200
}
