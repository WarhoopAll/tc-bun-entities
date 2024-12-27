package model

import "github.com/uptrace/bun"

type Addons struct {
	Name string `json:"name,omitempty"`
	Crc int `json:"crc,omitempty"`
}

type AddonsSlice []Addons

type DBAddons struct {
	bun.BaseModel `bun:"table:addons,alias:addons"`
	Name string `bun:"name"`
	Crc int `bun:"crc"`
}

type DBAddonsSlice []DBAddons

func (entry *Addons) ToDB() *DBAddons {
	if entry == nil {
		return nil
	}
	return &DBAddons{
		Name: entry.Name,
		Crc: entry.Crc,
	}
}

func (entry *DBAddons) ToWeb() *Addons {
	if entry == nil {
		return nil
	}
	return &Addons{
		Name: entry.Name,
		Crc: entry.Crc,
	}
}

func (data AddonsSlice) ToDB() DBAddonsSlice {
	result := make(DBAddonsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAddonsSlice) ToWeb() AddonsSlice {
	result := make(AddonsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
