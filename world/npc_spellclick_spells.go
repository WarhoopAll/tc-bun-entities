package model

import "github.com/uptrace/bun"

type NpcSpellclickSpells struct {
	NpcEntry int `json:"npc_entry,omitempty"`
	SpellId int `json:"spell_id,omitempty"`
	CastFlags int8 `json:"cast_flags,omitempty"`
	UserType int16 `json:"user_type,omitempty"`
}

type NpcSpellclickSpellsSlice []NpcSpellclickSpells

type DBNpcSpellclickSpells struct {
	bun.BaseModel `bun:"table:npc_spellclick_spells,alias:npc_spellclick_spells"`
	NpcEntry int `bun:"npc_entry"`
	SpellId int `bun:"spell_id"`
	CastFlags int8 `bun:"cast_flags"`
	UserType int16 `bun:"user_type"`
}

type DBNpcSpellclickSpellsSlice []DBNpcSpellclickSpells

func (entry *NpcSpellclickSpells) ToDB() *DBNpcSpellclickSpells {
	if entry == nil {
		return nil
	}
	return &DBNpcSpellclickSpells{
		NpcEntry: entry.NpcEntry,
		SpellId: entry.SpellId,
		CastFlags: entry.CastFlags,
		UserType: entry.UserType,
	}
}

func (entry *DBNpcSpellclickSpells) ToWeb() *NpcSpellclickSpells {
	if entry == nil {
		return nil
	}
	return &NpcSpellclickSpells{
		NpcEntry: entry.NpcEntry,
		SpellId: entry.SpellId,
		CastFlags: entry.CastFlags,
		UserType: entry.UserType,
	}
}

func (data NpcSpellclickSpellsSlice) ToDB() DBNpcSpellclickSpellsSlice {
	result := make(DBNpcSpellclickSpellsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBNpcSpellclickSpellsSlice) ToWeb() NpcSpellclickSpellsSlice {
	result := make(NpcSpellclickSpellsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
