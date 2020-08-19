package timecom

import (
	"fmt"

	"github.com/dvrkps/dojo/monorepo"
)

// New creates website.
func New() *monorepo.Website {
	return &monorepo.Website{
		Key:   "time_com",
		Index: index,
		Parse: parse,
	}
}

func index(w *monorepo.Website) (string, error) {
	return fmt.Sprintf("%v: index", w.Key), nil
}

func parse(w *monorepo.Website) (string, error) {
	return fmt.Sprintf("%v: parse", w.Key), nil
}
