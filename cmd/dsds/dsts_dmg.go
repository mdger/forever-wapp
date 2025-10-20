package main

import (
	"fmt"

	"github.com/mdger/forever-wapp/dsts"
)

func main() {
	lastDamage := 0.0
	skillPwr, atk, def := 11, 2710, 1000
	fmt.Printf("atk %v def %v skillPwr %v\n", atk, def, skillPwr)
	variance := 0.045454545
	for i := 1; i < 100; i++ {
		damage := dsts.CalcDamage(i, skillPwr, atk, def)
		if lastDamage == 0.0 {
			fmt.Printf("[%v] %v dmg - %.2f %% increase \n", i, int(damage), 0.0)
		} else {
			fmt.Printf("[%v] %v dmg %v min %v max - %.2f%% dmg increase \n", i, int(damage), int(damage-(damage*variance)), int(damage+(damage*variance)), (damage-lastDamage)/lastDamage*100)
		}
		lastDamage = damage
	}
}
