package model

import "github.com/uptrace/bun"

type LfgDungeonRewards struct {
	DungeonId int `json:"dungeonid,omitempty"`
	MaxLevel int8 `json:"maxlevel,omitempty"`
	FirstQuestId int `json:"firstquestid,omitempty"`
	OtherQuestId int `json:"otherquestid,omitempty"`
}

type LfgDungeonRewardsSlice []LfgDungeonRewards

type DBLfgDungeonRewards struct {
	bun.BaseModel `bun:"table:lfg_dungeon_rewards,alias:lfg_dungeon_rewards"`
	DungeonId int `bun:"dungeonId"`
	MaxLevel int8 `bun:"maxLevel"`
	FirstQuestId int `bun:"firstQuestId"`
	OtherQuestId int `bun:"otherQuestId"`
}

type DBLfgDungeonRewardsSlice []DBLfgDungeonRewards

func (entry *LfgDungeonRewards) ToDB() *DBLfgDungeonRewards {
	if entry == nil {
		return nil
	}
	return &DBLfgDungeonRewards{
		DungeonId: entry.DungeonId,
		MaxLevel: entry.MaxLevel,
		FirstQuestId: entry.FirstQuestId,
		OtherQuestId: entry.OtherQuestId,
	}
}

func (entry *DBLfgDungeonRewards) ToWeb() *LfgDungeonRewards {
	if entry == nil {
		return nil
	}
	return &LfgDungeonRewards{
		DungeonId: entry.DungeonId,
		MaxLevel: entry.MaxLevel,
		FirstQuestId: entry.FirstQuestId,
		OtherQuestId: entry.OtherQuestId,
	}
}

func (data LfgDungeonRewardsSlice) ToDB() DBLfgDungeonRewardsSlice {
	result := make(DBLfgDungeonRewardsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBLfgDungeonRewardsSlice) ToWeb() LfgDungeonRewardsSlice {
	result := make(LfgDungeonRewardsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
