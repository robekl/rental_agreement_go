package main

import (
	"fmt"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"math"
	"time"
)

type Receipt struct {
	ChargeDays             int
	PreDiscountChargeCents int
	DiscountPercent        int
	DiscountCents          int
	FinalChargeCents       int
}

func CreateReceipt(agreement Agreement) (Receipt, error) {
	chargeDays, err := calculateChargeDays(agreement)
	if err != nil {
		return Receipt{}, err
	}
	chargeCents := calculateChargeCents(agreement, chargeDays)
	discountCents := calculateDiscountCents(agreement, chargeCents)

	receipt := Receipt{
		chargeDays,
		chargeCents,
		agreement.DiscountPercent,
		discountCents,
		chargeCents - discountCents,
	}

	return receipt, nil
}

func calculateDiscountCents(agreement Agreement, preDiscountChargeCents int) int {
	return int(math.Round(float64(preDiscountChargeCents*agreement.DiscountPercent) / 100))
}

func calculateChargeCents(agreement Agreement, days int) int {
	return agreement.Charge.dailyChargeCents * days
}

// calculateChargeDays calculates the number of days, starting
// on the day after the checkout date, that should be charged for
// renting. Excludes days that are a holiday, when the item
func calculateChargeDays(agreement Agreement) (int, error) {
	charge := agreement.Charge

	holidays, err := GetHolidays(agreement.CheckoutDate, agreement.RentalDays)
	if err != nil {
		return 0, err
	}
	date := agreement.CheckoutDate.AddDate(0, 0, 1)
	chargeDays := 0

	rentalDays := agreement.RentalDays

	for rentalDays > 0 {
		_, dateIsAHoliday := holidays[date]
		// Count chargeable days, excluding “no charge” days
		if (isWeekend(date.Weekday()) && !charge.IsChargedOnWeekend) ||
			(isWeekday(date.Weekday()) && !charge.IsChargedOnWeekday) ||
			(dateIsAHoliday && !charge.IsChargedOnHoliday) {
			// this type of day is excluded
		} else {
			chargeDays++
		}
		date = date.AddDate(0, 0, 1)

		rentalDays--
	}

	return chargeDays, nil
}

func isWeekend(day time.Weekday) bool {
	return day == time.Saturday || day == time.Sunday
}

func isWeekday(day time.Weekday) bool {
	return !isWeekend(day)
}

func PrintReceipt(agreement Agreement, receipt Receipt) {
	fmt.Println("Tool code: " + agreement.Tool.toolCode)
	fmt.Println("Tool type: " + agreement.Tool.toolType)
	fmt.Println("Tool brand: " + agreement.Tool.toolBrand)
	fmt.Printf("Rental days: %d\n", agreement.RentalDays)
	fmt.Println("Check out date: " + agreement.CheckoutDate.Format(DateLayout))
	fmt.Println("Due date: " + agreement.CheckoutDate.AddDate(0, 0, agreement.RentalDays).Format(DateLayout))
	fmt.Println("Daily rental charge: " + currencyString(agreement.Charge.dailyChargeCents))
	fmt.Printf("Charge days: %d\n", receipt.ChargeDays)
	fmt.Println("Pre-discount charge: " + currencyString(receipt.PreDiscountChargeCents))
	fmt.Printf("Discount percent: %d%%\n", agreement.DiscountPercent)
	fmt.Println("Discount amount: " + currencyString(receipt.DiscountCents))
	fmt.Println("Final charge: " + currencyString(receipt.FinalChargeCents))
}

func currencyString(cents int) string {
	dec := number.Decimal(float64(cents)/100, number.Scale(2))
	p := message.NewPrinter(language.English)
	return p.Sprintf("%v%v", currency.Symbol(currency.USD), dec)
}
