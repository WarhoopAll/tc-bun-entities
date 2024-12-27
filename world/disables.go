package model

import "github.com/uptrace/bun"

type Disables struct {
	SourceType int `json:"sourcetype,omitempty"`
	Entry int `json:"entry,omitempty"`
	Flags int16 `json:"flags,omitempty"`
	Params0 string `json:"params_0,omitempty"`
	Params1 string `json:"params_1,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type DisablesSlice []Disables

type DBDisables struct {
	bun.BaseModel `bun:"table:disables,alias:disables"`
	SourceType int `bun:"sourceType"`
	Entry int `bun:"entry"`
	Flags int16 `bun:"flags"`
	Params0 string `bun:"params_0"`
	Params1 string `bun:"params_1"`
	Comment string `bun:"comment"`
}

type DBDisablesSlice []DBDisables

func (entry *Disables) ToDB() *DBDisables {
	if entry == nil {
		return nil
	}
	return &DBDisables{
		SourceType: entry.SourceType,
		Entry: entry.Entry,
		Flags: entry.Flags,
		Params0: entry.Params0,
		Params1: entry.Params1,
		Comment: entry.Comment,
	}
}

func (entry *DBDisables) ToWeb() *Disables {
	if entry == nil {
		return nil
	}
	return &Disables{
		SourceType: entry.SourceType,
		Entry: entry.Entry,
		Flags: entry.Flags,
		Params0: entry.Params0,
		Params1: entry.Params1,
		Comment: entry.Comment,
	}
}

func (data DisablesSlice) ToDB() DBDisablesSlice {
	result := make(DBDisablesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBDisablesSlice) ToWeb() DisablesSlice {
	result := make(DisablesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
