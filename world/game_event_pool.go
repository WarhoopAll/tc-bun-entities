package model

import "github.com/uptrace/bun"

type GameEventPool struct {
	EventEntry int8 `json:"evententry,omitempty"`
	PoolEntry int `json:"pool_entry,omitempty"`
}

type GameEventPoolSlice []GameEventPool

type DBGameEventPool struct {
	bun.BaseModel `bun:"table:game_event_pool,alias:game_event_pool"`
	EventEntry int8 `bun:"eventEntry"`
	PoolEntry int `bun:"pool_entry"`
}

type DBGameEventPoolSlice []DBGameEventPool

func (entry *GameEventPool) ToDB() *DBGameEventPool {
	if entry == nil {
		return nil
	}
	return &DBGameEventPool{
		EventEntry: entry.EventEntry,
		PoolEntry: entry.PoolEntry,
	}
}

func (entry *DBGameEventPool) ToWeb() *GameEventPool {
	if entry == nil {
		return nil
	}
	return &GameEventPool{
		EventEntry: entry.EventEntry,
		PoolEntry: entry.PoolEntry,
	}
}

func (data GameEventPoolSlice) ToDB() DBGameEventPoolSlice {
	result := make(DBGameEventPoolSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventPoolSlice) ToWeb() GameEventPoolSlice {
	result := make(GameEventPoolSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
