package model

import "github.com/uptrace/bun"

type CharacterDeclinedname struct {
	Guid int `json:"guid,omitempty"`
	Genitive string `json:"genitive,omitempty"`
	Dative string `json:"dative,omitempty"`
	Accusative string `json:"accusative,omitempty"`
	Instrumental string `json:"instrumental,omitempty"`
	Prepositional string `json:"prepositional,omitempty"`
}

type CharacterDeclinednameSlice []CharacterDeclinedname

type DBCharacterDeclinedname struct {
	bun.BaseModel `bun:"table:character_declinedname,alias:character_declinedname"`
	Guid int `bun:"guid"`
	Genitive string `bun:"genitive"`
	Dative string `bun:"dative"`
	Accusative string `bun:"accusative"`
	Instrumental string `bun:"instrumental"`
	Prepositional string `bun:"prepositional"`
}

type DBCharacterDeclinednameSlice []DBCharacterDeclinedname

func (entry *CharacterDeclinedname) ToDB() *DBCharacterDeclinedname {
	if entry == nil {
		return nil
	}
	return &DBCharacterDeclinedname{
		Guid: entry.Guid,
		Genitive: entry.Genitive,
		Dative: entry.Dative,
		Accusative: entry.Accusative,
		Instrumental: entry.Instrumental,
		Prepositional: entry.Prepositional,
	}
}

func (entry *DBCharacterDeclinedname) ToWeb() *CharacterDeclinedname {
	if entry == nil {
		return nil
	}
	return &CharacterDeclinedname{
		Guid: entry.Guid,
		Genitive: entry.Genitive,
		Dative: entry.Dative,
		Accusative: entry.Accusative,
		Instrumental: entry.Instrumental,
		Prepositional: entry.Prepositional,
	}
}

func (data CharacterDeclinednameSlice) ToDB() DBCharacterDeclinednameSlice {
	result := make(DBCharacterDeclinednameSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterDeclinednameSlice) ToWeb() CharacterDeclinednameSlice {
	result := make(CharacterDeclinednameSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
