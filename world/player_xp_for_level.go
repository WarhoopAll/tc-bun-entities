package model

import "github.com/uptrace/bun"

type PlayerXpForLevel struct {
	Level int8 `json:"level,omitempty"`
	Experience int `json:"experience,omitempty"`
}

type PlayerXpForLevelSlice []PlayerXpForLevel

type DBPlayerXpForLevel struct {
	bun.BaseModel `bun:"table:player_xp_for_level,alias:player_xp_for_level"`
	Level int8 `bun:"Level"`
	Experience int `bun:"Experience"`
}

type DBPlayerXpForLevelSlice []DBPlayerXpForLevel

func (entry *PlayerXpForLevel) ToDB() *DBPlayerXpForLevel {
	if entry == nil {
		return nil
	}
	return &DBPlayerXpForLevel{
		Level: entry.Level,
		Experience: entry.Experience,
	}
}

func (entry *DBPlayerXpForLevel) ToWeb() *PlayerXpForLevel {
	if entry == nil {
		return nil
	}
	return &PlayerXpForLevel{
		Level: entry.Level,
		Experience: entry.Experience,
	}
}

func (data PlayerXpForLevelSlice) ToDB() DBPlayerXpForLevelSlice {
	result := make(DBPlayerXpForLevelSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerXpForLevelSlice) ToWeb() PlayerXpForLevelSlice {
	result := make(PlayerXpForLevelSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
