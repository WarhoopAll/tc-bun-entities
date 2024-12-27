package model

import "github.com/uptrace/bun"

type LfgData struct {
	Guid int `json:"guid,omitempty"`
	Dungeon int `json:"dungeon,omitempty"`
	State int8 `json:"state,omitempty"`
}

type LfgDataSlice []LfgData

type DBLfgData struct {
	bun.BaseModel `bun:"table:lfg_data,alias:lfg_data"`
	Guid int `bun:"guid"`
	Dungeon int `bun:"dungeon"`
	State int8 `bun:"state"`
}

type DBLfgDataSlice []DBLfgData

func (entry *LfgData) ToDB() *DBLfgData {
	if entry == nil {
		return nil
	}
	return &DBLfgData{
		Guid: entry.Guid,
		Dungeon: entry.Dungeon,
		State: entry.State,
	}
}

func (entry *DBLfgData) ToWeb() *LfgData {
	if entry == nil {
		return nil
	}
	return &LfgData{
		Guid: entry.Guid,
		Dungeon: entry.Dungeon,
		State: entry.State,
	}
}

func (data LfgDataSlice) ToDB() DBLfgDataSlice {
	result := make(DBLfgDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBLfgDataSlice) ToWeb() LfgDataSlice {
	result := make(LfgDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
