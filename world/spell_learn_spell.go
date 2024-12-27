package model

import "github.com/uptrace/bun"

type SpellLearnSpell struct {
	Entry int `json:"entry,omitempty"`
	SpellID int `json:"spellid,omitempty"`
	Active int8 `json:"active,omitempty"`
}

type SpellLearnSpellSlice []SpellLearnSpell

type DBSpellLearnSpell struct {
	bun.BaseModel `bun:"table:spell_learn_spell,alias:spell_learn_spell"`
	Entry int `bun:"entry"`
	SpellID int `bun:"SpellID"`
	Active int8 `bun:"Active"`
}

type DBSpellLearnSpellSlice []DBSpellLearnSpell

func (entry *SpellLearnSpell) ToDB() *DBSpellLearnSpell {
	if entry == nil {
		return nil
	}
	return &DBSpellLearnSpell{
		Entry: entry.Entry,
		SpellID: entry.SpellID,
		Active: entry.Active,
	}
}

func (entry *DBSpellLearnSpell) ToWeb() *SpellLearnSpell {
	if entry == nil {
		return nil
	}
	return &SpellLearnSpell{
		Entry: entry.Entry,
		SpellID: entry.SpellID,
		Active: entry.Active,
	}
}

func (data SpellLearnSpellSlice) ToDB() DBSpellLearnSpellSlice {
	result := make(DBSpellLearnSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellLearnSpellSlice) ToWeb() SpellLearnSpellSlice {
	result := make(SpellLearnSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
