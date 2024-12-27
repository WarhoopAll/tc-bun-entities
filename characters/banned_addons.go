package model

import "github.com/uptrace/bun"

type BannedAddons struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type BannedAddonsSlice []BannedAddons

type DBBannedAddons struct {
	bun.BaseModel `bun:"table:banned_addons,alias:banned_addons"`
	Id int `bun:"Id"`
	Name string `bun:"Name"`
	Version string `bun:"Version"`
	Timestamp time.Time `bun:"Timestamp"`
}

type DBBannedAddonsSlice []DBBannedAddons

func (entry *BannedAddons) ToDB() *DBBannedAddons {
	if entry == nil {
		return nil
	}
	return &DBBannedAddons{
		Id: entry.Id,
		Name: entry.Name,
		Version: entry.Version,
		Timestamp: entry.Timestamp,
	}
}

func (entry *DBBannedAddons) ToWeb() *BannedAddons {
	if entry == nil {
		return nil
	}
	return &BannedAddons{
		Id: entry.Id,
		Name: entry.Name,
		Version: entry.Version,
		Timestamp: entry.Timestamp,
	}
}

func (data BannedAddonsSlice) ToDB() DBBannedAddonsSlice {
	result := make(DBBannedAddonsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBannedAddonsSlice) ToWeb() BannedAddonsSlice {
	result := make(BannedAddonsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
