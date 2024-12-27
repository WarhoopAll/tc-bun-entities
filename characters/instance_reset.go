package model

import "github.com/uptrace/bun"

type InstanceReset struct {
	Mapid int16 `json:"mapid,omitempty"`
	Difficulty int8 `json:"difficulty,omitempty"`
	Resettime int `json:"resettime,omitempty"`
}

type InstanceResetSlice []InstanceReset

type DBInstanceReset struct {
	bun.BaseModel `bun:"table:instance_reset,alias:instance_reset"`
	Mapid int16 `bun:"mapid"`
	Difficulty int8 `bun:"difficulty"`
	Resettime int `bun:"resettime"`
}

type DBInstanceResetSlice []DBInstanceReset

func (entry *InstanceReset) ToDB() *DBInstanceReset {
	if entry == nil {
		return nil
	}
	return &DBInstanceReset{
		Mapid: entry.Mapid,
		Difficulty: entry.Difficulty,
		Resettime: entry.Resettime,
	}
}

func (entry *DBInstanceReset) ToWeb() *InstanceReset {
	if entry == nil {
		return nil
	}
	return &InstanceReset{
		Mapid: entry.Mapid,
		Difficulty: entry.Difficulty,
		Resettime: entry.Resettime,
	}
}

func (data InstanceResetSlice) ToDB() DBInstanceResetSlice {
	result := make(DBInstanceResetSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBInstanceResetSlice) ToWeb() InstanceResetSlice {
	result := make(InstanceResetSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
