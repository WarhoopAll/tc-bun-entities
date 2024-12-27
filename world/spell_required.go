package model

import "github.com/uptrace/bun"

type SpellRequired struct {
	SpellId int `json:"spell_id,omitempty"`
	ReqSpell int `json:"req_spell,omitempty"`
}

type SpellRequiredSlice []SpellRequired

type DBSpellRequired struct {
	bun.BaseModel `bun:"table:spell_required,alias:spell_required"`
	SpellId int `bun:"spell_id"`
	ReqSpell int `bun:"req_spell"`
}

type DBSpellRequiredSlice []DBSpellRequired

func (entry *SpellRequired) ToDB() *DBSpellRequired {
	if entry == nil {
		return nil
	}
	return &DBSpellRequired{
		SpellId: entry.SpellId,
		ReqSpell: entry.ReqSpell,
	}
}

func (entry *DBSpellRequired) ToWeb() *SpellRequired {
	if entry == nil {
		return nil
	}
	return &SpellRequired{
		SpellId: entry.SpellId,
		ReqSpell: entry.ReqSpell,
	}
}

func (data SpellRequiredSlice) ToDB() DBSpellRequiredSlice {
	result := make(DBSpellRequiredSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellRequiredSlice) ToWeb() SpellRequiredSlice {
	result := make(SpellRequiredSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
