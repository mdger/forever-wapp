package main

import (
	"github.com/mdger/forever-wapp/mt2"
)

type Preset struct{}

func main() {
	numTotal := 0

	for i, clan := range mt2.Clans {
		for j, clan := range mt2.Clans {
			if i == j {
				continue
			}
			numTotal++
		}
	}
	// // TODO: calculate max number of primary clan + secondary clan with card combos
	// fmt.Sprintln("")
}
