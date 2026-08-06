package main

type locationArea struct {
	Id int
	Name string
	Game_Index int
	Encounter_Method_Rates []encounterMethodRate
	Location namedAPIResource
	Names []name
	Pokemon_Encounters []pokemonEncounter
}

type encounterMethodRate struct {
	Encounter_Method namedAPIResource
	Version_Details []encounterVersionDetails
}

type namedAPIResource struct {
	Name string
	Url string
}

type encounterVersionDetails struct {
	Rate int
	Version namedAPIResource
}

type name struct {
	Name string
	Language namedAPIResource
}

type pokemonEncounter struct {
	Pokemon namedAPIResource
	Version_Details []versionEncounterDetail
}

type versionEncounterDetail struct {
	Version namedAPIResource
	Max_Chance int
	Encounter_Details []encounter
}

type encounter struct {
	Min_Level int
	Max_Level int
	Condition_Values []namedAPIResource
	Chance int
	Method namedAPIResource
}
