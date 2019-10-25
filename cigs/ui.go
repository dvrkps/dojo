package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Usage holds command usage description.
const Usage = `Usage
	cigs [options]... [values]...

Options:
	-d	goal in days
	-s	since last cig(yyyy-mm-dd)
`

// Version is application version.
const Version = "0.3.6"

// daysResult returns titled days result.
func daysResult(from, to time.Time, days int) string {
	d := DateSince(from, days)
	return fmt.Sprintf("\nDays:\n%-4d %s (%d)",
		days,
		// date
		d.Format("2006-01-02"),
		// days
		DaysBetween(to, d),
	)
}

func main() {
	// flags
	since := flag.String("s", "", "since last cig(yyyy-mm-dd)")
	days := flag.Int("d", 0, "goal in days")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, Usage)
	}
	// output
	fmt.Printf("cigs %s\n\n", Version)
	flag.Parse()
	if flag.NFlag() < 1 {
		flag.Usage()
	}
	// dates
	today := parseDate("")
	from := parseDate(*since)
	fmt.Printf("Today:\n%v\n", DaysBetween(from, today))
	if *days > 0 {
		fmt.Println(daysResult(from, today, *days))
	}
}

// parseDate returns parsed or current date.
func parseDate(s string) time.Time {
	d, err := time.Parse("2006-01-02", s)
	if err != nil {
		d = time.Now()
	}
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
}
