package model

import "github.com/uptrace/bun"

type GameobjectQueststarter struct {
	Id int `json:"id,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type GameobjectQueststarterSlice []GameobjectQueststarter

type DBGameobjectQueststarter struct {
	bun.BaseModel `bun:"table:gameobject_queststarter,alias:gameobject_queststarter"`
	Id int `bun:"id"`
	Quest int `bun:"quest"`
}

type DBGameobjectQueststarterSlice []DBGameobjectQueststarter

func (entry *GameobjectQueststarter) ToDB() *DBGameobjectQueststarter {
	if entry == nil {
		return nil
	}
	return &DBGameobjectQueststarter{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (entry *DBGameobjectQueststarter) ToWeb() *GameobjectQueststarter {
	if entry == nil {
		return nil
	}
	return &GameobjectQueststarter{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (data GameobjectQueststarterSlice) ToDB() DBGameobjectQueststarterSlice {
	result := make(DBGameobjectQueststarterSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectQueststarterSlice) ToWeb() GameobjectQueststarterSlice {
	result := make(GameobjectQueststarterSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
