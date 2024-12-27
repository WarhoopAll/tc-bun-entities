package model

import "github.com/uptrace/bun"

type CharacterSpell struct {
	Guid int `json:"guid,omitempty"`
	Spell int32 `json:"spell,omitempty"`
	Active int8 `json:"active,omitempty"`
	Disabled int8 `json:"disabled,omitempty"`
}

type CharacterSpellSlice []CharacterSpell

type DBCharacterSpell struct {
	bun.BaseModel `bun:"table:character_spell,alias:character_spell"`
	Guid int `bun:"guid"`
	Spell int32 `bun:"spell"`
	Active int8 `bun:"active"`
	Disabled int8 `bun:"disabled"`
}

type DBCharacterSpellSlice []DBCharacterSpell

func (entry *CharacterSpell) ToDB() *DBCharacterSpell {
	if entry == nil {
		return nil
	}
	return &DBCharacterSpell{
		Guid: entry.Guid,
		Spell: entry.Spell,
		Active: entry.Active,
		Disabled: entry.Disabled,
	}
}

func (entry *DBCharacterSpell) ToWeb() *CharacterSpell {
	if entry == nil {
		return nil
	}
	return &CharacterSpell{
		Guid: entry.Guid,
		Spell: entry.Spell,
		Active: entry.Active,
		Disabled: entry.Disabled,
	}
}

func (data CharacterSpellSlice) ToDB() DBCharacterSpellSlice {
	result := make(DBCharacterSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterSpellSlice) ToWeb() CharacterSpellSlice {
	result := make(CharacterSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
