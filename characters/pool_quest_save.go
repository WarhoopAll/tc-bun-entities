package model

import "github.com/uptrace/bun"

type PoolQuestSave struct {
	PoolId int `json:"pool_id,omitempty"`
	QuestId int `json:"quest_id,omitempty"`
}

type PoolQuestSaveSlice []PoolQuestSave

type DBPoolQuestSave struct {
	bun.BaseModel `bun:"table:pool_quest_save,alias:pool_quest_save"`
	PoolId int `bun:"pool_id"`
	QuestId int `bun:"quest_id"`
}

type DBPoolQuestSaveSlice []DBPoolQuestSave

func (entry *PoolQuestSave) ToDB() *DBPoolQuestSave {
	if entry == nil {
		return nil
	}
	return &DBPoolQuestSave{
		PoolId: entry.PoolId,
		QuestId: entry.QuestId,
	}
}

func (entry *DBPoolQuestSave) ToWeb() *PoolQuestSave {
	if entry == nil {
		return nil
	}
	return &PoolQuestSave{
		PoolId: entry.PoolId,
		QuestId: entry.QuestId,
	}
}

func (data PoolQuestSaveSlice) ToDB() DBPoolQuestSaveSlice {
	result := make(DBPoolQuestSaveSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPoolQuestSaveSlice) ToWeb() PoolQuestSaveSlice {
	result := make(PoolQuestSaveSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
