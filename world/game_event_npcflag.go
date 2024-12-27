package model

import "github.com/uptrace/bun"

type GameEventNpcflag struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Guid int `json:"guid,omitempty"`
	Npcflag int `json:"npcflag,omitempty"`
}

type GameEventNpcflagSlice []GameEventNpcflag

type DBGameEventNpcflag struct {
	bun.BaseModel `bun:"table:game_event_npcflag,alias:game_event_npcflag"`
	EventEntry int8 `bun:"eventEntry"`
	Guid int `bun:"guid"`
	Npcflag int `bun:"npcflag"`
}

type DBGameEventNpcflagSlice []DBGameEventNpcflag

func (entry *GameEventNpcflag) ToDB() *DBGameEventNpcflag {
	if entry == nil {
		return nil
	}
	return &DBGameEventNpcflag{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
		Npcflag: entry.Npcflag,
	}
}

func (entry *DBGameEventNpcflag) ToWeb() *GameEventNpcflag {
	if entry == nil {
		return nil
	}
	return &GameEventNpcflag{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
		Npcflag: entry.Npcflag,
	}
}

func (data GameEventNpcflagSlice) ToDB() DBGameEventNpcflagSlice {
	result := make(DBGameEventNpcflagSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventNpcflagSlice) ToWeb() GameEventNpcflagSlice {
	result := make(GameEventNpcflagSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
