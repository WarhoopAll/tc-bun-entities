package model

import "github.com/uptrace/bun"

type CharacterQueststatusSeasonal struct {
	Guid int `json:"guid,omitempty"`
	Quest int `json:"quest,omitempty"`
	Event int `json:"event,omitempty"`
}

type CharacterQueststatusSeasonalSlice []CharacterQueststatusSeasonal

type DBCharacterQueststatusSeasonal struct {
	bun.BaseModel `bun:"table:character_queststatus_seasonal,alias:character_queststatus_seasonal"`
	Guid int `bun:"guid"`
	Quest int `bun:"quest"`
	Event int `bun:"event"`
}

type DBCharacterQueststatusSeasonalSlice []DBCharacterQueststatusSeasonal

func (entry *CharacterQueststatusSeasonal) ToDB() *DBCharacterQueststatusSeasonal {
	if entry == nil {
		return nil
	}
	return &DBCharacterQueststatusSeasonal{
		Guid: entry.Guid,
		Quest: entry.Quest,
		Event: entry.Event,
	}
}

func (entry *DBCharacterQueststatusSeasonal) ToWeb() *CharacterQueststatusSeasonal {
	if entry == nil {
		return nil
	}
	return &CharacterQueststatusSeasonal{
		Guid: entry.Guid,
		Quest: entry.Quest,
		Event: entry.Event,
	}
}

func (data CharacterQueststatusSeasonalSlice) ToDB() DBCharacterQueststatusSeasonalSlice {
	result := make(DBCharacterQueststatusSeasonalSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterQueststatusSeasonalSlice) ToWeb() CharacterQueststatusSeasonalSlice {
	result := make(CharacterQueststatusSeasonalSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
