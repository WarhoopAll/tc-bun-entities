package model

import "github.com/uptrace/bun"

type GameEventSeasonalQuestrelation struct {
	QuestId int `json:"questid,omitempty"`
	EventEntry int `json:"evententry,omitempty"`
}

type GameEventSeasonalQuestrelationSlice []GameEventSeasonalQuestrelation

type DBGameEventSeasonalQuestrelation struct {
	bun.BaseModel `bun:"table:game_event_seasonal_questrelation,alias:game_event_seasonal_questrelation"`
	QuestId int `bun:"questId"`
	EventEntry int `bun:"eventEntry"`
}

type DBGameEventSeasonalQuestrelationSlice []DBGameEventSeasonalQuestrelation

func (entry *GameEventSeasonalQuestrelation) ToDB() *DBGameEventSeasonalQuestrelation {
	if entry == nil {
		return nil
	}
	return &DBGameEventSeasonalQuestrelation{
		QuestId: entry.QuestId,
		EventEntry: entry.EventEntry,
	}
}

func (entry *DBGameEventSeasonalQuestrelation) ToWeb() *GameEventSeasonalQuestrelation {
	if entry == nil {
		return nil
	}
	return &GameEventSeasonalQuestrelation{
		QuestId: entry.QuestId,
		EventEntry: entry.EventEntry,
	}
}

func (data GameEventSeasonalQuestrelationSlice) ToDB() DBGameEventSeasonalQuestrelationSlice {
	result := make(DBGameEventSeasonalQuestrelationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventSeasonalQuestrelationSlice) ToWeb() GameEventSeasonalQuestrelationSlice {
	result := make(GameEventSeasonalQuestrelationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
