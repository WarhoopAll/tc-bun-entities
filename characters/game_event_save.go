package model

import "github.com/uptrace/bun"

type GameEventSave struct {
	EventEntry int8 `json:"evententry,omitempty"`
	State int8 `json:"state,omitempty"`
	NextStart int `json:"next_start,omitempty"`
}

type GameEventSaveSlice []GameEventSave

type DBGameEventSave struct {
	bun.BaseModel `bun:"table:game_event_save,alias:game_event_save"`
	EventEntry int8 `bun:"eventEntry"`
	State int8 `bun:"state"`
	NextStart int `bun:"next_start"`
}

type DBGameEventSaveSlice []DBGameEventSave

func (entry *GameEventSave) ToDB() *DBGameEventSave {
	if entry == nil {
		return nil
	}
	return &DBGameEventSave{
		EventEntry: entry.EventEntry,
		State: entry.State,
		NextStart: entry.NextStart,
	}
}

func (entry *DBGameEventSave) ToWeb() *GameEventSave {
	if entry == nil {
		return nil
	}
	return &GameEventSave{
		EventEntry: entry.EventEntry,
		State: entry.State,
		NextStart: entry.NextStart,
	}
}

func (data GameEventSaveSlice) ToDB() DBGameEventSaveSlice {
	result := make(DBGameEventSaveSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventSaveSlice) ToWeb() GameEventSaveSlice {
	result := make(GameEventSaveSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
