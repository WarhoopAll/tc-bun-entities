package model

import "github.com/uptrace/bun"

type WardenAction struct {
	WardenId int16 `json:"wardenid,omitempty"`
	Action int8 `json:"action,omitempty"`
}

type WardenActionSlice []WardenAction

type DBWardenAction struct {
	bun.BaseModel `bun:"table:warden_action,alias:warden_action"`
	WardenId int16 `bun:"wardenId"`
	Action int8 `bun:"action"`
}

type DBWardenActionSlice []DBWardenAction

func (entry *WardenAction) ToDB() *DBWardenAction {
	if entry == nil {
		return nil
	}
	return &DBWardenAction{
		WardenId: entry.WardenId,
		Action: entry.Action,
	}
}

func (entry *DBWardenAction) ToWeb() *WardenAction {
	if entry == nil {
		return nil
	}
	return &WardenAction{
		WardenId: entry.WardenId,
		Action: entry.Action,
	}
}

func (data WardenActionSlice) ToDB() DBWardenActionSlice {
	result := make(DBWardenActionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBWardenActionSlice) ToWeb() WardenActionSlice {
	result := make(WardenActionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
