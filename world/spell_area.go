package model

import "github.com/uptrace/bun"

type SpellArea struct {
	Spell int `json:"spell,omitempty"`
	Area int `json:"area,omitempty"`
	QuestStart int `json:"quest_start,omitempty"`
	QuestEnd int `json:"quest_end,omitempty"`
	AuraSpell int `json:"aura_spell,omitempty"`
	Racemask int `json:"racemask,omitempty"`
	Gender int8 `json:"gender,omitempty"`
	Autocast int8 `json:"autocast,omitempty"`
	QuestStartStatus int `json:"quest_start_status,omitempty"`
	QuestEndStatus int `json:"quest_end_status,omitempty"`
}

type SpellAreaSlice []SpellArea

type DBSpellArea struct {
	bun.BaseModel `bun:"table:spell_area,alias:spell_area"`
	Spell int `bun:"spell"`
	Area int `bun:"area"`
	QuestStart int `bun:"quest_start"`
	QuestEnd int `bun:"quest_end"`
	AuraSpell int `bun:"aura_spell"`
	Racemask int `bun:"racemask"`
	Gender int8 `bun:"gender"`
	Autocast int8 `bun:"autocast"`
	QuestStartStatus int `bun:"quest_start_status"`
	QuestEndStatus int `bun:"quest_end_status"`
}

type DBSpellAreaSlice []DBSpellArea

func (entry *SpellArea) ToDB() *DBSpellArea {
	if entry == nil {
		return nil
	}
	return &DBSpellArea{
		Spell: entry.Spell,
		Area: entry.Area,
		QuestStart: entry.QuestStart,
		QuestEnd: entry.QuestEnd,
		AuraSpell: entry.AuraSpell,
		Racemask: entry.Racemask,
		Gender: entry.Gender,
		Autocast: entry.Autocast,
		QuestStartStatus: entry.QuestStartStatus,
		QuestEndStatus: entry.QuestEndStatus,
	}
}

func (entry *DBSpellArea) ToWeb() *SpellArea {
	if entry == nil {
		return nil
	}
	return &SpellArea{
		Spell: entry.Spell,
		Area: entry.Area,
		QuestStart: entry.QuestStart,
		QuestEnd: entry.QuestEnd,
		AuraSpell: entry.AuraSpell,
		Racemask: entry.Racemask,
		Gender: entry.Gender,
		Autocast: entry.Autocast,
		QuestStartStatus: entry.QuestStartStatus,
		QuestEndStatus: entry.QuestEndStatus,
	}
}

func (data SpellAreaSlice) ToDB() DBSpellAreaSlice {
	result := make(DBSpellAreaSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellAreaSlice) ToWeb() SpellAreaSlice {
	result := make(SpellAreaSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
