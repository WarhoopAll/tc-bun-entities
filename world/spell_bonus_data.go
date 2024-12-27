package model

import "github.com/uptrace/bun"

type SpellBonusData struct {
	Entry int `json:"entry,omitempty"`
	DirectBonus float64 `json:"direct_bonus,omitempty"`
	DotBonus float64 `json:"dot_bonus,omitempty"`
	ApBonus float64 `json:"ap_bonus,omitempty"`
	ApDotBonus float64 `json:"ap_dot_bonus,omitempty"`
	Comments string `json:"comments,omitempty"`
}

type SpellBonusDataSlice []SpellBonusData

type DBSpellBonusData struct {
	bun.BaseModel `bun:"table:spell_bonus_data,alias:spell_bonus_data"`
	Entry int `bun:"entry"`
	DirectBonus float64 `bun:"direct_bonus"`
	DotBonus float64 `bun:"dot_bonus"`
	ApBonus float64 `bun:"ap_bonus"`
	ApDotBonus float64 `bun:"ap_dot_bonus"`
	Comments string `bun:"comments"`
}

type DBSpellBonusDataSlice []DBSpellBonusData

func (entry *SpellBonusData) ToDB() *DBSpellBonusData {
	if entry == nil {
		return nil
	}
	return &DBSpellBonusData{
		Entry: entry.Entry,
		DirectBonus: entry.DirectBonus,
		DotBonus: entry.DotBonus,
		ApBonus: entry.ApBonus,
		ApDotBonus: entry.ApDotBonus,
		Comments: entry.Comments,
	}
}

func (entry *DBSpellBonusData) ToWeb() *SpellBonusData {
	if entry == nil {
		return nil
	}
	return &SpellBonusData{
		Entry: entry.Entry,
		DirectBonus: entry.DirectBonus,
		DotBonus: entry.DotBonus,
		ApBonus: entry.ApBonus,
		ApDotBonus: entry.ApDotBonus,
		Comments: entry.Comments,
	}
}

func (data SpellBonusDataSlice) ToDB() DBSpellBonusDataSlice {
	result := make(DBSpellBonusDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellBonusDataSlice) ToWeb() SpellBonusDataSlice {
	result := make(SpellBonusDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
