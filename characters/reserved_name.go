package model

import "github.com/uptrace/bun"

type ReservedName struct {
	Name string `json:"name,omitempty"`
}

type ReservedNameSlice []ReservedName

type DBReservedName struct {
	bun.BaseModel `bun:"table:reserved_name,alias:reserved_name"`
	Name string `bun:"name"`
}

type DBReservedNameSlice []DBReservedName

func (entry *ReservedName) ToDB() *DBReservedName {
	if entry == nil {
		return nil
	}
	return &DBReservedName{
		Name: entry.Name,
	}
}

func (entry *DBReservedName) ToWeb() *ReservedName {
	if entry == nil {
		return nil
	}
	return &ReservedName{
		Name: entry.Name,
	}
}

func (data ReservedNameSlice) ToDB() DBReservedNameSlice {
	result := make(DBReservedNameSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBReservedNameSlice) ToWeb() ReservedNameSlice {
	result := make(ReservedNameSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
