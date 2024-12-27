package model

import "github.com/uptrace/bun"

type Instance struct {
	Id int `json:"id,omitempty"`
	Map int16 `json:"map,omitempty"`
	Resettime int `json:"resettime,omitempty"`
	Difficulty int8 `json:"difficulty,omitempty"`
	CompletedEncounters int `json:"completedencounters,omitempty"`
	Data string `json:"data,omitempty"`
}

type InstanceSlice []Instance

type DBInstance struct {
	bun.BaseModel `bun:"table:instance,alias:instance"`
	Id int `bun:"id"`
	Map int16 `bun:"map"`
	Resettime int `bun:"resettime"`
	Difficulty int8 `bun:"difficulty"`
	CompletedEncounters int `bun:"completedEncounters"`
	Data string `bun:"data"`
}

type DBInstanceSlice []DBInstance

func (entry *Instance) ToDB() *DBInstance {
	if entry == nil {
		return nil
	}
	return &DBInstance{
		Id: entry.Id,
		Map: entry.Map,
		Resettime: entry.Resettime,
		Difficulty: entry.Difficulty,
		CompletedEncounters: entry.CompletedEncounters,
		Data: entry.Data,
	}
}

func (entry *DBInstance) ToWeb() *Instance {
	if entry == nil {
		return nil
	}
	return &Instance{
		Id: entry.Id,
		Map: entry.Map,
		Resettime: entry.Resettime,
		Difficulty: entry.Difficulty,
		CompletedEncounters: entry.CompletedEncounters,
		Data: entry.Data,
	}
}

func (data InstanceSlice) ToDB() DBInstanceSlice {
	result := make(DBInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBInstanceSlice) ToWeb() InstanceSlice {
	result := make(InstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
