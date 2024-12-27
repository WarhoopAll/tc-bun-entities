package model

import "github.com/uptrace/bun"

type PlayerLevelstats struct {
	Race int8 `json:"race,omitempty"`
	Class int8 `json:"class,omitempty"`
	Level int8 `json:"level,omitempty"`
	Str int8 `json:"str,omitempty"`
	Agi int8 `json:"agi,omitempty"`
	Sta int8 `json:"sta,omitempty"`
	Inte int8 `json:"inte,omitempty"`
	Spi int8 `json:"spi,omitempty"`
}

type PlayerLevelstatsSlice []PlayerLevelstats

type DBPlayerLevelstats struct {
	bun.BaseModel `bun:"table:player_levelstats,alias:player_levelstats"`
	Race int8 `bun:"race"`
	Class int8 `bun:"class"`
	Level int8 `bun:"level"`
	Str int8 `bun:"str"`
	Agi int8 `bun:"agi"`
	Sta int8 `bun:"sta"`
	Inte int8 `bun:"inte"`
	Spi int8 `bun:"spi"`
}

type DBPlayerLevelstatsSlice []DBPlayerLevelstats

func (entry *PlayerLevelstats) ToDB() *DBPlayerLevelstats {
	if entry == nil {
		return nil
	}
	return &DBPlayerLevelstats{
		Race: entry.Race,
		Class: entry.Class,
		Level: entry.Level,
		Str: entry.Str,
		Agi: entry.Agi,
		Sta: entry.Sta,
		Inte: entry.Inte,
		Spi: entry.Spi,
	}
}

func (entry *DBPlayerLevelstats) ToWeb() *PlayerLevelstats {
	if entry == nil {
		return nil
	}
	return &PlayerLevelstats{
		Race: entry.Race,
		Class: entry.Class,
		Level: entry.Level,
		Str: entry.Str,
		Agi: entry.Agi,
		Sta: entry.Sta,
		Inte: entry.Inte,
		Spi: entry.Spi,
	}
}

func (data PlayerLevelstatsSlice) ToDB() DBPlayerLevelstatsSlice {
	result := make(DBPlayerLevelstatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerLevelstatsSlice) ToWeb() PlayerLevelstatsSlice {
	result := make(PlayerLevelstatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
