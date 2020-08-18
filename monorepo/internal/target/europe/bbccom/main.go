package bbccom

import "github.com/dvrkps/dojo/monorepo/internal/target"

func New() target.Target {
	return target.Target{
		Key: "bbc_com",
	}
}
