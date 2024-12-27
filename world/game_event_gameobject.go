package model

import "github.com/uptrace/bun"

type GameEventGameobject struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Guid int `json:"guid,omitempty"`
}

type GameEventGameobjectSlice []GameEventGameobject

type DBGameEventGameobject struct {
	bun.BaseModel `bun:"table:game_event_gameobject,alias:game_event_gameobject"`
	EventEntry int8 `bun:"eventEntry"`
	Guid int `bun:"guid"`
}

type DBGameEventGameobjectSlice []DBGameEventGameobject

func (entry *GameEventGameobject) ToDB() *DBGameEventGameobject {
	if entry == nil {
		return nil
	}
	return &DBGameEventGameobject{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
	}
}

func (entry *DBGameEventGameobject) ToWeb() *GameEventGameobject {
	if entry == nil {
		return nil
	}
	return &GameEventGameobject{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
	}
}

func (data GameEventGameobjectSlice) ToDB() DBGameEventGameobjectSlice {
	result := make(DBGameEventGameobjectSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventGameobjectSlice) ToWeb() GameEventGameobjectSlice {
	result := make(GameEventGameobjectSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
