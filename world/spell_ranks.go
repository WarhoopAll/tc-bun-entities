package model

import "github.com/uptrace/bun"

type SpellRanks struct {
	FirstSpellId int `json:"first_spell_id,omitempty"`
	SpellId int `json:"spell_id,omitempty"`
	Rank int8 `json:"rank,omitempty"`
}

type SpellRanksSlice []SpellRanks

type DBSpellRanks struct {
	bun.BaseModel `bun:"table:spell_ranks,alias:spell_ranks"`
	FirstSpellId int `bun:"first_spell_id"`
	SpellId int `bun:"spell_id"`
	Rank int8 `bun:"rank"`
}

type DBSpellRanksSlice []DBSpellRanks

func (entry *SpellRanks) ToDB() *DBSpellRanks {
	if entry == nil {
		return nil
	}
	return &DBSpellRanks{
		FirstSpellId: entry.FirstSpellId,
		SpellId: entry.SpellId,
		Rank: entry.Rank,
	}
}

func (entry *DBSpellRanks) ToWeb() *SpellRanks {
	if entry == nil {
		return nil
	}
	return &SpellRanks{
		FirstSpellId: entry.FirstSpellId,
		SpellId: entry.SpellId,
		Rank: entry.Rank,
	}
}

func (data SpellRanksSlice) ToDB() DBSpellRanksSlice {
	result := make(DBSpellRanksSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellRanksSlice) ToWeb() SpellRanksSlice {
	result := make(SpellRanksSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
