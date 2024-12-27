package model

import "github.com/uptrace/bun"

type BattlegroundDeserters struct {
	Guid int `json:"guid,omitempty"`
	Type int8 `json:"type,omitempty"`
	Datetime time.Time `json:"datetime,omitempty"`
}

type BattlegroundDesertersSlice []BattlegroundDeserters

type DBBattlegroundDeserters struct {
	bun.BaseModel `bun:"table:battleground_deserters,alias:battleground_deserters"`
	Guid int `bun:"guid"`
	Type int8 `bun:"type"`
	Datetime time.Time `bun:"datetime"`
}

type DBBattlegroundDesertersSlice []DBBattlegroundDeserters

func (entry *BattlegroundDeserters) ToDB() *DBBattlegroundDeserters {
	if entry == nil {
		return nil
	}
	return &DBBattlegroundDeserters{
		Guid: entry.Guid,
		Type: entry.Type,
		Datetime: entry.Datetime,
	}
}

func (entry *DBBattlegroundDeserters) ToWeb() *BattlegroundDeserters {
	if entry == nil {
		return nil
	}
	return &BattlegroundDeserters{
		Guid: entry.Guid,
		Type: entry.Type,
		Datetime: entry.Datetime,
	}
}

func (data BattlegroundDesertersSlice) ToDB() DBBattlegroundDesertersSlice {
	result := make(DBBattlegroundDesertersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBattlegroundDesertersSlice) ToWeb() BattlegroundDesertersSlice {
	result := make(BattlegroundDesertersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
