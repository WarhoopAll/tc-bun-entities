package model

import "github.com/uptrace/bun"

type SpellThreat struct {
	Entry int `json:"entry,omitempty"`
	FlatMod int `json:"flatmod,omitempty"`
	PctMod float64 `json:"pctmod,omitempty"`
	ApPctMod float64 `json:"appctmod,omitempty"`
}

type SpellThreatSlice []SpellThreat

type DBSpellThreat struct {
	bun.BaseModel `bun:"table:spell_threat,alias:spell_threat"`
	Entry int `bun:"entry"`
	FlatMod int `bun:"flatMod"`
	PctMod float64 `bun:"pctMod"`
	ApPctMod float64 `bun:"apPctMod"`
}

type DBSpellThreatSlice []DBSpellThreat

func (entry *SpellThreat) ToDB() *DBSpellThreat {
	if entry == nil {
		return nil
	}
	return &DBSpellThreat{
		Entry: entry.Entry,
		FlatMod: entry.FlatMod,
		PctMod: entry.PctMod,
		ApPctMod: entry.ApPctMod,
	}
}

func (entry *DBSpellThreat) ToWeb() *SpellThreat {
	if entry == nil {
		return nil
	}
	return &SpellThreat{
		Entry: entry.Entry,
		FlatMod: entry.FlatMod,
		PctMod: entry.PctMod,
		ApPctMod: entry.ApPctMod,
	}
}

func (data SpellThreatSlice) ToDB() DBSpellThreatSlice {
	result := make(DBSpellThreatSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellThreatSlice) ToWeb() SpellThreatSlice {
	result := make(SpellThreatSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
