package model

import "github.com/uptrace/bun"

type Transports struct {
	Guid int `json:"guid,omitempty"`
	Entry int `json:"entry,omitempty"`
	Name string `json:"name,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
}

type TransportsSlice []Transports

type DBTransports struct {
	bun.BaseModel `bun:"table:transports,alias:transports"`
	Guid int `bun:"guid"`
	Entry int `bun:"entry"`
	Name string `bun:"name"`
	ScriptName string `bun:"ScriptName"`
}

type DBTransportsSlice []DBTransports

func (entry *Transports) ToDB() *DBTransports {
	if entry == nil {
		return nil
	}
	return &DBTransports{
		Guid: entry.Guid,
		Entry: entry.Entry,
		Name: entry.Name,
		ScriptName: entry.ScriptName,
	}
}

func (entry *DBTransports) ToWeb() *Transports {
	if entry == nil {
		return nil
	}
	return &Transports{
		Guid: entry.Guid,
		Entry: entry.Entry,
		Name: entry.Name,
		ScriptName: entry.ScriptName,
	}
}

func (data TransportsSlice) ToDB() DBTransportsSlice {
	result := make(DBTransportsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBTransportsSlice) ToWeb() TransportsSlice {
	result := make(TransportsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
