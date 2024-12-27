package model

import "github.com/uptrace/bun"

type CreatureQueststarter struct {
	Id int `json:"id,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type CreatureQueststarterSlice []CreatureQueststarter

type DBCreatureQueststarter struct {
	bun.BaseModel `bun:"table:creature_queststarter,alias:creature_queststarter"`
	Id int `bun:"id"`
	Quest int `bun:"quest"`
}

type DBCreatureQueststarterSlice []DBCreatureQueststarter

func (entry *CreatureQueststarter) ToDB() *DBCreatureQueststarter {
	if entry == nil {
		return nil
	}
	return &DBCreatureQueststarter{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (entry *DBCreatureQueststarter) ToWeb() *CreatureQueststarter {
	if entry == nil {
		return nil
	}
	return &CreatureQueststarter{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (data CreatureQueststarterSlice) ToDB() DBCreatureQueststarterSlice {
	result := make(DBCreatureQueststarterSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureQueststarterSlice) ToWeb() CreatureQueststarterSlice {
	result := make(CreatureQueststarterSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
