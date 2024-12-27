package model

import "github.com/uptrace/bun"

type EventScripts struct {
	Id int `json:"id,omitempty"`
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

type EventScriptsSlice []EventScripts

type DBEventScripts struct {
	bun.BaseModel `bun:"table:event_scripts,alias:event_scripts"`
	Id int `bun:"id"`
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

type DBEventScriptsSlice []DBEventScripts

func (entry *EventScripts) ToDB() *DBEventScripts {
	if entry == nil {
		return nil
	}
	return &DBEventScripts{
		Id: entry.Id,
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

func (entry *DBEventScripts) ToWeb() *EventScripts {
	if entry == nil {
		return nil
	}
	return &EventScripts{
		Id: entry.Id,
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

func (data EventScriptsSlice) ToDB() DBEventScriptsSlice {
	result := make(DBEventScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBEventScriptsSlice) ToWeb() EventScriptsSlice {
	result := make(EventScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
