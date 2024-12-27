package model

import "github.com/uptrace/bun"

type CreatureQuestitem struct {
	CreatureEntry int `json:"creatureentry,omitempty"`
	Idx int `json:"idx,omitempty"`
	ItemId int `json:"itemid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type CreatureQuestitemSlice []CreatureQuestitem

type DBCreatureQuestitem struct {
	bun.BaseModel `bun:"table:creature_questitem,alias:creature_questitem"`
	CreatureEntry int `bun:"CreatureEntry"`
	Idx int `bun:"Idx"`
	ItemId int `bun:"ItemId"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBCreatureQuestitemSlice []DBCreatureQuestitem

func (entry *CreatureQuestitem) ToDB() *DBCreatureQuestitem {
	if entry == nil {
		return nil
	}
	return &DBCreatureQuestitem{
		CreatureEntry: entry.CreatureEntry,
		Idx: entry.Idx,
		ItemId: entry.ItemId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBCreatureQuestitem) ToWeb() *CreatureQuestitem {
	if entry == nil {
		return nil
	}
	return &CreatureQuestitem{
		CreatureEntry: entry.CreatureEntry,
		Idx: entry.Idx,
		ItemId: entry.ItemId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data CreatureQuestitemSlice) ToDB() DBCreatureQuestitemSlice {
	result := make(DBCreatureQuestitemSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureQuestitemSlice) ToWeb() CreatureQuestitemSlice {
	result := make(CreatureQuestitemSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
