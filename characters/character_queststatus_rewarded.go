package model

import "github.com/uptrace/bun"

type CharacterQueststatusRewarded struct {
	Guid int `json:"guid,omitempty"`
	Quest int `json:"quest,omitempty"`
	Active int8 `json:"active,omitempty"`
}

type CharacterQueststatusRewardedSlice []CharacterQueststatusRewarded

type DBCharacterQueststatusRewarded struct {
	bun.BaseModel `bun:"table:character_queststatus_rewarded,alias:character_queststatus_rewarded"`
	Guid int `bun:"guid"`
	Quest int `bun:"quest"`
	Active int8 `bun:"active"`
}

type DBCharacterQueststatusRewardedSlice []DBCharacterQueststatusRewarded

func (entry *CharacterQueststatusRewarded) ToDB() *DBCharacterQueststatusRewarded {
	if entry == nil {
		return nil
	}
	return &DBCharacterQueststatusRewarded{
		Guid: entry.Guid,
		Quest: entry.Quest,
		Active: entry.Active,
	}
}

func (entry *DBCharacterQueststatusRewarded) ToWeb() *CharacterQueststatusRewarded {
	if entry == nil {
		return nil
	}
	return &CharacterQueststatusRewarded{
		Guid: entry.Guid,
		Quest: entry.Quest,
		Active: entry.Active,
	}
}

func (data CharacterQueststatusRewardedSlice) ToDB() DBCharacterQueststatusRewardedSlice {
	result := make(DBCharacterQueststatusRewardedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterQueststatusRewardedSlice) ToWeb() CharacterQueststatusRewardedSlice {
	result := make(CharacterQueststatusRewardedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
