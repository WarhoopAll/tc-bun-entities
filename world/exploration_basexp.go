package model

import "github.com/uptrace/bun"

type ExplorationBasexp struct {
	Level int8 `json:"level,omitempty"`
	Basexp int `json:"basexp,omitempty"`
}

type ExplorationBasexpSlice []ExplorationBasexp

type DBExplorationBasexp struct {
	bun.BaseModel `bun:"table:exploration_basexp,alias:exploration_basexp"`
	Level int8 `bun:"level"`
	Basexp int `bun:"basexp"`
}

type DBExplorationBasexpSlice []DBExplorationBasexp

func (entry *ExplorationBasexp) ToDB() *DBExplorationBasexp {
	if entry == nil {
		return nil
	}
	return &DBExplorationBasexp{
		Level: entry.Level,
		Basexp: entry.Basexp,
	}
}

func (entry *DBExplorationBasexp) ToWeb() *ExplorationBasexp {
	if entry == nil {
		return nil
	}
	return &ExplorationBasexp{
		Level: entry.Level,
		Basexp: entry.Basexp,
	}
}

func (data ExplorationBasexpSlice) ToDB() DBExplorationBasexpSlice {
	result := make(DBExplorationBasexpSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBExplorationBasexpSlice) ToWeb() ExplorationBasexpSlice {
	result := make(ExplorationBasexpSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
