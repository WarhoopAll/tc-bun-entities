package model

import "github.com/uptrace/bun"

type GameEventBattlegroundHoliday struct {
	EventEntry int8 `json:"evententry,omitempty"`
	BattlegroundID int `json:"battlegroundid,omitempty"`
}

type GameEventBattlegroundHolidaySlice []GameEventBattlegroundHoliday

type DBGameEventBattlegroundHoliday struct {
	bun.BaseModel `bun:"table:game_event_battleground_holiday,alias:game_event_battleground_holiday"`
	EventEntry int8 `bun:"EventEntry"`
	BattlegroundID int `bun:"BattlegroundID"`
}

type DBGameEventBattlegroundHolidaySlice []DBGameEventBattlegroundHoliday

func (entry *GameEventBattlegroundHoliday) ToDB() *DBGameEventBattlegroundHoliday {
	if entry == nil {
		return nil
	}
	return &DBGameEventBattlegroundHoliday{
		EventEntry: entry.EventEntry,
		BattlegroundID: entry.BattlegroundID,
	}
}

func (entry *DBGameEventBattlegroundHoliday) ToWeb() *GameEventBattlegroundHoliday {
	if entry == nil {
		return nil
	}
	return &GameEventBattlegroundHoliday{
		EventEntry: entry.EventEntry,
		BattlegroundID: entry.BattlegroundID,
	}
}

func (data GameEventBattlegroundHolidaySlice) ToDB() DBGameEventBattlegroundHolidaySlice {
	result := make(DBGameEventBattlegroundHolidaySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventBattlegroundHolidaySlice) ToWeb() GameEventBattlegroundHolidaySlice {
	result := make(GameEventBattlegroundHolidaySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
