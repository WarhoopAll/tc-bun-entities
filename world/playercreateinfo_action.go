package model

import "github.com/uptrace/bun"

type PlayercreateinfoAction struct {
	Race int8 `json:"race,omitempty"`
	Class int8 `json:"class,omitempty"`
	Button int16 `json:"button,omitempty"`
	Action int `json:"action,omitempty"`
	Type int16 `json:"type,omitempty"`
}

type PlayercreateinfoActionSlice []PlayercreateinfoAction

type DBPlayercreateinfoAction struct {
	bun.BaseModel `bun:"table:playercreateinfo_action,alias:playercreateinfo_action"`
	Race int8 `bun:"race"`
	Class int8 `bun:"class"`
	Button int16 `bun:"button"`
	Action int `bun:"action"`
	Type int16 `bun:"type"`
}

type DBPlayercreateinfoActionSlice []DBPlayercreateinfoAction

func (entry *PlayercreateinfoAction) ToDB() *DBPlayercreateinfoAction {
	if entry == nil {
		return nil
	}
	return &DBPlayercreateinfoAction{
		Race: entry.Race,
		Class: entry.Class,
		Button: entry.Button,
		Action: entry.Action,
		Type: entry.Type,
	}
}

func (entry *DBPlayercreateinfoAction) ToWeb() *PlayercreateinfoAction {
	if entry == nil {
		return nil
	}
	return &PlayercreateinfoAction{
		Race: entry.Race,
		Class: entry.Class,
		Button: entry.Button,
		Action: entry.Action,
		Type: entry.Type,
	}
}

func (data PlayercreateinfoActionSlice) ToDB() DBPlayercreateinfoActionSlice {
	result := make(DBPlayercreateinfoActionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayercreateinfoActionSlice) ToWeb() PlayercreateinfoActionSlice {
	result := make(PlayercreateinfoActionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
