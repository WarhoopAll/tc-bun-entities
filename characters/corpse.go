package model

import "github.com/uptrace/bun"

type Corpse struct {
	Guid int `json:"guid,omitempty"`
	PosX float64 `json:"posx,omitempty"`
	PosY float64 `json:"posy,omitempty"`
	PosZ float64 `json:"posz,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	MapId int16 `json:"mapid,omitempty"`
	PhaseMask int `json:"phasemask,omitempty"`
	DisplayId int `json:"displayid,omitempty"`
	ItemCache string `json:"itemcache,omitempty"`
	Bytes1 int `json:"bytes1,omitempty"`
	Bytes2 int `json:"bytes2,omitempty"`
	GuildId int `json:"guildid,omitempty"`
	Flags int8 `json:"flags,omitempty"`
	DynFlags int8 `json:"dynflags,omitempty"`
	Time int `json:"time,omitempty"`
	CorpseType int8 `json:"corpsetype,omitempty"`
	InstanceId int `json:"instanceid,omitempty"`
}

type CorpseSlice []Corpse

type DBCorpse struct {
	bun.BaseModel `bun:"table:corpse,alias:corpse"`
	Guid int `bun:"guid"`
	PosX float64 `bun:"posX"`
	PosY float64 `bun:"posY"`
	PosZ float64 `bun:"posZ"`
	Orientation float64 `bun:"orientation"`
	MapId int16 `bun:"mapId"`
	PhaseMask int `bun:"phaseMask"`
	DisplayId int `bun:"displayId"`
	ItemCache string `bun:"itemCache"`
	Bytes1 int `bun:"bytes1"`
	Bytes2 int `bun:"bytes2"`
	GuildId int `bun:"guildId"`
	Flags int8 `bun:"flags"`
	DynFlags int8 `bun:"dynFlags"`
	Time int `bun:"time"`
	CorpseType int8 `bun:"corpseType"`
	InstanceId int `bun:"instanceId"`
}

type DBCorpseSlice []DBCorpse

func (entry *Corpse) ToDB() *DBCorpse {
	if entry == nil {
		return nil
	}
	return &DBCorpse{
		Guid: entry.Guid,
		PosX: entry.PosX,
		PosY: entry.PosY,
		PosZ: entry.PosZ,
		Orientation: entry.Orientation,
		MapId: entry.MapId,
		PhaseMask: entry.PhaseMask,
		DisplayId: entry.DisplayId,
		ItemCache: entry.ItemCache,
		Bytes1: entry.Bytes1,
		Bytes2: entry.Bytes2,
		GuildId: entry.GuildId,
		Flags: entry.Flags,
		DynFlags: entry.DynFlags,
		Time: entry.Time,
		CorpseType: entry.CorpseType,
		InstanceId: entry.InstanceId,
	}
}

func (entry *DBCorpse) ToWeb() *Corpse {
	if entry == nil {
		return nil
	}
	return &Corpse{
		Guid: entry.Guid,
		PosX: entry.PosX,
		PosY: entry.PosY,
		PosZ: entry.PosZ,
		Orientation: entry.Orientation,
		MapId: entry.MapId,
		PhaseMask: entry.PhaseMask,
		DisplayId: entry.DisplayId,
		ItemCache: entry.ItemCache,
		Bytes1: entry.Bytes1,
		Bytes2: entry.Bytes2,
		GuildId: entry.GuildId,
		Flags: entry.Flags,
		DynFlags: entry.DynFlags,
		Time: entry.Time,
		CorpseType: entry.CorpseType,
		InstanceId: entry.InstanceId,
	}
}

func (data CorpseSlice) ToDB() DBCorpseSlice {
	result := make(DBCorpseSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCorpseSlice) ToWeb() CorpseSlice {
	result := make(CorpseSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
