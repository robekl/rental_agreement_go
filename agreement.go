package main

import (
	"errors"
	"time"
)

type Agreement struct {
	Tool            Tool
	RentalDays      int
	DiscountPercent int
	CheckoutDate    time.Time
	Charge          Charge
}

func CreateAgreement(tool Tool, days int, discountPercent int, checkoutDate time.Time) (Agreement, error) {
	charge, ok := Charges[tool.toolType]
	if !ok {
		return Agreement{}, errors.New("invalid tool type")
	}

	return Agreement{tool, days, discountPercent, checkoutDate, charge}, nil
}
