package model

import "github.com/uptrace/bun"

type GameEventArenaSeasons struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Season int8 `json:"season,omitempty"`
}

type GameEventArenaSeasonsSlice []GameEventArenaSeasons

type DBGameEventArenaSeasons struct {
	bun.BaseModel `bun:"table:game_event_arena_seasons,alias:game_event_arena_seasons"`
	EventEntry int8 `bun:"eventEntry"`
	Season int8 `bun:"season"`
}

type DBGameEventArenaSeasonsSlice []DBGameEventArenaSeasons

func (entry *GameEventArenaSeasons) ToDB() *DBGameEventArenaSeasons {
	if entry == nil {
		return nil
	}
	return &DBGameEventArenaSeasons{
		EventEntry: entry.EventEntry,
		Season: entry.Season,
	}
}

func (entry *DBGameEventArenaSeasons) ToWeb() *GameEventArenaSeasons {
	if entry == nil {
		return nil
	}
	return &GameEventArenaSeasons{
		EventEntry: entry.EventEntry,
		Season: entry.Season,
	}
}

func (data GameEventArenaSeasonsSlice) ToDB() DBGameEventArenaSeasonsSlice {
	result := make(DBGameEventArenaSeasonsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventArenaSeasonsSlice) ToWeb() GameEventArenaSeasonsSlice {
	result := make(GameEventArenaSeasonsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
