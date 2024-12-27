package model

import "github.com/uptrace/bun"

type VehicleTemplateAccessory struct {
	Entry int `json:"entry,omitempty"`
	AccessoryEntry int `json:"accessory_entry,omitempty"`
	SeatId int8 `json:"seat_id,omitempty"`
	Minion int8 `json:"minion,omitempty"`
	Description string `json:"description,omitempty"`
	Summontype int8 `json:"summontype,omitempty"`
	Summontimer int `json:"summontimer,omitempty"`
}

type VehicleTemplateAccessorySlice []VehicleTemplateAccessory

type DBVehicleTemplateAccessory struct {
	bun.BaseModel `bun:"table:vehicle_template_accessory,alias:vehicle_template_accessory"`
	Entry int `bun:"entry"`
	AccessoryEntry int `bun:"accessory_entry"`
	SeatId int8 `bun:"seat_id"`
	Minion int8 `bun:"minion"`
	Description string `bun:"description"`
	Summontype int8 `bun:"summontype"`
	Summontimer int `bun:"summontimer"`
}

type DBVehicleTemplateAccessorySlice []DBVehicleTemplateAccessory

func (entry *VehicleTemplateAccessory) ToDB() *DBVehicleTemplateAccessory {
	if entry == nil {
		return nil
	}
	return &DBVehicleTemplateAccessory{
		Entry: entry.Entry,
		AccessoryEntry: entry.AccessoryEntry,
		SeatId: entry.SeatId,
		Minion: entry.Minion,
		Description: entry.Description,
		Summontype: entry.Summontype,
		Summontimer: entry.Summontimer,
	}
}

func (entry *DBVehicleTemplateAccessory) ToWeb() *VehicleTemplateAccessory {
	if entry == nil {
		return nil
	}
	return &VehicleTemplateAccessory{
		Entry: entry.Entry,
		AccessoryEntry: entry.AccessoryEntry,
		SeatId: entry.SeatId,
		Minion: entry.Minion,
		Description: entry.Description,
		Summontype: entry.Summontype,
		Summontimer: entry.Summontimer,
	}
}

func (data VehicleTemplateAccessorySlice) ToDB() DBVehicleTemplateAccessorySlice {
	result := make(DBVehicleTemplateAccessorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVehicleTemplateAccessorySlice) ToWeb() VehicleTemplateAccessorySlice {
	result := make(VehicleTemplateAccessorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
