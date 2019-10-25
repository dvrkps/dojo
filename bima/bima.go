package main

import "math"

// BMI calculate BMI from height and mass.
func BMI(height, kg float64) float64 {
	if height <= 0 || kg <= 0 {
		return 0
	}
	result := kg / math.Pow(height, 2)
	return round(result, 1)
}

// Goal returns first lower value in BMI.
func Goal(bmi float64) float64 {
	if bmi < 18.6 {
		bmi = 18.6
	}
	return round(bmi-0.1, 1)
}

// Kg calculate mass from height and BMI.
func Kg(height, bmi float64) float64 {
	weight := bmi * math.Pow(height, 2)
	return round(weight, 2)
}

// Range returns normal range BMIs.
func Range() []float64 {
	return []float64{18.5, 19, 20, 21, 22, 23, 24, 24.99}
}

// round returns rounded result to places decimals.
func round(value float64, places uint8) float64 {
	if value < 0 {
		return 0
	}
	pow := math.Pow(10, float64(places))
	digit := pow * value
	_, div := math.Modf(digit)
	result := math.Floor(digit)
	if div >= 0.5 {
		result = math.Ceil(digit)
	}
	return result / pow
}
