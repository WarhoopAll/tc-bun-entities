package model

import "github.com/uptrace/bun"

type PvpstatsBattlegrounds struct {
	Id int `json:"id,omitempty"`
	WinnerFaction int8 `json:"winner_faction,omitempty"`
	BracketId int8 `json:"bracket_id,omitempty"`
	Type int8 `json:"type,omitempty"`
	Date time.Time `json:"date,omitempty"`
}

type PvpstatsBattlegroundsSlice []PvpstatsBattlegrounds

type DBPvpstatsBattlegrounds struct {
	bun.BaseModel `bun:"table:pvpstats_battlegrounds,alias:pvpstats_battlegrounds"`
	Id int `bun:"id"`
	WinnerFaction int8 `bun:"winner_faction"`
	BracketId int8 `bun:"bracket_id"`
	Type int8 `bun:"type"`
	Date time.Time `bun:"date"`
}

type DBPvpstatsBattlegroundsSlice []DBPvpstatsBattlegrounds

func (entry *PvpstatsBattlegrounds) ToDB() *DBPvpstatsBattlegrounds {
	if entry == nil {
		return nil
	}
	return &DBPvpstatsBattlegrounds{
		Id: entry.Id,
		WinnerFaction: entry.WinnerFaction,
		BracketId: entry.BracketId,
		Type: entry.Type,
		Date: entry.Date,
	}
}

func (entry *DBPvpstatsBattlegrounds) ToWeb() *PvpstatsBattlegrounds {
	if entry == nil {
		return nil
	}
	return &PvpstatsBattlegrounds{
		Id: entry.Id,
		WinnerFaction: entry.WinnerFaction,
		BracketId: entry.BracketId,
		Type: entry.Type,
		Date: entry.Date,
	}
}

func (data PvpstatsBattlegroundsSlice) ToDB() DBPvpstatsBattlegroundsSlice {
	result := make(DBPvpstatsBattlegroundsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPvpstatsBattlegroundsSlice) ToWeb() PvpstatsBattlegroundsSlice {
	result := make(PvpstatsBattlegroundsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
