/*
Package dsts contains code related to Digimon Story Time Stranger
*/
package dsts

func CalcDamage(level, skillPwr, atk, def int) float64 {
	fAtk := float64(atk)
	fLevel := float64(level)
	fSkillPwr := float64(skillPwr)
	fDef := float64(def)
	return ((fAtk * 80 * (1 + fLevel*fLevel/9801)) * fSkillPwr / (fDef * 35)) * (1 + fLevel*(1+fLevel/30)/100)
}
