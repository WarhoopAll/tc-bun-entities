package model

import "github.com/uptrace/bun"

type SpellLinkedSpell struct {
	SpellTrigger int `json:"spell_trigger,omitempty"`
	SpellEffect int `json:"spell_effect,omitempty"`
	Type int8 `json:"type,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type SpellLinkedSpellSlice []SpellLinkedSpell

type DBSpellLinkedSpell struct {
	bun.BaseModel `bun:"table:spell_linked_spell,alias:spell_linked_spell"`
	SpellTrigger int `bun:"spell_trigger"`
	SpellEffect int `bun:"spell_effect"`
	Type int8 `bun:"type"`
	Comment string `bun:"comment"`
}

type DBSpellLinkedSpellSlice []DBSpellLinkedSpell

func (entry *SpellLinkedSpell) ToDB() *DBSpellLinkedSpell {
	if entry == nil {
		return nil
	}
	return &DBSpellLinkedSpell{
		SpellTrigger: entry.SpellTrigger,
		SpellEffect: entry.SpellEffect,
		Type: entry.Type,
		Comment: entry.Comment,
	}
}

func (entry *DBSpellLinkedSpell) ToWeb() *SpellLinkedSpell {
	if entry == nil {
		return nil
	}
	return &SpellLinkedSpell{
		SpellTrigger: entry.SpellTrigger,
		SpellEffect: entry.SpellEffect,
		Type: entry.Type,
		Comment: entry.Comment,
	}
}

func (data SpellLinkedSpellSlice) ToDB() DBSpellLinkedSpellSlice {
	result := make(DBSpellLinkedSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellLinkedSpellSlice) ToWeb() SpellLinkedSpellSlice {
	result := make(SpellLinkedSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
