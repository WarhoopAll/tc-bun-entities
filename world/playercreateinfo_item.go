package model

import "github.com/uptrace/bun"

type PlayercreateinfoItem struct {
	Race int8 `json:"race,omitempty"`
	Class int8 `json:"class,omitempty"`
	Itemid int `json:"itemid,omitempty"`
	Amount int8 `json:"amount,omitempty"`
}

type PlayercreateinfoItemSlice []PlayercreateinfoItem

type DBPlayercreateinfoItem struct {
	bun.BaseModel `bun:"table:playercreateinfo_item,alias:playercreateinfo_item"`
	Race int8 `bun:"race"`
	Class int8 `bun:"class"`
	Itemid int `bun:"itemid"`
	Amount int8 `bun:"amount"`
}

type DBPlayercreateinfoItemSlice []DBPlayercreateinfoItem

func (entry *PlayercreateinfoItem) ToDB() *DBPlayercreateinfoItem {
	if entry == nil {
		return nil
	}
	return &DBPlayercreateinfoItem{
		Race: entry.Race,
		Class: entry.Class,
		Itemid: entry.Itemid,
		Amount: entry.Amount,
	}
}

func (entry *DBPlayercreateinfoItem) ToWeb() *PlayercreateinfoItem {
	if entry == nil {
		return nil
	}
	return &PlayercreateinfoItem{
		Race: entry.Race,
		Class: entry.Class,
		Itemid: entry.Itemid,
		Amount: entry.Amount,
	}
}

func (data PlayercreateinfoItemSlice) ToDB() DBPlayercreateinfoItemSlice {
	result := make(DBPlayercreateinfoItemSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayercreateinfoItemSlice) ToWeb() PlayercreateinfoItemSlice {
	result := make(PlayercreateinfoItemSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
