package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Usage holds command usage description.
const Usage = `Usage
	bima [options]... [values]...

Options:
	-h	height(m)
	-g	show goals
	-kg	weight(kg)
	-r	show normal range
`

// Version is command version.
const Version = "0.4.6"

func main() {
	// output
	fmt.Printf("bima %s\n\n", Version)
	height, weight, goals, rng, err := parseFlags()
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	current := BMI(*height, *weight)
	fmt.Print(printRow(*height, current))
	if *goals {
		fmt.Print("\nGoals:\n")
		for c, i := current, 0; i < 5; i++ {
			c = Goal(c)
			fmt.Print(printRow(*height, c))
		}
	}
	if *rng && *height > 0 {
		fmt.Print("\nNormal range:\n")
		for _, c := range Range() {
			fmt.Print(printRow(*height, c))
		}
	}
}

func parseFlags() (*float64, *float64, *bool, *bool, error) {
	// flags
	h := flag.Float64("h", 0, "height(m)")
	w := flag.Float64("kg", 0, "weight(kg)")
	g := flag.Bool("g", false, "show goals")
	rng := flag.Bool("r", false, "show normal range values")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, Usage)
	}
	flag.Parse()
	if flag.NFlag() < 1 {
		return nil, nil, nil, nil, errors.New("no flags")
	}

	return h, w, g, rng, nil
}

// printRow returns formated row string.
func printRow(height, bmi float64) string {
	return fmt.Sprintf("%4.1f (%4.2f kg)\n", bmi, Kg(height, bmi))
}
