package model

import "github.com/uptrace/bun"

type CharacterQueststatusMonthly struct {
	Guid int `json:"guid,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type CharacterQueststatusMonthlySlice []CharacterQueststatusMonthly

type DBCharacterQueststatusMonthly struct {
	bun.BaseModel `bun:"table:character_queststatus_monthly,alias:character_queststatus_monthly"`
	Guid int `bun:"guid"`
	Quest int `bun:"quest"`
}

type DBCharacterQueststatusMonthlySlice []DBCharacterQueststatusMonthly

func (entry *CharacterQueststatusMonthly) ToDB() *DBCharacterQueststatusMonthly {
	if entry == nil {
		return nil
	}
	return &DBCharacterQueststatusMonthly{
		Guid: entry.Guid,
		Quest: entry.Quest,
	}
}

func (entry *DBCharacterQueststatusMonthly) ToWeb() *CharacterQueststatusMonthly {
	if entry == nil {
		return nil
	}
	return &CharacterQueststatusMonthly{
		Guid: entry.Guid,
		Quest: entry.Quest,
	}
}

func (data CharacterQueststatusMonthlySlice) ToDB() DBCharacterQueststatusMonthlySlice {
	result := make(DBCharacterQueststatusMonthlySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterQueststatusMonthlySlice) ToWeb() CharacterQueststatusMonthlySlice {
	result := make(CharacterQueststatusMonthlySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
