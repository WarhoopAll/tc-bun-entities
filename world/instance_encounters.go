package model

import "github.com/uptrace/bun"

type InstanceEncounters struct {
	Entry int `json:"entry,omitempty"`
	CreditType int8 `json:"credittype,omitempty"`
	CreditEntry int `json:"creditentry,omitempty"`
	LastEncounterDungeon int16 `json:"lastencounterdungeon,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type InstanceEncountersSlice []InstanceEncounters

type DBInstanceEncounters struct {
	bun.BaseModel `bun:"table:instance_encounters,alias:instance_encounters"`
	Entry int `bun:"entry"`
	CreditType int8 `bun:"creditType"`
	CreditEntry int `bun:"creditEntry"`
	LastEncounterDungeon int16 `bun:"lastEncounterDungeon"`
	Comment string `bun:"comment"`
}

type DBInstanceEncountersSlice []DBInstanceEncounters

func (entry *InstanceEncounters) ToDB() *DBInstanceEncounters {
	if entry == nil {
		return nil
	}
	return &DBInstanceEncounters{
		Entry: entry.Entry,
		CreditType: entry.CreditType,
		CreditEntry: entry.CreditEntry,
		LastEncounterDungeon: entry.LastEncounterDungeon,
		Comment: entry.Comment,
	}
}

func (entry *DBInstanceEncounters) ToWeb() *InstanceEncounters {
	if entry == nil {
		return nil
	}
	return &InstanceEncounters{
		Entry: entry.Entry,
		CreditType: entry.CreditType,
		CreditEntry: entry.CreditEntry,
		LastEncounterDungeon: entry.LastEncounterDungeon,
		Comment: entry.Comment,
	}
}

func (data InstanceEncountersSlice) ToDB() DBInstanceEncountersSlice {
	result := make(DBInstanceEncountersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBInstanceEncountersSlice) ToWeb() InstanceEncountersSlice {
	result := make(InstanceEncountersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
