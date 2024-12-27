package model

import "github.com/uptrace/bun"

type SpellScripts struct {
	Id int `json:"id,omitempty"`
	EffIndex int8 `json:"effindex,omitempty"`
	Delay int `json:"delay,omitempty"`
	Command int `json:"command,omitempty"`
	Datalong int `json:"datalong,omitempty"`
	Datalong2 int `json:"datalong2,omitempty"`
	Dataint int `json:"dataint,omitempty"`
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
	O float64 `json:"o,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type SpellScriptsSlice []SpellScripts

type DBSpellScripts struct {
	bun.BaseModel `bun:"table:spell_scripts,alias:spell_scripts"`
	Id int `bun:"id"`
	EffIndex int8 `bun:"effIndex"`
	Delay int `bun:"delay"`
	Command int `bun:"command"`
	Datalong int `bun:"datalong"`
	Datalong2 int `bun:"datalong2"`
	Dataint int `bun:"dataint"`
	X float64 `bun:"x"`
	Y float64 `bun:"y"`
	Z float64 `bun:"z"`
	O float64 `bun:"o"`
	Comment string `bun:"Comment"`
}

type DBSpellScriptsSlice []DBSpellScripts

func (entry *SpellScripts) ToDB() *DBSpellScripts {
	if entry == nil {
		return nil
	}
	return &DBSpellScripts{
		Id: entry.Id,
		EffIndex: entry.EffIndex,
		Delay: entry.Delay,
		Command: entry.Command,
		Datalong: entry.Datalong,
		Datalong2: entry.Datalong2,
		Dataint: entry.Dataint,
		X: entry.X,
		Y: entry.Y,
		Z: entry.Z,
		O: entry.O,
		Comment: entry.Comment,
	}
}

func (entry *DBSpellScripts) ToWeb() *SpellScripts {
	if entry == nil {
		return nil
	}
	return &SpellScripts{
		Id: entry.Id,
		EffIndex: entry.EffIndex,
		Delay: entry.Delay,
		Command: entry.Command,
		Datalong: entry.Datalong,
		Datalong2: entry.Datalong2,
		Dataint: entry.Dataint,
		X: entry.X,
		Y: entry.Y,
		Z: entry.Z,
		O: entry.O,
		Comment: entry.Comment,
	}
}

func (data SpellScriptsSlice) ToDB() DBSpellScriptsSlice {
	result := make(DBSpellScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellScriptsSlice) ToWeb() SpellScriptsSlice {
	result := make(SpellScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
