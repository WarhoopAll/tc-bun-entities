package model

import "github.com/uptrace/bun"

type GameEventPrerequisite struct {
	EventEntry int8 `json:"evententry,omitempty"`
	PrerequisiteEvent int `json:"prerequisite_event,omitempty"`
}

type GameEventPrerequisiteSlice []GameEventPrerequisite

type DBGameEventPrerequisite struct {
	bun.BaseModel `bun:"table:game_event_prerequisite,alias:game_event_prerequisite"`
	EventEntry int8 `bun:"eventEntry"`
	PrerequisiteEvent int `bun:"prerequisite_event"`
}

type DBGameEventPrerequisiteSlice []DBGameEventPrerequisite

func (entry *GameEventPrerequisite) ToDB() *DBGameEventPrerequisite {
	if entry == nil {
		return nil
	}
	return &DBGameEventPrerequisite{
		EventEntry: entry.EventEntry,
		PrerequisiteEvent: entry.PrerequisiteEvent,
	}
}

func (entry *DBGameEventPrerequisite) ToWeb() *GameEventPrerequisite {
	if entry == nil {
		return nil
	}
	return &GameEventPrerequisite{
		EventEntry: entry.EventEntry,
		PrerequisiteEvent: entry.PrerequisiteEvent,
	}
}

func (data GameEventPrerequisiteSlice) ToDB() DBGameEventPrerequisiteSlice {
	result := make(DBGameEventPrerequisiteSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventPrerequisiteSlice) ToWeb() GameEventPrerequisiteSlice {
	result := make(GameEventPrerequisiteSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
