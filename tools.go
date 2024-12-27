package main

type Tool struct {
	toolCode  string
	toolType  string
	toolBrand string
}

var Tools = map[string]Tool{
	"CHNS": {"CHNS", "Chainsaw", "Stihl"},
	"LADW": {"LADW", "Ladder", "Werner"},
	"JAKD": {"JAKD", "Jackhammer", "Dewalt"},
	"JAKR": {"JAKR", "Jackhammer", "Ridgid"},
}
