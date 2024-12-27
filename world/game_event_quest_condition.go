package model

import "github.com/uptrace/bun"

type GameEventQuestCondition struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Quest int `json:"quest,omitempty"`
	ConditionId int `json:"condition_id,omitempty"`
	Num float64 `json:"num,omitempty"`
}

type GameEventQuestConditionSlice []GameEventQuestCondition

type DBGameEventQuestCondition struct {
	bun.BaseModel `bun:"table:game_event_quest_condition,alias:game_event_quest_condition"`
	EventEntry int8 `bun:"eventEntry"`
	Quest int `bun:"quest"`
	ConditionId int `bun:"condition_id"`
	Num float64 `bun:"num"`
}

type DBGameEventQuestConditionSlice []DBGameEventQuestCondition

func (entry *GameEventQuestCondition) ToDB() *DBGameEventQuestCondition {
	if entry == nil {
		return nil
	}
	return &DBGameEventQuestCondition{
		EventEntry: entry.EventEntry,
		Quest: entry.Quest,
		ConditionId: entry.ConditionId,
		Num: entry.Num,
	}
}

func (entry *DBGameEventQuestCondition) ToWeb() *GameEventQuestCondition {
	if entry == nil {
		return nil
	}
	return &GameEventQuestCondition{
		EventEntry: entry.EventEntry,
		Quest: entry.Quest,
		ConditionId: entry.ConditionId,
		Num: entry.Num,
	}
}

func (data GameEventQuestConditionSlice) ToDB() DBGameEventQuestConditionSlice {
	result := make(DBGameEventQuestConditionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventQuestConditionSlice) ToWeb() GameEventQuestConditionSlice {
	result := make(GameEventQuestConditionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
