package model

import "github.com/uptrace/bun"

type GameEventGameobjectQuest struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Id int `json:"id,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type GameEventGameobjectQuestSlice []GameEventGameobjectQuest

type DBGameEventGameobjectQuest struct {
	bun.BaseModel `bun:"table:game_event_gameobject_quest,alias:game_event_gameobject_quest"`
	EventEntry int8 `bun:"eventEntry"`
	Id int `bun:"id"`
	Quest int `bun:"quest"`
}

type DBGameEventGameobjectQuestSlice []DBGameEventGameobjectQuest

func (entry *GameEventGameobjectQuest) ToDB() *DBGameEventGameobjectQuest {
	if entry == nil {
		return nil
	}
	return &DBGameEventGameobjectQuest{
		EventEntry: entry.EventEntry,
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (entry *DBGameEventGameobjectQuest) ToWeb() *GameEventGameobjectQuest {
	if entry == nil {
		return nil
	}
	return &GameEventGameobjectQuest{
		EventEntry: entry.EventEntry,
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (data GameEventGameobjectQuestSlice) ToDB() DBGameEventGameobjectQuestSlice {
	result := make(DBGameEventGameobjectQuestSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventGameobjectQuestSlice) ToWeb() GameEventGameobjectQuestSlice {
	result := make(GameEventGameobjectQuestSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
