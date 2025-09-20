/*
Package mt2 contains code related to interfacing with mt2 related code
*/
package mt2

type Clan struct {
	Name      string
	Champion1 string
	Champion2 string
}

var Clans = []Clan{
	{"Banished", "Fel", "Talos"},
	{"Pyreborne", "Lord Fenix", "Lady Gilda"},
	{"Luna Coven", "Ekka", "Arduhn"},
	{"Underlegion", "Bolete the Guillotine", "Madame Lionsmane"},
	{"Lazarus League", "Orechi", "Baron Grael"},
	{"Hellhorned", "Hornbreaker Prince", "Shardtail Queen"},
	{"Awoken", "The Sentient", "Wyldenten"},
	{"Stygian Guard", "Tethys Titansbane", "Solgard the Martyr"},
	{"Umbra", "Penumbra", "Primordium"},
	{"Melting Remnant", "Rector Flicker", "Little Fade"},
}
