package model

import "github.com/uptrace/bun"

type BattlemasterEntry struct {
	Entry int `json:"entry,omitempty"`
	BgTemplate int `json:"bg_template,omitempty"`
}

type BattlemasterEntrySlice []BattlemasterEntry

type DBBattlemasterEntry struct {
	bun.BaseModel `bun:"table:battlemaster_entry,alias:battlemaster_entry"`
	Entry int `bun:"entry"`
	BgTemplate int `bun:"bg_template"`
}

type DBBattlemasterEntrySlice []DBBattlemasterEntry

func (entry *BattlemasterEntry) ToDB() *DBBattlemasterEntry {
	if entry == nil {
		return nil
	}
	return &DBBattlemasterEntry{
		Entry: entry.Entry,
		BgTemplate: entry.BgTemplate,
	}
}

func (entry *DBBattlemasterEntry) ToWeb() *BattlemasterEntry {
	if entry == nil {
		return nil
	}
	return &BattlemasterEntry{
		Entry: entry.Entry,
		BgTemplate: entry.BgTemplate,
	}
}

func (data BattlemasterEntrySlice) ToDB() DBBattlemasterEntrySlice {
	result := make(DBBattlemasterEntrySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBattlemasterEntrySlice) ToWeb() BattlemasterEntrySlice {
	result := make(BattlemasterEntrySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
