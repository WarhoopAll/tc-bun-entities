package model

import "github.com/uptrace/bun"

type AreatriggerTavern struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type AreatriggerTavernSlice []AreatriggerTavern

type DBAreatriggerTavern struct {
	bun.BaseModel `bun:"table:areatrigger_tavern,alias:areatrigger_tavern"`
	Id int `bun:"id"`
	Name string `bun:"name"`
}

type DBAreatriggerTavernSlice []DBAreatriggerTavern

func (entry *AreatriggerTavern) ToDB() *DBAreatriggerTavern {
	if entry == nil {
		return nil
	}
	return &DBAreatriggerTavern{
		Id: entry.Id,
		Name: entry.Name,
	}
}

func (entry *DBAreatriggerTavern) ToWeb() *AreatriggerTavern {
	if entry == nil {
		return nil
	}
	return &AreatriggerTavern{
		Id: entry.Id,
		Name: entry.Name,
	}
}

func (data AreatriggerTavernSlice) ToDB() DBAreatriggerTavernSlice {
	result := make(DBAreatriggerTavernSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAreatriggerTavernSlice) ToWeb() AreatriggerTavernSlice {
	result := make(AreatriggerTavernSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
