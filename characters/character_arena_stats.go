package model

import "github.com/uptrace/bun"

type CharacterArenaStats struct {
	Guid int `json:"guid,omitempty"`
	Slot int8 `json:"slot,omitempty"`
	MatchMakerRating int16 `json:"matchmakerrating,omitempty"`
}

type CharacterArenaStatsSlice []CharacterArenaStats

type DBCharacterArenaStats struct {
	bun.BaseModel `bun:"table:character_arena_stats,alias:character_arena_stats"`
	Guid int `bun:"guid"`
	Slot int8 `bun:"slot"`
	MatchMakerRating int16 `bun:"matchMakerRating"`
}

type DBCharacterArenaStatsSlice []DBCharacterArenaStats

func (entry *CharacterArenaStats) ToDB() *DBCharacterArenaStats {
	if entry == nil {
		return nil
	}
	return &DBCharacterArenaStats{
		Guid: entry.Guid,
		Slot: entry.Slot,
		MatchMakerRating: entry.MatchMakerRating,
	}
}

func (entry *DBCharacterArenaStats) ToWeb() *CharacterArenaStats {
	if entry == nil {
		return nil
	}
	return &CharacterArenaStats{
		Guid: entry.Guid,
		Slot: entry.Slot,
		MatchMakerRating: entry.MatchMakerRating,
	}
}

func (data CharacterArenaStatsSlice) ToDB() DBCharacterArenaStatsSlice {
	result := make(DBCharacterArenaStatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterArenaStatsSlice) ToWeb() CharacterArenaStatsSlice {
	result := make(CharacterArenaStatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
