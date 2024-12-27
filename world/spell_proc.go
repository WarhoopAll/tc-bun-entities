package model

import "github.com/uptrace/bun"

type SpellProc struct {
	SpellId int `json:"spellid,omitempty"`
	SchoolMask int8 `json:"schoolmask,omitempty"`
	SpellFamilyName int16 `json:"spellfamilyname,omitempty"`
	SpellFamilyMask0 int `json:"spellfamilymask0,omitempty"`
	SpellFamilyMask1 int `json:"spellfamilymask1,omitempty"`
	SpellFamilyMask2 int `json:"spellfamilymask2,omitempty"`
	ProcFlags int `json:"procflags,omitempty"`
	SpellTypeMask int `json:"spelltypemask,omitempty"`
	SpellPhaseMask int `json:"spellphasemask,omitempty"`
	HitMask int `json:"hitmask,omitempty"`
	AttributesMask int `json:"attributesmask,omitempty"`
	DisableEffectsMask int `json:"disableeffectsmask,omitempty"`
	ProcsPerMinute float64 `json:"procsperminute,omitempty"`
	Chance float64 `json:"chance,omitempty"`
	Cooldown int `json:"cooldown,omitempty"`
	Charges int8 `json:"charges,omitempty"`
}

type SpellProcSlice []SpellProc

type DBSpellProc struct {
	bun.BaseModel `bun:"table:spell_proc,alias:spell_proc"`
	SpellId int `bun:"SpellId"`
	SchoolMask int8 `bun:"SchoolMask"`
	SpellFamilyName int16 `bun:"SpellFamilyName"`
	SpellFamilyMask0 int `bun:"SpellFamilyMask0"`
	SpellFamilyMask1 int `bun:"SpellFamilyMask1"`
	SpellFamilyMask2 int `bun:"SpellFamilyMask2"`
	ProcFlags int `bun:"ProcFlags"`
	SpellTypeMask int `bun:"SpellTypeMask"`
	SpellPhaseMask int `bun:"SpellPhaseMask"`
	HitMask int `bun:"HitMask"`
	AttributesMask int `bun:"AttributesMask"`
	DisableEffectsMask int `bun:"DisableEffectsMask"`
	ProcsPerMinute float64 `bun:"ProcsPerMinute"`
	Chance float64 `bun:"Chance"`
	Cooldown int `bun:"Cooldown"`
	Charges int8 `bun:"Charges"`
}

type DBSpellProcSlice []DBSpellProc

func (entry *SpellProc) ToDB() *DBSpellProc {
	if entry == nil {
		return nil
	}
	return &DBSpellProc{
		SpellId: entry.SpellId,
		SchoolMask: entry.SchoolMask,
		SpellFamilyName: entry.SpellFamilyName,
		SpellFamilyMask0: entry.SpellFamilyMask0,
		SpellFamilyMask1: entry.SpellFamilyMask1,
		SpellFamilyMask2: entry.SpellFamilyMask2,
		ProcFlags: entry.ProcFlags,
		SpellTypeMask: entry.SpellTypeMask,
		SpellPhaseMask: entry.SpellPhaseMask,
		HitMask: entry.HitMask,
		AttributesMask: entry.AttributesMask,
		DisableEffectsMask: entry.DisableEffectsMask,
		ProcsPerMinute: entry.ProcsPerMinute,
		Chance: entry.Chance,
		Cooldown: entry.Cooldown,
		Charges: entry.Charges,
	}
}

func (entry *DBSpellProc) ToWeb() *SpellProc {
	if entry == nil {
		return nil
	}
	return &SpellProc{
		SpellId: entry.SpellId,
		SchoolMask: entry.SchoolMask,
		SpellFamilyName: entry.SpellFamilyName,
		SpellFamilyMask0: entry.SpellFamilyMask0,
		SpellFamilyMask1: entry.SpellFamilyMask1,
		SpellFamilyMask2: entry.SpellFamilyMask2,
		ProcFlags: entry.ProcFlags,
		SpellTypeMask: entry.SpellTypeMask,
		SpellPhaseMask: entry.SpellPhaseMask,
		HitMask: entry.HitMask,
		AttributesMask: entry.AttributesMask,
		DisableEffectsMask: entry.DisableEffectsMask,
		ProcsPerMinute: entry.ProcsPerMinute,
		Chance: entry.Chance,
		Cooldown: entry.Cooldown,
		Charges: entry.Charges,
	}
}

func (data SpellProcSlice) ToDB() DBSpellProcSlice {
	result := make(DBSpellProcSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellProcSlice) ToWeb() SpellProcSlice {
	result := make(SpellProcSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
