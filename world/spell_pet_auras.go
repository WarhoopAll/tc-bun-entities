package model

import "github.com/uptrace/bun"

type SpellPetAuras struct {
	Spell int `json:"spell,omitempty"`
	EffectId int8 `json:"effectid,omitempty"`
	Pet int `json:"pet,omitempty"`
	Aura int `json:"aura,omitempty"`
}

type SpellPetAurasSlice []SpellPetAuras

type DBSpellPetAuras struct {
	bun.BaseModel `bun:"table:spell_pet_auras,alias:spell_pet_auras"`
	Spell int `bun:"spell"`
	EffectId int8 `bun:"effectId"`
	Pet int `bun:"pet"`
	Aura int `bun:"aura"`
}

type DBSpellPetAurasSlice []DBSpellPetAuras

func (entry *SpellPetAuras) ToDB() *DBSpellPetAuras {
	if entry == nil {
		return nil
	}
	return &DBSpellPetAuras{
		Spell: entry.Spell,
		EffectId: entry.EffectId,
		Pet: entry.Pet,
		Aura: entry.Aura,
	}
}

func (entry *DBSpellPetAuras) ToWeb() *SpellPetAuras {
	if entry == nil {
		return nil
	}
	return &SpellPetAuras{
		Spell: entry.Spell,
		EffectId: entry.EffectId,
		Pet: entry.Pet,
		Aura: entry.Aura,
	}
}

func (data SpellPetAurasSlice) ToDB() DBSpellPetAurasSlice {
	result := make(DBSpellPetAurasSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellPetAurasSlice) ToWeb() SpellPetAurasSlice {
	result := make(SpellPetAurasSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
