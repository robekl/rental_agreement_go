package main

type Charge struct {
	Type               string
	dailyChargeCents   int
	IsChargedOnWeekday bool
	IsChargedOnWeekend bool
	IsChargedOnHoliday bool
}

var Charges = map[string]Charge{
	"Ladder":     {"Ladder", 199, true, true, false},
	"Chainsaw":   {"Chainsaw", 149, true, false, true},
	"Jackhammer": {"Jackhammer", 299, true, false, false},
}
