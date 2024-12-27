package model

import "github.com/uptrace/bun"

type PetSpell struct {
	Guid int `json:"guid,omitempty"`
	Spell int32 `json:"spell,omitempty"`
	Active int8 `json:"active,omitempty"`
}

type PetSpellSlice []PetSpell

type DBPetSpell struct {
	bun.BaseModel `bun:"table:pet_spell,alias:pet_spell"`
	Guid int `bun:"guid"`
	Spell int32 `bun:"spell"`
	Active int8 `bun:"active"`
}

type DBPetSpellSlice []DBPetSpell

func (entry *PetSpell) ToDB() *DBPetSpell {
	if entry == nil {
		return nil
	}
	return &DBPetSpell{
		Guid: entry.Guid,
		Spell: entry.Spell,
		Active: entry.Active,
	}
}

func (entry *DBPetSpell) ToWeb() *PetSpell {
	if entry == nil {
		return nil
	}
	return &PetSpell{
		Guid: entry.Guid,
		Spell: entry.Spell,
		Active: entry.Active,
	}
}

func (data PetSpellSlice) ToDB() DBPetSpellSlice {
	result := make(DBPetSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPetSpellSlice) ToWeb() PetSpellSlice {
	result := make(PetSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
