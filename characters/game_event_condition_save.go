package model

import "github.com/uptrace/bun"

type GameEventConditionSave struct {
	EventEntry int8 `json:"evententry,omitempty"`
	ConditionId int `json:"condition_id,omitempty"`
	Done float64 `json:"done,omitempty"`
}

type GameEventConditionSaveSlice []GameEventConditionSave

type DBGameEventConditionSave struct {
	bun.BaseModel `bun:"table:game_event_condition_save,alias:game_event_condition_save"`
	EventEntry int8 `bun:"eventEntry"`
	ConditionId int `bun:"condition_id"`
	Done float64 `bun:"done"`
}

type DBGameEventConditionSaveSlice []DBGameEventConditionSave

func (entry *GameEventConditionSave) ToDB() *DBGameEventConditionSave {
	if entry == nil {
		return nil
	}
	return &DBGameEventConditionSave{
		EventEntry: entry.EventEntry,
		ConditionId: entry.ConditionId,
		Done: entry.Done,
	}
}

func (entry *DBGameEventConditionSave) ToWeb() *GameEventConditionSave {
	if entry == nil {
		return nil
	}
	return &GameEventConditionSave{
		EventEntry: entry.EventEntry,
		ConditionId: entry.ConditionId,
		Done: entry.Done,
	}
}

func (data GameEventConditionSaveSlice) ToDB() DBGameEventConditionSaveSlice {
	result := make(DBGameEventConditionSaveSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventConditionSaveSlice) ToWeb() GameEventConditionSaveSlice {
	result := make(GameEventConditionSaveSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
