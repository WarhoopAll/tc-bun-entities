package model

import "github.com/uptrace/bun"

type SpellGroup struct {
	Id int `json:"id,omitempty"`
	SpellId int `json:"spell_id,omitempty"`
}

type SpellGroupSlice []SpellGroup

type DBSpellGroup struct {
	bun.BaseModel `bun:"table:spell_group,alias:spell_group"`
	Id int `bun:"id"`
	SpellId int `bun:"spell_id"`
}

type DBSpellGroupSlice []DBSpellGroup

func (entry *SpellGroup) ToDB() *DBSpellGroup {
	if entry == nil {
		return nil
	}
	return &DBSpellGroup{
		Id: entry.Id,
		SpellId: entry.SpellId,
	}
}

func (entry *DBSpellGroup) ToWeb() *SpellGroup {
	if entry == nil {
		return nil
	}
	return &SpellGroup{
		Id: entry.Id,
		SpellId: entry.SpellId,
	}
}

func (data SpellGroupSlice) ToDB() DBSpellGroupSlice {
	result := make(DBSpellGroupSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellGroupSlice) ToWeb() SpellGroupSlice {
	result := make(SpellGroupSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
