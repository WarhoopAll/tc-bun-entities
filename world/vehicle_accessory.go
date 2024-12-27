package model

import "github.com/uptrace/bun"

type VehicleAccessory struct {
	Guid int `json:"guid,omitempty"`
	AccessoryEntry int `json:"accessory_entry,omitempty"`
	SeatId int8 `json:"seat_id,omitempty"`
	Minion int8 `json:"minion,omitempty"`
	Description string `json:"description,omitempty"`
	Summontype int8 `json:"summontype,omitempty"`
	Summontimer int `json:"summontimer,omitempty"`
}

type VehicleAccessorySlice []VehicleAccessory

type DBVehicleAccessory struct {
	bun.BaseModel `bun:"table:vehicle_accessory,alias:vehicle_accessory"`
	Guid int `bun:"guid"`
	AccessoryEntry int `bun:"accessory_entry"`
	SeatId int8 `bun:"seat_id"`
	Minion int8 `bun:"minion"`
	Description string `bun:"description"`
	Summontype int8 `bun:"summontype"`
	Summontimer int `bun:"summontimer"`
}

type DBVehicleAccessorySlice []DBVehicleAccessory

func (entry *VehicleAccessory) ToDB() *DBVehicleAccessory {
	if entry == nil {
		return nil
	}
	return &DBVehicleAccessory{
		Guid: entry.Guid,
		AccessoryEntry: entry.AccessoryEntry,
		SeatId: entry.SeatId,
		Minion: entry.Minion,
		Description: entry.Description,
		Summontype: entry.Summontype,
		Summontimer: entry.Summontimer,
	}
}

func (entry *DBVehicleAccessory) ToWeb() *VehicleAccessory {
	if entry == nil {
		return nil
	}
	return &VehicleAccessory{
		Guid: entry.Guid,
		AccessoryEntry: entry.AccessoryEntry,
		SeatId: entry.SeatId,
		Minion: entry.Minion,
		Description: entry.Description,
		Summontype: entry.Summontype,
		Summontimer: entry.Summontimer,
	}
}

func (data VehicleAccessorySlice) ToDB() DBVehicleAccessorySlice {
	result := make(DBVehicleAccessorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVehicleAccessorySlice) ToWeb() VehicleAccessorySlice {
	result := make(VehicleAccessorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
