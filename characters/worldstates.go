package model

import "github.com/uptrace/bun"

type Worldstates struct {
	Entry int `json:"entry,omitempty"`
	Value int `json:"value,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type WorldstatesSlice []Worldstates

type DBWorldstates struct {
	bun.BaseModel `bun:"table:worldstates,alias:worldstates"`
	Entry int `bun:"entry"`
	Value int `bun:"value"`
	Comment string `bun:"comment"`
}

type DBWorldstatesSlice []DBWorldstates

func (entry *Worldstates) ToDB() *DBWorldstates {
	if entry == nil {
		return nil
	}
	return &DBWorldstates{
		Entry: entry.Entry,
		Value: entry.Value,
		Comment: entry.Comment,
	}
}

func (entry *DBWorldstates) ToWeb() *Worldstates {
	if entry == nil {
		return nil
	}
	return &Worldstates{
		Entry: entry.Entry,
		Value: entry.Value,
		Comment: entry.Comment,
	}
}

func (data WorldstatesSlice) ToDB() DBWorldstatesSlice {
	result := make(DBWorldstatesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBWorldstatesSlice) ToWeb() WorldstatesSlice {
	result := make(WorldstatesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
