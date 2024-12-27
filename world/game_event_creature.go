package model

import "github.com/uptrace/bun"

type GameEventCreature struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Guid int `json:"guid,omitempty"`
}

type GameEventCreatureSlice []GameEventCreature

type DBGameEventCreature struct {
	bun.BaseModel `bun:"table:game_event_creature,alias:game_event_creature"`
	EventEntry int8 `bun:"eventEntry"`
	Guid int `bun:"guid"`
}

type DBGameEventCreatureSlice []DBGameEventCreature

func (entry *GameEventCreature) ToDB() *DBGameEventCreature {
	if entry == nil {
		return nil
	}
	return &DBGameEventCreature{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
	}
}

func (entry *DBGameEventCreature) ToWeb() *GameEventCreature {
	if entry == nil {
		return nil
	}
	return &GameEventCreature{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
	}
}

func (data GameEventCreatureSlice) ToDB() DBGameEventCreatureSlice {
	result := make(DBGameEventCreatureSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventCreatureSlice) ToWeb() GameEventCreatureSlice {
	result := make(GameEventCreatureSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
