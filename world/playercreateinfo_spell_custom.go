package model

import "github.com/uptrace/bun"

type PlayercreateinfoSpellCustom struct {
	Racemask int `json:"racemask,omitempty"`
	Classmask int `json:"classmask,omitempty"`
	Spell int `json:"spell,omitempty"`
	Note string `json:"note,omitempty"`
}

type PlayercreateinfoSpellCustomSlice []PlayercreateinfoSpellCustom

type DBPlayercreateinfoSpellCustom struct {
	bun.BaseModel `bun:"table:playercreateinfo_spell_custom,alias:playercreateinfo_spell_custom"`
	Racemask int `bun:"racemask"`
	Classmask int `bun:"classmask"`
	Spell int `bun:"Spell"`
	Note string `bun:"Note"`
}

type DBPlayercreateinfoSpellCustomSlice []DBPlayercreateinfoSpellCustom

func (entry *PlayercreateinfoSpellCustom) ToDB() *DBPlayercreateinfoSpellCustom {
	if entry == nil {
		return nil
	}
	return &DBPlayercreateinfoSpellCustom{
		Racemask: entry.Racemask,
		Classmask: entry.Classmask,
		Spell: entry.Spell,
		Note: entry.Note,
	}
}

func (entry *DBPlayercreateinfoSpellCustom) ToWeb() *PlayercreateinfoSpellCustom {
	if entry == nil {
		return nil
	}
	return &PlayercreateinfoSpellCustom{
		Racemask: entry.Racemask,
		Classmask: entry.Classmask,
		Spell: entry.Spell,
		Note: entry.Note,
	}
}

func (data PlayercreateinfoSpellCustomSlice) ToDB() DBPlayercreateinfoSpellCustomSlice {
	result := make(DBPlayercreateinfoSpellCustomSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayercreateinfoSpellCustomSlice) ToWeb() PlayercreateinfoSpellCustomSlice {
	result := make(PlayercreateinfoSpellCustomSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
