package main

import "os"

func main() {
	b, err := os.ReadFile("testdata/booking-bjelovar.html")
	if err != nil {
		//		log.Printf("readfile: %v", err)
	}

	_ = b

}

type location struct {
	name string
	id   int64
	bID  int64
}

var bjelovar = location{
	name: "Bjelovar",
	id:   1,
	bID:  -75332,
}

// https://www.booking.com/searchresults.html?
// 	dest_type=city&
//	dest_id=-75332&
//	checkout_monthday=28&
//	checkin_month=7&
//	checkout_month=7&
//	checkout_year=2022&
//	checkin_monthday=27&
//	checkin_year=2022
