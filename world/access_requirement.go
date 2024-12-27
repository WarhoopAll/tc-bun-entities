package model

import "github.com/uptrace/bun"

type AccessRequirement struct {
	MapId int `json:"mapid,omitempty"`
	Difficulty int8 `json:"difficulty,omitempty"`
	LevelMin int8 `json:"level_min,omitempty"`
	LevelMax int8 `json:"level_max,omitempty"`
	ItemLevel int16 `json:"item_level,omitempty"`
	Item int `json:"item,omitempty"`
	Item2 int `json:"item2,omitempty"`
	QuestDoneA int `json:"quest_done_a,omitempty"`
	QuestDoneH int `json:"quest_done_h,omitempty"`
	CompletedAchievement int `json:"completed_achievement,omitempty"`
	QuestFailedText string `json:"quest_failed_text,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type AccessRequirementSlice []AccessRequirement

type DBAccessRequirement struct {
	bun.BaseModel `bun:"table:access_requirement,alias:access_requirement"`
	MapId int `bun:"mapId"`
	Difficulty int8 `bun:"difficulty"`
	LevelMin int8 `bun:"level_min"`
	LevelMax int8 `bun:"level_max"`
	ItemLevel int16 `bun:"item_level"`
	Item int `bun:"item"`
	Item2 int `bun:"item2"`
	QuestDoneA int `bun:"quest_done_A"`
	QuestDoneH int `bun:"quest_done_H"`
	CompletedAchievement int `bun:"completed_achievement"`
	QuestFailedText string `bun:"quest_failed_text"`
	Comment string `bun:"comment"`
}

type DBAccessRequirementSlice []DBAccessRequirement

func (entry *AccessRequirement) ToDB() *DBAccessRequirement {
	if entry == nil {
		return nil
	}
	return &DBAccessRequirement{
		MapId: entry.MapId,
		Difficulty: entry.Difficulty,
		LevelMin: entry.LevelMin,
		LevelMax: entry.LevelMax,
		ItemLevel: entry.ItemLevel,
		Item: entry.Item,
		Item2: entry.Item2,
		QuestDoneA: entry.QuestDoneA,
		QuestDoneH: entry.QuestDoneH,
		CompletedAchievement: entry.CompletedAchievement,
		QuestFailedText: entry.QuestFailedText,
		Comment: entry.Comment,
	}
}

func (entry *DBAccessRequirement) ToWeb() *AccessRequirement {
	if entry == nil {
		return nil
	}
	return &AccessRequirement{
		MapId: entry.MapId,
		Difficulty: entry.Difficulty,
		LevelMin: entry.LevelMin,
		LevelMax: entry.LevelMax,
		ItemLevel: entry.ItemLevel,
		Item: entry.Item,
		Item2: entry.Item2,
		QuestDoneA: entry.QuestDoneA,
		QuestDoneH: entry.QuestDoneH,
		CompletedAchievement: entry.CompletedAchievement,
		QuestFailedText: entry.QuestFailedText,
		Comment: entry.Comment,
	}
}

func (data AccessRequirementSlice) ToDB() DBAccessRequirementSlice {
	result := make(DBAccessRequirementSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAccessRequirementSlice) ToWeb() AccessRequirementSlice {
	result := make(AccessRequirementSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
