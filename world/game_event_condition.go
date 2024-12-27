package model

import "github.com/uptrace/bun"

type GameEventCondition struct {
	EventEntry int8 `json:"evententry,omitempty"`
	ConditionId int `json:"condition_id,omitempty"`
	ReqNum float64 `json:"req_num,omitempty"`
	MaxWorldStateField int16 `json:"max_world_state_field,omitempty"`
	DoneWorldStateField int16 `json:"done_world_state_field,omitempty"`
	Description string `json:"description,omitempty"`
}

type GameEventConditionSlice []GameEventCondition

type DBGameEventCondition struct {
	bun.BaseModel `bun:"table:game_event_condition,alias:game_event_condition"`
	EventEntry int8 `bun:"eventEntry"`
	ConditionId int `bun:"condition_id"`
	ReqNum float64 `bun:"req_num"`
	MaxWorldStateField int16 `bun:"max_world_state_field"`
	DoneWorldStateField int16 `bun:"done_world_state_field"`
	Description string `bun:"description"`
}

type DBGameEventConditionSlice []DBGameEventCondition

func (entry *GameEventCondition) ToDB() *DBGameEventCondition {
	if entry == nil {
		return nil
	}
	return &DBGameEventCondition{
		EventEntry: entry.EventEntry,
		ConditionId: entry.ConditionId,
		ReqNum: entry.ReqNum,
		MaxWorldStateField: entry.MaxWorldStateField,
		DoneWorldStateField: entry.DoneWorldStateField,
		Description: entry.Description,
	}
}

func (entry *DBGameEventCondition) ToWeb() *GameEventCondition {
	if entry == nil {
		return nil
	}
	return &GameEventCondition{
		EventEntry: entry.EventEntry,
		ConditionId: entry.ConditionId,
		ReqNum: entry.ReqNum,
		MaxWorldStateField: entry.MaxWorldStateField,
		DoneWorldStateField: entry.DoneWorldStateField,
		Description: entry.Description,
	}
}

func (data GameEventConditionSlice) ToDB() DBGameEventConditionSlice {
	result := make(DBGameEventConditionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventConditionSlice) ToWeb() GameEventConditionSlice {
	result := make(GameEventConditionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
