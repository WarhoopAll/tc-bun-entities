package model

import "github.com/uptrace/bun"

type GameEvent struct {
	EventEntry int8 `json:"evententry,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime time.Time `json:"end_time,omitempty"`
	Occurence int `json:"occurence,omitempty"`
	Length int `json:"length,omitempty"`
	Holiday int `json:"holiday,omitempty"`
	HolidayStage int8 `json:"holidaystage,omitempty"`
	Description string `json:"description,omitempty"`
	WorldEvent int8 `json:"world_event,omitempty"`
	Announce int8 `json:"announce,omitempty"`
}

type GameEventSlice []GameEvent

type DBGameEvent struct {
	bun.BaseModel `bun:"table:game_event,alias:game_event"`
	EventEntry int8 `bun:"eventEntry"`
	StartTime time.Time `bun:"start_time"`
	EndTime time.Time `bun:"end_time"`
	Occurence int `bun:"occurence"`
	Length int `bun:"length"`
	Holiday int `bun:"holiday"`
	HolidayStage int8 `bun:"holidayStage"`
	Description string `bun:"description"`
	WorldEvent int8 `bun:"world_event"`
	Announce int8 `bun:"announce"`
}

type DBGameEventSlice []DBGameEvent

func (entry *GameEvent) ToDB() *DBGameEvent {
	if entry == nil {
		return nil
	}
	return &DBGameEvent{
		EventEntry: entry.EventEntry,
		StartTime: entry.StartTime,
		EndTime: entry.EndTime,
		Occurence: entry.Occurence,
		Length: entry.Length,
		Holiday: entry.Holiday,
		HolidayStage: entry.HolidayStage,
		Description: entry.Description,
		WorldEvent: entry.WorldEvent,
		Announce: entry.Announce,
	}
}

func (entry *DBGameEvent) ToWeb() *GameEvent {
	if entry == nil {
		return nil
	}
	return &GameEvent{
		EventEntry: entry.EventEntry,
		StartTime: entry.StartTime,
		EndTime: entry.EndTime,
		Occurence: entry.Occurence,
		Length: entry.Length,
		Holiday: entry.Holiday,
		HolidayStage: entry.HolidayStage,
		Description: entry.Description,
		WorldEvent: entry.WorldEvent,
		Announce: entry.Announce,
	}
}

func (data GameEventSlice) ToDB() DBGameEventSlice {
	result := make(DBGameEventSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventSlice) ToWeb() GameEventSlice {
	result := make(GameEventSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
