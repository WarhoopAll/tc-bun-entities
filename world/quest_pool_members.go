package model

import "github.com/uptrace/bun"

type QuestPoolMembers struct {
	QuestId int `json:"questid,omitempty"`
	PoolId int `json:"poolid,omitempty"`
	PoolIndex int8 `json:"poolindex,omitempty"`
	Description string `json:"description,omitempty"`
}

type QuestPoolMembersSlice []QuestPoolMembers

type DBQuestPoolMembers struct {
	bun.BaseModel `bun:"table:quest_pool_members,alias:quest_pool_members"`
	QuestId int `bun:"questId"`
	PoolId int `bun:"poolId"`
	PoolIndex int8 `bun:"poolIndex"`
	Description string `bun:"description"`
}

type DBQuestPoolMembersSlice []DBQuestPoolMembers

func (entry *QuestPoolMembers) ToDB() *DBQuestPoolMembers {
	if entry == nil {
		return nil
	}
	return &DBQuestPoolMembers{
		QuestId: entry.QuestId,
		PoolId: entry.PoolId,
		PoolIndex: entry.PoolIndex,
		Description: entry.Description,
	}
}

func (entry *DBQuestPoolMembers) ToWeb() *QuestPoolMembers {
	if entry == nil {
		return nil
	}
	return &QuestPoolMembers{
		QuestId: entry.QuestId,
		PoolId: entry.PoolId,
		PoolIndex: entry.PoolIndex,
		Description: entry.Description,
	}
}

func (data QuestPoolMembersSlice) ToDB() DBQuestPoolMembersSlice {
	result := make(DBQuestPoolMembersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestPoolMembersSlice) ToWeb() QuestPoolMembersSlice {
	result := make(QuestPoolMembersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
