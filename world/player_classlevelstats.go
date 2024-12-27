package model

import "github.com/uptrace/bun"

type PlayerClasslevelstats struct {
	Class int8 `json:"class,omitempty"`
	Level int8 `json:"level,omitempty"`
	Basehp int16 `json:"basehp,omitempty"`
	Basemana int16 `json:"basemana,omitempty"`
}

type PlayerClasslevelstatsSlice []PlayerClasslevelstats

type DBPlayerClasslevelstats struct {
	bun.BaseModel `bun:"table:player_classlevelstats,alias:player_classlevelstats"`
	Class int8 `bun:"class"`
	Level int8 `bun:"level"`
	Basehp int16 `bun:"basehp"`
	Basemana int16 `bun:"basemana"`
}

type DBPlayerClasslevelstatsSlice []DBPlayerClasslevelstats

func (entry *PlayerClasslevelstats) ToDB() *DBPlayerClasslevelstats {
	if entry == nil {
		return nil
	}
	return &DBPlayerClasslevelstats{
		Class: entry.Class,
		Level: entry.Level,
		Basehp: entry.Basehp,
		Basemana: entry.Basemana,
	}
}

func (entry *DBPlayerClasslevelstats) ToWeb() *PlayerClasslevelstats {
	if entry == nil {
		return nil
	}
	return &PlayerClasslevelstats{
		Class: entry.Class,
		Level: entry.Level,
		Basehp: entry.Basehp,
		Basemana: entry.Basemana,
	}
}

func (data PlayerClasslevelstatsSlice) ToDB() DBPlayerClasslevelstatsSlice {
	result := make(DBPlayerClasslevelstatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerClasslevelstatsSlice) ToWeb() PlayerClasslevelstatsSlice {
	result := make(PlayerClasslevelstatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
