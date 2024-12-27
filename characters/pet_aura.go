package model

import "github.com/uptrace/bun"

type PetAura struct {
	Guid int `json:"guid,omitempty"`
	CasterGuid int `json:"casterguid,omitempty"`
	Spell int32 `json:"spell,omitempty"`
	EffectMask int8 `json:"effectmask,omitempty"`
	RecalculateMask int8 `json:"recalculatemask,omitempty"`
	StackCount int8 `json:"stackcount,omitempty"`
	Amount0 int32 `json:"amount0,omitempty"`
	Amount1 int32 `json:"amount1,omitempty"`
	Amount2 int32 `json:"amount2,omitempty"`
	BaseAmount0 int32 `json:"base_amount0,omitempty"`
	BaseAmount1 int32 `json:"base_amount1,omitempty"`
	BaseAmount2 int32 `json:"base_amount2,omitempty"`
	MaxDuration int `json:"maxduration,omitempty"`
	RemainTime int `json:"remaintime,omitempty"`
	RemainCharges int8 `json:"remaincharges,omitempty"`
	CritChance float64 `json:"critchance,omitempty"`
	ApplyResilience int8 `json:"applyresilience,omitempty"`
}

type PetAuraSlice []PetAura

type DBPetAura struct {
	bun.BaseModel `bun:"table:pet_aura,alias:pet_aura"`
	Guid int `bun:"guid"`
	CasterGuid int `bun:"casterGuid"`
	Spell int32 `bun:"spell"`
	EffectMask int8 `bun:"effectMask"`
	RecalculateMask int8 `bun:"recalculateMask"`
	StackCount int8 `bun:"stackCount"`
	Amount0 int32 `bun:"amount0"`
	Amount1 int32 `bun:"amount1"`
	Amount2 int32 `bun:"amount2"`
	BaseAmount0 int32 `bun:"base_amount0"`
	BaseAmount1 int32 `bun:"base_amount1"`
	BaseAmount2 int32 `bun:"base_amount2"`
	MaxDuration int `bun:"maxDuration"`
	RemainTime int `bun:"remainTime"`
	RemainCharges int8 `bun:"remainCharges"`
	CritChance float64 `bun:"critChance"`
	ApplyResilience int8 `bun:"applyResilience"`
}

type DBPetAuraSlice []DBPetAura

func (entry *PetAura) ToDB() *DBPetAura {
	if entry == nil {
		return nil
	}
	return &DBPetAura{
		Guid: entry.Guid,
		CasterGuid: entry.CasterGuid,
		Spell: entry.Spell,
		EffectMask: entry.EffectMask,
		RecalculateMask: entry.RecalculateMask,
		StackCount: entry.StackCount,
		Amount0: entry.Amount0,
		Amount1: entry.Amount1,
		Amount2: entry.Amount2,
		BaseAmount0: entry.BaseAmount0,
		BaseAmount1: entry.BaseAmount1,
		BaseAmount2: entry.BaseAmount2,
		MaxDuration: entry.MaxDuration,
		RemainTime: entry.RemainTime,
		RemainCharges: entry.RemainCharges,
		CritChance: entry.CritChance,
		ApplyResilience: entry.ApplyResilience,
	}
}

func (entry *DBPetAura) ToWeb() *PetAura {
	if entry == nil {
		return nil
	}
	return &PetAura{
		Guid: entry.Guid,
		CasterGuid: entry.CasterGuid,
		Spell: entry.Spell,
		EffectMask: entry.EffectMask,
		RecalculateMask: entry.RecalculateMask,
		StackCount: entry.StackCount,
		Amount0: entry.Amount0,
		Amount1: entry.Amount1,
		Amount2: entry.Amount2,
		BaseAmount0: entry.BaseAmount0,
		BaseAmount1: entry.BaseAmount1,
		BaseAmount2: entry.BaseAmount2,
		MaxDuration: entry.MaxDuration,
		RemainTime: entry.RemainTime,
		RemainCharges: entry.RemainCharges,
		CritChance: entry.CritChance,
		ApplyResilience: entry.ApplyResilience,
	}
}

func (data PetAuraSlice) ToDB() DBPetAuraSlice {
	result := make(DBPetAuraSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPetAuraSlice) ToWeb() PetAuraSlice {
	result := make(PetAuraSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
