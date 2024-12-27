package model

import "github.com/uptrace/bun"

type GameobjectQuestitem struct {
	GameObjectEntry int `json:"gameobjectentry,omitempty"`
	Idx int `json:"idx,omitempty"`
	ItemId int `json:"itemid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type GameobjectQuestitemSlice []GameobjectQuestitem

type DBGameobjectQuestitem struct {
	bun.BaseModel `bun:"table:gameobject_questitem,alias:gameobject_questitem"`
	GameObjectEntry int `bun:"GameObjectEntry"`
	Idx int `bun:"Idx"`
	ItemId int `bun:"ItemId"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBGameobjectQuestitemSlice []DBGameobjectQuestitem

func (entry *GameobjectQuestitem) ToDB() *DBGameobjectQuestitem {
	if entry == nil {
		return nil
	}
	return &DBGameobjectQuestitem{
		GameObjectEntry: entry.GameObjectEntry,
		Idx: entry.Idx,
		ItemId: entry.ItemId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBGameobjectQuestitem) ToWeb() *GameobjectQuestitem {
	if entry == nil {
		return nil
	}
	return &GameobjectQuestitem{
		GameObjectEntry: entry.GameObjectEntry,
		Idx: entry.Idx,
		ItemId: entry.ItemId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data GameobjectQuestitemSlice) ToDB() DBGameobjectQuestitemSlice {
	result := make(DBGameobjectQuestitemSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectQuestitemSlice) ToWeb() GameobjectQuestitemSlice {
	result := make(GameobjectQuestitemSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
