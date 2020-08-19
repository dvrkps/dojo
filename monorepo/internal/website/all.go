package website

import (
	"fmt"

	"github.com/dvrkps/dojo/monorepo"
	"github.com/dvrkps/dojo/monorepo/internal/website/europe/bbccom"
	"github.com/dvrkps/dojo/monorepo/internal/website/usa/timecom"
)

func All() (map[string]*monorepo.Website, error) {
	ws := []*monorepo.Website{
		bbccom.New(),
		timecom.New(),
	}

	all := make(map[string]*monorepo.Website, len(ws))
	for _, w := range ws {
		if _, ok := all[w.Key]; ok {
			return nil, fmt.Errorf("duplicate key: %q", w.Key)
		}

		all[w.Key] = w
	}

	return all, nil
}
