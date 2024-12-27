package model

import "github.com/uptrace/bun"

type TrinityString struct {
	Entry int `json:"entry,omitempty"`
	ContentDefault string `json:"content_default,omitempty"`
	ContentLoc1 string `json:"content_loc1,omitempty"`
	ContentLoc2 string `json:"content_loc2,omitempty"`
	ContentLoc3 string `json:"content_loc3,omitempty"`
	ContentLoc4 string `json:"content_loc4,omitempty"`
	ContentLoc5 string `json:"content_loc5,omitempty"`
	ContentLoc6 string `json:"content_loc6,omitempty"`
	ContentLoc7 string `json:"content_loc7,omitempty"`
	ContentLoc8 string `json:"content_loc8,omitempty"`
}

type TrinityStringSlice []TrinityString

type DBTrinityString struct {
	bun.BaseModel `bun:"table:trinity_string,alias:trinity_string"`
	Entry int `bun:"entry"`
	ContentDefault string `bun:"content_default"`
	ContentLoc1 string `bun:"content_loc1"`
	ContentLoc2 string `bun:"content_loc2"`
	ContentLoc3 string `bun:"content_loc3"`
	ContentLoc4 string `bun:"content_loc4"`
	ContentLoc5 string `bun:"content_loc5"`
	ContentLoc6 string `bun:"content_loc6"`
	ContentLoc7 string `bun:"content_loc7"`
	ContentLoc8 string `bun:"content_loc8"`
}

type DBTrinityStringSlice []DBTrinityString

func (entry *TrinityString) ToDB() *DBTrinityString {
	if entry == nil {
		return nil
	}
	return &DBTrinityString{
		Entry: entry.Entry,
		ContentDefault: entry.ContentDefault,
		ContentLoc1: entry.ContentLoc1,
		ContentLoc2: entry.ContentLoc2,
		ContentLoc3: entry.ContentLoc3,
		ContentLoc4: entry.ContentLoc4,
		ContentLoc5: entry.ContentLoc5,
		ContentLoc6: entry.ContentLoc6,
		ContentLoc7: entry.ContentLoc7,
		ContentLoc8: entry.ContentLoc8,
	}
}

func (entry *DBTrinityString) ToWeb() *TrinityString {
	if entry == nil {
		return nil
	}
	return &TrinityString{
		Entry: entry.Entry,
		ContentDefault: entry.ContentDefault,
		ContentLoc1: entry.ContentLoc1,
		ContentLoc2: entry.ContentLoc2,
		ContentLoc3: entry.ContentLoc3,
		ContentLoc4: entry.ContentLoc4,
		ContentLoc5: entry.ContentLoc5,
		ContentLoc6: entry.ContentLoc6,
		ContentLoc7: entry.ContentLoc7,
		ContentLoc8: entry.ContentLoc8,
	}
}

func (data TrinityStringSlice) ToDB() DBTrinityStringSlice {
	result := make(DBTrinityStringSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBTrinityStringSlice) ToWeb() TrinityStringSlice {
	result := make(TrinityStringSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
