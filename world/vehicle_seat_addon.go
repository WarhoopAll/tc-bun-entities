package model

import "github.com/uptrace/bun"

type VehicleSeatAddon struct {
	SeatEntry int `json:"seatentry,omitempty"`
	SeatOrientation float64 `json:"seatorientation,omitempty"`
	ExitParamX float64 `json:"exitparamx,omitempty"`
	ExitParamY float64 `json:"exitparamy,omitempty"`
	ExitParamZ float64 `json:"exitparamz,omitempty"`
	ExitParamO float64 `json:"exitparamo,omitempty"`
	ExitParamValue bool `json:"exitparamvalue,omitempty"`
}

type VehicleSeatAddonSlice []VehicleSeatAddon

type DBVehicleSeatAddon struct {
	bun.BaseModel `bun:"table:vehicle_seat_addon,alias:vehicle_seat_addon"`
	SeatEntry int `bun:"SeatEntry"`
	SeatOrientation float64 `bun:"SeatOrientation"`
	ExitParamX float64 `bun:"ExitParamX"`
	ExitParamY float64 `bun:"ExitParamY"`
	ExitParamZ float64 `bun:"ExitParamZ"`
	ExitParamO float64 `bun:"ExitParamO"`
	ExitParamValue bool `bun:"ExitParamValue"`
}

type DBVehicleSeatAddonSlice []DBVehicleSeatAddon

func (entry *VehicleSeatAddon) ToDB() *DBVehicleSeatAddon {
	if entry == nil {
		return nil
	}
	return &DBVehicleSeatAddon{
		SeatEntry: entry.SeatEntry,
		SeatOrientation: entry.SeatOrientation,
		ExitParamX: entry.ExitParamX,
		ExitParamY: entry.ExitParamY,
		ExitParamZ: entry.ExitParamZ,
		ExitParamO: entry.ExitParamO,
		ExitParamValue: entry.ExitParamValue,
	}
}

func (entry *DBVehicleSeatAddon) ToWeb() *VehicleSeatAddon {
	if entry == nil {
		return nil
	}
	return &VehicleSeatAddon{
		SeatEntry: entry.SeatEntry,
		SeatOrientation: entry.SeatOrientation,
		ExitParamX: entry.ExitParamX,
		ExitParamY: entry.ExitParamY,
		ExitParamZ: entry.ExitParamZ,
		ExitParamO: entry.ExitParamO,
		ExitParamValue: entry.ExitParamValue,
	}
}

func (data VehicleSeatAddonSlice) ToDB() DBVehicleSeatAddonSlice {
	result := make(DBVehicleSeatAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVehicleSeatAddonSlice) ToWeb() VehicleSeatAddonSlice {
	result := make(VehicleSeatAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
