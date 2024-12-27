package model

import "github.com/uptrace/bun"

type CreatureQuestender struct {
	Id int `json:"id,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type CreatureQuestenderSlice []CreatureQuestender

type DBCreatureQuestender struct {
	bun.BaseModel `bun:"table:creature_questender,alias:creature_questender"`
	Id int `bun:"id"`
	Quest int `bun:"quest"`
}

type DBCreatureQuestenderSlice []DBCreatureQuestender

func (entry *CreatureQuestender) ToDB() *DBCreatureQuestender {
	if entry == nil {
		return nil
	}
	return &DBCreatureQuestender{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (entry *DBCreatureQuestender) ToWeb() *CreatureQuestender {
	if entry == nil {
		return nil
	}
	return &CreatureQuestender{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (data CreatureQuestenderSlice) ToDB() DBCreatureQuestenderSlice {
	result := make(DBCreatureQuestenderSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureQuestenderSlice) ToWeb() CreatureQuestenderSlice {
	result := make(CreatureQuestenderSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
