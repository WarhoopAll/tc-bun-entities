package model

import "github.com/uptrace/bun"

type GameEventCreatureQuest struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Id int `json:"id,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type GameEventCreatureQuestSlice []GameEventCreatureQuest

type DBGameEventCreatureQuest struct {
	bun.BaseModel `bun:"table:game_event_creature_quest,alias:game_event_creature_quest"`
	EventEntry int8 `bun:"eventEntry"`
	Id int `bun:"id"`
	Quest int `bun:"quest"`
}

type DBGameEventCreatureQuestSlice []DBGameEventCreatureQuest

func (entry *GameEventCreatureQuest) ToDB() *DBGameEventCreatureQuest {
	if entry == nil {
		return nil
	}
	return &DBGameEventCreatureQuest{
		EventEntry: entry.EventEntry,
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (entry *DBGameEventCreatureQuest) ToWeb() *GameEventCreatureQuest {
	if entry == nil {
		return nil
	}
	return &GameEventCreatureQuest{
		EventEntry: entry.EventEntry,
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (data GameEventCreatureQuestSlice) ToDB() DBGameEventCreatureQuestSlice {
	result := make(DBGameEventCreatureQuestSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventCreatureQuestSlice) ToWeb() GameEventCreatureQuestSlice {
	result := make(GameEventCreatureQuestSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
