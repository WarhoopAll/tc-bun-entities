package model

import "github.com/uptrace/bun"

type CharacterQueststatusDaily struct {
	Guid int `json:"guid,omitempty"`
	Quest int `json:"quest,omitempty"`
	Time int `json:"time,omitempty"`
}

type CharacterQueststatusDailySlice []CharacterQueststatusDaily

type DBCharacterQueststatusDaily struct {
	bun.BaseModel `bun:"table:character_queststatus_daily,alias:character_queststatus_daily"`
	Guid int `bun:"guid"`
	Quest int `bun:"quest"`
	Time int `bun:"time"`
}

type DBCharacterQueststatusDailySlice []DBCharacterQueststatusDaily

func (entry *CharacterQueststatusDaily) ToDB() *DBCharacterQueststatusDaily {
	if entry == nil {
		return nil
	}
	return &DBCharacterQueststatusDaily{
		Guid: entry.Guid,
		Quest: entry.Quest,
		Time: entry.Time,
	}
}

func (entry *DBCharacterQueststatusDaily) ToWeb() *CharacterQueststatusDaily {
	if entry == nil {
		return nil
	}
	return &CharacterQueststatusDaily{
		Guid: entry.Guid,
		Quest: entry.Quest,
		Time: entry.Time,
	}
}

func (data CharacterQueststatusDailySlice) ToDB() DBCharacterQueststatusDailySlice {
	result := make(DBCharacterQueststatusDailySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterQueststatusDailySlice) ToWeb() CharacterQueststatusDailySlice {
	result := make(CharacterQueststatusDailySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
