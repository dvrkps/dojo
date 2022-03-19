package main

import "fmt"

type itemFloat struct {
	Count    float64
	Eur      float64
	Hrk      float64
	Rate     float64
	TotalEur float64
	TotalHrk float64
}

func (i *itemFloat) SetEur(count, eur float64) {
	i.Rate = exchangeRate

	i.Count = count
	i.Eur = eur
	i.TotalEur = i.Count * i.Eur

	i.Hrk = calcHrkFloat(i.Eur)
	i.TotalHrk = i.Count * i.Hrk
}

func (i *itemFloat) SetHrk(count, hrk float64) {
	i.Rate = exchangeRate

	i.Count = count
	i.Hrk = hrk
	i.TotalHrk = i.Count * i.Hrk

	i.Eur = calcEurFloat(i.Hrk)
	i.TotalEur = i.Count * i.Eur
}

func (i *itemFloat) String() string {
	s := fmt.Sprintf("%6.2f %6.2f %6.2f %6.2f %8.2f %8.2f",
		i.Count,
		i.Eur,
		i.Hrk,
		i.Rate,
		i.TotalEur,
		i.TotalHrk,
	)

	return s
}

func calcHrkFloat(eur float64) float64 {
	return eur * exchangeRate
}

func calcEurFloat(hrk float64) float64 {
	return hrk / exchangeRate
}

func sumFloat(items []itemFloat) (float64, float64) {
	var sumEur, sumHrk float64
	for _, i := range items {
		sumEur = sumEur + i.Eur
		sumHrk = sumHrk + i.Hrk
	}

	return sumEur, sumHrk
}
