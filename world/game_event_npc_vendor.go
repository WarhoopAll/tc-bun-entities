package model

import "github.com/uptrace/bun"

type GameEventNpcVendor struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Guid int `json:"guid,omitempty"`
	Slot int16 `json:"slot,omitempty"`
	Item int `json:"item,omitempty"`
	Maxcount int `json:"maxcount,omitempty"`
	Incrtime int `json:"incrtime,omitempty"`
	ExtendedCost int `json:"extendedcost,omitempty"`
}

type GameEventNpcVendorSlice []GameEventNpcVendor

type DBGameEventNpcVendor struct {
	bun.BaseModel `bun:"table:game_event_npc_vendor,alias:game_event_npc_vendor"`
	EventEntry int8 `bun:"eventEntry"`
	Guid int `bun:"guid"`
	Slot int16 `bun:"slot"`
	Item int `bun:"item"`
	Maxcount int `bun:"maxcount"`
	Incrtime int `bun:"incrtime"`
	ExtendedCost int `bun:"ExtendedCost"`
}

type DBGameEventNpcVendorSlice []DBGameEventNpcVendor

func (entry *GameEventNpcVendor) ToDB() *DBGameEventNpcVendor {
	if entry == nil {
		return nil
	}
	return &DBGameEventNpcVendor{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
		Slot: entry.Slot,
		Item: entry.Item,
		Maxcount: entry.Maxcount,
		Incrtime: entry.Incrtime,
		ExtendedCost: entry.ExtendedCost,
	}
}

func (entry *DBGameEventNpcVendor) ToWeb() *GameEventNpcVendor {
	if entry == nil {
		return nil
	}
	return &GameEventNpcVendor{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
		Slot: entry.Slot,
		Item: entry.Item,
		Maxcount: entry.Maxcount,
		Incrtime: entry.Incrtime,
		ExtendedCost: entry.ExtendedCost,
	}
}

func (data GameEventNpcVendorSlice) ToDB() DBGameEventNpcVendorSlice {
	result := make(DBGameEventNpcVendorSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventNpcVendorSlice) ToWeb() GameEventNpcVendorSlice {
	result := make(GameEventNpcVendorSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
