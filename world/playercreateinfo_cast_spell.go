package model

import "github.com/uptrace/bun"

type PlayercreateinfoCastSpell struct {
	RaceMask int `json:"racemask,omitempty"`
	ClassMask int `json:"classmask,omitempty"`
	Spell int `json:"spell,omitempty"`
	Note string `json:"note,omitempty"`
}

type PlayercreateinfoCastSpellSlice []PlayercreateinfoCastSpell

type DBPlayercreateinfoCastSpell struct {
	bun.BaseModel `bun:"table:playercreateinfo_cast_spell,alias:playercreateinfo_cast_spell"`
	RaceMask int `bun:"raceMask"`
	ClassMask int `bun:"classMask"`
	Spell int `bun:"spell"`
	Note string `bun:"note"`
}

type DBPlayercreateinfoCastSpellSlice []DBPlayercreateinfoCastSpell

func (entry *PlayercreateinfoCastSpell) ToDB() *DBPlayercreateinfoCastSpell {
	if entry == nil {
		return nil
	}
	return &DBPlayercreateinfoCastSpell{
		RaceMask: entry.RaceMask,
		ClassMask: entry.ClassMask,
		Spell: entry.Spell,
		Note: entry.Note,
	}
}

func (entry *DBPlayercreateinfoCastSpell) ToWeb() *PlayercreateinfoCastSpell {
	if entry == nil {
		return nil
	}
	return &PlayercreateinfoCastSpell{
		RaceMask: entry.RaceMask,
		ClassMask: entry.ClassMask,
		Spell: entry.Spell,
		Note: entry.Note,
	}
}

func (data PlayercreateinfoCastSpellSlice) ToDB() DBPlayercreateinfoCastSpellSlice {
	result := make(DBPlayercreateinfoCastSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayercreateinfoCastSpellSlice) ToWeb() PlayercreateinfoCastSpellSlice {
	result := make(PlayercreateinfoCastSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
