package model

import "github.com/uptrace/bun"

type CreatureTemplateSpell struct {
	CreatureID int `json:"creatureid,omitempty"`
	Index int8 `json:"index,omitempty"`
	Spell int `json:"spell,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type CreatureTemplateSpellSlice []CreatureTemplateSpell

type DBCreatureTemplateSpell struct {
	bun.BaseModel `bun:"table:creature_template_spell,alias:creature_template_spell"`
	CreatureID int `bun:"CreatureID"`
	Index int8 `bun:"Index"`
	Spell int `bun:"Spell"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBCreatureTemplateSpellSlice []DBCreatureTemplateSpell

func (entry *CreatureTemplateSpell) ToDB() *DBCreatureTemplateSpell {
	if entry == nil {
		return nil
	}
	return &DBCreatureTemplateSpell{
		CreatureID: entry.CreatureID,
		Index: entry.Index,
		Spell: entry.Spell,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBCreatureTemplateSpell) ToWeb() *CreatureTemplateSpell {
	if entry == nil {
		return nil
	}
	return &CreatureTemplateSpell{
		CreatureID: entry.CreatureID,
		Index: entry.Index,
		Spell: entry.Spell,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data CreatureTemplateSpellSlice) ToDB() DBCreatureTemplateSpellSlice {
	result := make(DBCreatureTemplateSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureTemplateSpellSlice) ToWeb() CreatureTemplateSpellSlice {
	result := make(CreatureTemplateSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
