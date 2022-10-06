package main

import (
	"fmt"
	"os"
	"strconv"
)

var Data []Friend

type Friend struct {
	Name, Address, Job, Reason string
}

func init() {
	Data = []Friend{
		{Name: "Klarrisa Scutts", Address: "2472 Kennedy Court", Job: "Food Chemist", Reason: "Neutrogena Wet Skin Kids Beach and Pool"},
		{Name: "Gerrilee Elflain", Address: "0 Beilfuss Court", Job: "Electrical Engineer", Reason: "VIIBRYD"},
		{Name: "Karin Brasher", Address: "8 Vahlen Junction", Job: "Analog Circuit Design manager", Reason: "NITROFURANTOIN"},
		{Name: "Quintana Rackham", Address: "0998 Old Gate Avenue", Job: "Sales Associate", Reason: "Xylocaine"},
		{Name: "Suzette Spawton", Address: "5137 Summit Alley", Job: "Internal Auditor", Reason: "Equaline Anti Diarrheal"},
		{Name: "Leonardo Lenham", Address: "39 Springview Pass", Job: "Assistant Professor", Reason: "Eye Drops AC"},
		{Name: "Alfonso Aldhous", Address: "625 Prairieview Point", Job: "Web Developer IV", Reason: "Premier Value Famotidine"},
		{Name: "Lisle Backman", Address: "3 Meadow Valley Drive", Job: "VP Sales", Reason: "Losartan Potassium and Hydrochlorothiazide"},
		{Name: "Earle MacShirie", Address: "050 Lunder Place", Job: "Design Engineer", Reason: "SALIX NIGRA POLLEN"},
		{Name: "Jeannine Weedall", Address: "26 Fairview Road", Job: "Assistant Professor", Reason: "Levetiracetam"},
	}
}

func main() {
	input := os.Args[1:]
	inputNumber, err := strconv.Atoi(input[0])
	if err != nil {
		panic("cannot convert string to int")
	}
	fmt.Println(Data[inputNumber-1])
}
