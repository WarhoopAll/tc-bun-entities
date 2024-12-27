package model

import "github.com/uptrace/bun"

type AreatriggerInvolvedrelation struct {
	Id int `json:"id,omitempty"`
	Quest int `json:"quest,omitempty"`
}

type AreatriggerInvolvedrelationSlice []AreatriggerInvolvedrelation

type DBAreatriggerInvolvedrelation struct {
	bun.BaseModel `bun:"table:areatrigger_involvedrelation,alias:areatrigger_involvedrelation"`
	Id int `bun:"id"`
	Quest int `bun:"quest"`
}

type DBAreatriggerInvolvedrelationSlice []DBAreatriggerInvolvedrelation

func (entry *AreatriggerInvolvedrelation) ToDB() *DBAreatriggerInvolvedrelation {
	if entry == nil {
		return nil
	}
	return &DBAreatriggerInvolvedrelation{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (entry *DBAreatriggerInvolvedrelation) ToWeb() *AreatriggerInvolvedrelation {
	if entry == nil {
		return nil
	}
	return &AreatriggerInvolvedrelation{
		Id: entry.Id,
		Quest: entry.Quest,
	}
}

func (data AreatriggerInvolvedrelationSlice) ToDB() DBAreatriggerInvolvedrelationSlice {
	result := make(DBAreatriggerInvolvedrelationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAreatriggerInvolvedrelationSlice) ToWeb() AreatriggerInvolvedrelationSlice {
	result := make(AreatriggerInvolvedrelationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
