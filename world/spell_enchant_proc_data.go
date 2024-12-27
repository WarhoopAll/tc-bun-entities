package model

import "github.com/uptrace/bun"

type SpellEnchantProcData struct {
	EnchantID int `json:"enchantid,omitempty"`
	Chance float64 `json:"chance,omitempty"`
	ProcsPerMinute float64 `json:"procsperminute,omitempty"`
	HitMask int `json:"hitmask,omitempty"`
	AttributesMask int `json:"attributesmask,omitempty"`
}

type SpellEnchantProcDataSlice []SpellEnchantProcData

type DBSpellEnchantProcData struct {
	bun.BaseModel `bun:"table:spell_enchant_proc_data,alias:spell_enchant_proc_data"`
	EnchantID int `bun:"EnchantID"`
	Chance float64 `bun:"Chance"`
	ProcsPerMinute float64 `bun:"ProcsPerMinute"`
	HitMask int `bun:"HitMask"`
	AttributesMask int `bun:"AttributesMask"`
}

type DBSpellEnchantProcDataSlice []DBSpellEnchantProcData

func (entry *SpellEnchantProcData) ToDB() *DBSpellEnchantProcData {
	if entry == nil {
		return nil
	}
	return &DBSpellEnchantProcData{
		EnchantID: entry.EnchantID,
		Chance: entry.Chance,
		ProcsPerMinute: entry.ProcsPerMinute,
		HitMask: entry.HitMask,
		AttributesMask: entry.AttributesMask,
	}
}

func (entry *DBSpellEnchantProcData) ToWeb() *SpellEnchantProcData {
	if entry == nil {
		return nil
	}
	return &SpellEnchantProcData{
		EnchantID: entry.EnchantID,
		Chance: entry.Chance,
		ProcsPerMinute: entry.ProcsPerMinute,
		HitMask: entry.HitMask,
		AttributesMask: entry.AttributesMask,
	}
}

func (data SpellEnchantProcDataSlice) ToDB() DBSpellEnchantProcDataSlice {
	result := make(DBSpellEnchantProcDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellEnchantProcDataSlice) ToWeb() SpellEnchantProcDataSlice {
	result := make(SpellEnchantProcDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
