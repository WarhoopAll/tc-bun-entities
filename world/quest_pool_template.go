package model

import "github.com/uptrace/bun"

type QuestPoolTemplate struct {
	PoolId int `json:"poolid,omitempty"`
	NumActive int `json:"numactive,omitempty"`
	Description string `json:"description,omitempty"`
}

type QuestPoolTemplateSlice []QuestPoolTemplate

type DBQuestPoolTemplate struct {
	bun.BaseModel `bun:"table:quest_pool_template,alias:quest_pool_template"`
	PoolId int `bun:"poolId"`
	NumActive int `bun:"numActive"`
	Description string `bun:"description"`
}

type DBQuestPoolTemplateSlice []DBQuestPoolTemplate

func (entry *QuestPoolTemplate) ToDB() *DBQuestPoolTemplate {
	if entry == nil {
		return nil
	}
	return &DBQuestPoolTemplate{
		PoolId: entry.PoolId,
		NumActive: entry.NumActive,
		Description: entry.Description,
	}
}

func (entry *DBQuestPoolTemplate) ToWeb() *QuestPoolTemplate {
	if entry == nil {
		return nil
	}
	return &QuestPoolTemplate{
		PoolId: entry.PoolId,
		NumActive: entry.NumActive,
		Description: entry.Description,
	}
}

func (data QuestPoolTemplateSlice) ToDB() DBQuestPoolTemplateSlice {
	result := make(DBQuestPoolTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestPoolTemplateSlice) ToWeb() QuestPoolTemplateSlice {
	result := make(QuestPoolTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
