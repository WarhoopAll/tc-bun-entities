package model

import "github.com/uptrace/bun"

type Autobroadcast struct {
	Realmid int `json:"realmid,omitempty"`
	Id int8 `json:"id,omitempty"`
	Weight int8 `json:"weight,omitempty"`
	Text string `json:"text,omitempty"`
}

type AutobroadcastSlice []Autobroadcast

type DBAutobroadcast struct {
	bun.BaseModel `bun:"table:autobroadcast,alias:autobroadcast"`
	Realmid int `bun:"realmid"`
	Id int8 `bun:"id"`
	Weight int8 `bun:"weight"`
	Text string `bun:"text"`
}

type DBAutobroadcastSlice []DBAutobroadcast

func (entry *Autobroadcast) ToDB() *DBAutobroadcast {
	if entry == nil {
		return nil
	}
	return &DBAutobroadcast{
		Realmid: entry.Realmid,
		Id: entry.Id,
		Weight: entry.Weight,
		Text: entry.Text,
	}
}

func (entry *DBAutobroadcast) ToWeb() *Autobroadcast {
	if entry == nil {
		return nil
	}
	return &Autobroadcast{
		Realmid: entry.Realmid,
		Id: entry.Id,
		Weight: entry.Weight,
		Text: entry.Text,
	}
}

func (data AutobroadcastSlice) ToDB() DBAutobroadcastSlice {
	result := make(DBAutobroadcastSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAutobroadcastSlice) ToWeb() AutobroadcastSlice {
	result := make(AutobroadcastSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
