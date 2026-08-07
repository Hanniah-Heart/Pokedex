package main

type pokemon struct {
	Id	int
	Name	string
	Base_Experience	int
	Height	int
	Is_Default	bool
	Order	int
	Weight	int
	Abilities	[]pokemonAbility
	Forms	[]namedAPIResource
	Game_Indices	[]versionGameIndex
	Held_Items	[]pokemonHeldItem
	Location_Area_Encounters	string
	Moves	[]pokemonMove
	Past_Types	[]pokemonTypePast
	Past_Abilities	[]pokemonAbilityPast
	Past_Stats	[]pokemonStatPast
	Sprites	pokemonSprites
	Cries	pokemonCries
	Species	namedAPIResource
	Stats	[]pokemonStat
	Types	[]pokemonType
}

type pokemonAbility struct {
	Is_Hidden	bool
	Slot	int
	Ability	namedAPIResource
}

type versionGameIndex struct {
	Game_Index	int
	Version	namedAPIResource
}

type pokemonHeldItem struct {
	Item namedAPIResource
	Version_Details	[]pokemonHeldItemVersion
}

type pokemonHeldItemVersion struct {
	Version	namedAPIResource
	Rarity int
}

type pokemonMove struct {
	Move	namedAPIResource
	Version_Group_Details	[]pokemonMoveVersion
}

type pokemonMoveVersion struct {
	Move_Learn_Method	namedAPIResource
	Version_Group	namedAPIResource
	Level_Learned_At	int
	Order	int
}

type pokemonTypePast struct {
	Generation	namedAPIResource
	Types	[]pokemonType
}

type pokemonType struct {
	Slot	int
	Type	namedAPIResource
}

type pokemonAbilityPast struct {
	Generation	namedAPIResource
	Abilities	[]pokemonAbility
}

type pokemonStatPast struct {
	Generation	namedAPIResource
	Stats	[]pokemonStat
}

type pokemonStat struct {
	Stat namedAPIResource
	Effort	int
	Base_Stat	int
}

type pokemonSprites struct {
	Front_Default string
	Front_Shiny string
	Front_Female	string
	Front_Shiny_Female	string
	Back_Default	string
	Back_Shiny	string
	Back_Female	string
	Back_Shiny_Female	string
}

type pokemonCries struct {
	Latest string
	Legacy string
}
