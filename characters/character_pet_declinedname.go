package model

import "github.com/uptrace/bun"

type CharacterPetDeclinedname struct {
	Id int `json:"id,omitempty"`
	Owner int `json:"owner,omitempty"`
	Genitive string `json:"genitive,omitempty"`
	Dative string `json:"dative,omitempty"`
	Accusative string `json:"accusative,omitempty"`
	Instrumental string `json:"instrumental,omitempty"`
	Prepositional string `json:"prepositional,omitempty"`
}

type CharacterPetDeclinednameSlice []CharacterPetDeclinedname

type DBCharacterPetDeclinedname struct {
	bun.BaseModel `bun:"table:character_pet_declinedname,alias:character_pet_declinedname"`
	Id int `bun:"id"`
	Owner int `bun:"owner"`
	Genitive string `bun:"genitive"`
	Dative string `bun:"dative"`
	Accusative string `bun:"accusative"`
	Instrumental string `bun:"instrumental"`
	Prepositional string `bun:"prepositional"`
}

type DBCharacterPetDeclinednameSlice []DBCharacterPetDeclinedname

func (entry *CharacterPetDeclinedname) ToDB() *DBCharacterPetDeclinedname {
	if entry == nil {
		return nil
	}
	return &DBCharacterPetDeclinedname{
		Id: entry.Id,
		Owner: entry.Owner,
		Genitive: entry.Genitive,
		Dative: entry.Dative,
		Accusative: entry.Accusative,
		Instrumental: entry.Instrumental,
		Prepositional: entry.Prepositional,
	}
}

func (entry *DBCharacterPetDeclinedname) ToWeb() *CharacterPetDeclinedname {
	if entry == nil {
		return nil
	}
	return &CharacterPetDeclinedname{
		Id: entry.Id,
		Owner: entry.Owner,
		Genitive: entry.Genitive,
		Dative: entry.Dative,
		Accusative: entry.Accusative,
		Instrumental: entry.Instrumental,
		Prepositional: entry.Prepositional,
	}
}

func (data CharacterPetDeclinednameSlice) ToDB() DBCharacterPetDeclinednameSlice {
	result := make(DBCharacterPetDeclinednameSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterPetDeclinednameSlice) ToWeb() CharacterPetDeclinednameSlice {
	result := make(CharacterPetDeclinednameSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
