package model

import "github.com/uptrace/bun"

type VehicleTemplate struct {
	CreatureId int `json:"creatureid,omitempty"`
	DespawnDelayMs int `json:"despawndelayms,omitempty"`
}

type VehicleTemplateSlice []VehicleTemplate

type DBVehicleTemplate struct {
	bun.BaseModel `bun:"table:vehicle_template,alias:vehicle_template"`
	CreatureId int `bun:"creatureId"`
	DespawnDelayMs int `bun:"despawnDelayMs"`
}

type DBVehicleTemplateSlice []DBVehicleTemplate

func (entry *VehicleTemplate) ToDB() *DBVehicleTemplate {
	if entry == nil {
		return nil
	}
	return &DBVehicleTemplate{
		CreatureId: entry.CreatureId,
		DespawnDelayMs: entry.DespawnDelayMs,
	}
}

func (entry *DBVehicleTemplate) ToWeb() *VehicleTemplate {
	if entry == nil {
		return nil
	}
	return &VehicleTemplate{
		CreatureId: entry.CreatureId,
		DespawnDelayMs: entry.DespawnDelayMs,
	}
}

func (data VehicleTemplateSlice) ToDB() DBVehicleTemplateSlice {
	result := make(DBVehicleTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVehicleTemplateSlice) ToWeb() VehicleTemplateSlice {
	result := make(VehicleTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
