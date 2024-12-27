package model

import "github.com/uptrace/bun"

type GameEventModelEquip struct {
	EventEntry int8 `json:"evententry,omitempty"`
	Guid int `json:"guid,omitempty"`
	Modelid int `json:"modelid,omitempty"`
	EquipmentId int8 `json:"equipment_id,omitempty"`
}

type GameEventModelEquipSlice []GameEventModelEquip

type DBGameEventModelEquip struct {
	bun.BaseModel `bun:"table:game_event_model_equip,alias:game_event_model_equip"`
	EventEntry int8 `bun:"eventEntry"`
	Guid int `bun:"guid"`
	Modelid int `bun:"modelid"`
	EquipmentId int8 `bun:"equipment_id"`
}

type DBGameEventModelEquipSlice []DBGameEventModelEquip

func (entry *GameEventModelEquip) ToDB() *DBGameEventModelEquip {
	if entry == nil {
		return nil
	}
	return &DBGameEventModelEquip{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
		Modelid: entry.Modelid,
		EquipmentId: entry.EquipmentId,
	}
}

func (entry *DBGameEventModelEquip) ToWeb() *GameEventModelEquip {
	if entry == nil {
		return nil
	}
	return &GameEventModelEquip{
		EventEntry: entry.EventEntry,
		Guid: entry.Guid,
		Modelid: entry.Modelid,
		EquipmentId: entry.EquipmentId,
	}
}

func (data GameEventModelEquipSlice) ToDB() DBGameEventModelEquipSlice {
	result := make(DBGameEventModelEquipSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameEventModelEquipSlice) ToWeb() GameEventModelEquipSlice {
	result := make(GameEventModelEquipSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
