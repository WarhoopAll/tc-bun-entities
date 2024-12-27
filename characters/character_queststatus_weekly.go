package model

import "github.com/uptrace/bun"

type CharacterQueststatusWeekly struct {
	Guid int `json:"guid,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type CharacterQueststatusWeeklySlice []CharacterQueststatusWeekly

type DBCharacterQueststatusWeekly struct {
	bun.BaseModel `bun:"table:character_queststatus_weekly,alias:character_queststatus_weekly"`
	Guid int `bun:"guid"`
	Quest int `bun:"quest"`
}

type DBCharacterQueststatusWeeklySlice []DBCharacterQueststatusWeekly

func (entry *CharacterQueststatusWeekly) ToDB() *DBCharacterQueststatusWeekly {
	if entry == nil {
		return nil
	}
	return &DBCharacterQueststatusWeekly{
		Guid: entry.Guid,
		Quest: entry.Quest,
	}
}

func (entry *DBCharacterQueststatusWeekly) ToWeb() *CharacterQueststatusWeekly {
	if entry == nil {
		return nil
	}
	return &CharacterQueststatusWeekly{
		Guid: entry.Guid,
		Quest: entry.Quest,
	}
}

func (data CharacterQueststatusWeeklySlice) ToDB() DBCharacterQueststatusWeeklySlice {
	result := make(DBCharacterQueststatusWeeklySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterQueststatusWeeklySlice) ToWeb() CharacterQueststatusWeeklySlice {
	result := make(CharacterQueststatusWeeklySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
