package model

import "github.com/uptrace/bun"

type GameobjectQuestender struct {
	Id int `json:"id,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type GameobjectQuestenderSlice []GameobjectQuestender

type DBGameobjectQuestender struct {
	bun.BaseModel `bun:"table:gameobject_questender,alias:gameobject_questender"`
	Id int `bun:"id"`
	Quest int `bun:"quest"`
}

type DBGameobjectQuestenderSlice []DBGameobjectQuestender

func (entry *GameobjectQuestender) ToDB() *DBGameobjectQuestender {
	if entry == nil {
		return nil
	}
	return &DBGameobjectQuestender{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (entry *DBGameobjectQuestender) ToWeb() *GameobjectQuestender {
	if entry == nil {
		return nil
	}
	return &GameobjectQuestender{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (data GameobjectQuestenderSlice) ToDB() DBGameobjectQuestenderSlice {
	result := make(DBGameobjectQuestenderSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectQuestenderSlice) ToWeb() GameobjectQuestenderSlice {
	result := make(GameobjectQuestenderSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
