package model

import "github.com/uptrace/bun"

type SpelldifficultyDbc struct {
	Id int `json:"id,omitempty"`
	Spellid0 int `json:"spellid0,omitempty"`
	Spellid1 int `json:"spellid1,omitempty"`
	Spellid2 int `json:"spellid2,omitempty"`
	Spellid3 int `json:"spellid3,omitempty"`
}

type SpelldifficultyDbcSlice []SpelldifficultyDbc

type DBSpelldifficultyDbc struct {
	bun.BaseModel `bun:"table:spelldifficulty_dbc,alias:spelldifficulty_dbc"`
	Id int `bun:"id"`
	Spellid0 int `bun:"spellid0"`
	Spellid1 int `bun:"spellid1"`
	Spellid2 int `bun:"spellid2"`
	Spellid3 int `bun:"spellid3"`
}

type DBSpelldifficultyDbcSlice []DBSpelldifficultyDbc

func (entry *SpelldifficultyDbc) ToDB() *DBSpelldifficultyDbc {
	if entry == nil {
		return nil
	}
	return &DBSpelldifficultyDbc{
		Id: entry.Id,
		Spellid0: entry.Spellid0,
		Spellid1: entry.Spellid1,
		Spellid2: entry.Spellid2,
		Spellid3: entry.Spellid3,
	}
}

func (entry *DBSpelldifficultyDbc) ToWeb() *SpelldifficultyDbc {
	if entry == nil {
		return nil
	}
	return &SpelldifficultyDbc{
		Id: entry.Id,
		Spellid0: entry.Spellid0,
		Spellid1: entry.Spellid1,
		Spellid2: entry.Spellid2,
		Spellid3: entry.Spellid3,
	}
}

func (data SpelldifficultyDbcSlice) ToDB() DBSpelldifficultyDbcSlice {
	result := make(DBSpelldifficultyDbcSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpelldifficultyDbcSlice) ToWeb() SpelldifficultyDbcSlice {
	result := make(SpelldifficultyDbcSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
