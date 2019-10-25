package main

import "testing"

func TestBMI(t *testing.T) {
	var tests = []struct {
		height float64
		kg     float64
		want   float64
	}{
		{1.87, 76.93, 22},
		{1.78, 79.21, 25},
		{0, 32, 0},
		{-10.5, 0, 0},
	}
	for _, tt := range tests {
		if got := BMI(tt.height, tt.kg); got != tt.want {
			t.Errorf("round(%v, %v) = %v; want %v",
				tt.height,
				tt.kg,
				got,
				tt.want,
			)
		}
	}
}

func TestGoal(t *testing.T) {
	var tests = []struct {
		bmi  float64
		want float64
	}{
		{22, 21.9},
		{22.5, 22.4},
		{-5, 18.5},
	}
	for _, tt := range tests {
		if got := Goal(tt.bmi); got != tt.want {
			t.Errorf("Goals(%v) = %v; want %v",
				tt.bmi,
				got,
				tt.want,
			)
		}
	}
}

func TestKg(t *testing.T) {
	var tests = []struct {
		height float64
		bmi    float64
		want   float64
	}{
		{1.87, 22, 76.93},
		{1.78, 25, 79.21},
		{0, 32, 0},
		{-10.5, 0, 0},
		{1, -1.40, 0},
	}
	for _, tt := range tests {
		if got := Kg(tt.height, tt.bmi); got != tt.want {
			t.Errorf("round(%v, %v) = %v; want %v",
				tt.height,
				tt.bmi,
				got,
				tt.want,
			)
		}
	}
}

func TestRange(t *testing.T) {
	want := []float64{18.5, 19, 20, 21, 22, 23, 24, 24.99}
	got := Range()
	for i, g := range got {
		w := want[i]
		if g != w {
			t.Errorf("got %v; want %v", g, w)
		}
	}
}

func TestRound(t *testing.T) {
	var tests = []struct {
		places uint8
		in     float64
		want   float64
	}{
		{2, 1.267, 1.27},
		{1, 3.149, 3.1},
		{1, 0, 0},
		{1, -1.40, 0},
	}
	for _, tt := range tests {
		if got := round(tt.in, tt.places); got != tt.want {
			t.Errorf("round(%v, %v) = %v; want %v",
				tt.in,
				tt.places,
				got,
				tt.want,
			)
		}
	}
}
