package main

import (
	"errors"
	"strconv"
	"time"
)

const DateLayout = "01/02/06"

func Checkout(toolCode, rentalDaysString, discountPercentString, checkoutDateString string) error {
	tool, ok := Tools[toolCode]
	if !ok {
		return errors.New("tool not found")
	}

	rentalDays, err := strconv.Atoi(rentalDaysString)
	if err != nil {
		return err
	}

	discountPercentage, err := strconv.Atoi(discountPercentString)
	if err != nil {
		return err
	}

	checkoutDate, err := time.ParseInLocation(DateLayout, checkoutDateString, time.Local)
	if err != nil {
		return err
	}

	if rentalDays <= 0 {
		return errors.New("the number of rental days must be 1 or greater")
	}

	if discountPercentage < 0 || discountPercentage > 100 {
		return errors.New("the discount percentage must be between 0 and 100")
	}

	agreement, err := CreateAgreement(tool, rentalDays, discountPercentage, checkoutDate)
	if err != nil {
		return err
	}

	receipt, err := CreateReceipt(agreement)
	if err != nil {
		return err
	}

	PrintReceipt(agreement, receipt)

	return nil
}
