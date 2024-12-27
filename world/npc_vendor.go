package model

import "github.com/uptrace/bun"

type NpcVendor struct {
	Entry int `json:"entry,omitempty"`
	Slot int16 `json:"slot,omitempty"`
	Item int `json:"item,omitempty"`
	Maxcount int8 `json:"maxcount,omitempty"`
	Incrtime int `json:"incrtime,omitempty"`
	ExtendedCost int `json:"extendedcost,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type NpcVendorSlice []NpcVendor

type DBNpcVendor struct {
	bun.BaseModel `bun:"table:npc_vendor,alias:npc_vendor"`
	Entry int `bun:"entry"`
	Slot int16 `bun:"slot"`
	Item int `bun:"item"`
	Maxcount int8 `bun:"maxcount"`
	Incrtime int `bun:"incrtime"`
	ExtendedCost int `bun:"ExtendedCost"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBNpcVendorSlice []DBNpcVendor

func (entry *NpcVendor) ToDB() *DBNpcVendor {
	if entry == nil {
		return nil
	}
	return &DBNpcVendor{
		Entry: entry.Entry,
		Slot: entry.Slot,
		Item: entry.Item,
		Maxcount: entry.Maxcount,
		Incrtime: entry.Incrtime,
		ExtendedCost: entry.ExtendedCost,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBNpcVendor) ToWeb() *NpcVendor {
	if entry == nil {
		return nil
	}
	return &NpcVendor{
		Entry: entry.Entry,
		Slot: entry.Slot,
		Item: entry.Item,
		Maxcount: entry.Maxcount,
		Incrtime: entry.Incrtime,
		ExtendedCost: entry.ExtendedCost,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data NpcVendorSlice) ToDB() DBNpcVendorSlice {
	result := make(DBNpcVendorSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBNpcVendorSlice) ToWeb() NpcVendorSlice {
	result := make(NpcVendorSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
