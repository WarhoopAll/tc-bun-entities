package model

import "github.com/uptrace/bun"

type WaypointData struct {
	Id int `json:"id,omitempty"`
	Point int `json:"point,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	Delay int `json:"delay,omitempty"`
	MoveType int `json:"move_type,omitempty"`
	Action int `json:"action,omitempty"`
	ActionChance int16 `json:"action_chance,omitempty"`
	Wpguid int `json:"wpguid,omitempty"`
}

type WaypointDataSlice []WaypointData

type DBWaypointData struct {
	bun.BaseModel `bun:"table:waypoint_data,alias:waypoint_data"`
	Id int `bun:"id"`
	Point int `bun:"point"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Orientation float64 `bun:"orientation"`
	Delay int `bun:"delay"`
	MoveType int `bun:"move_type"`
	Action int `bun:"action"`
	ActionChance int16 `bun:"action_chance"`
	Wpguid int `bun:"wpguid"`
}

type DBWaypointDataSlice []DBWaypointData

func (entry *WaypointData) ToDB() *DBWaypointData {
	if entry == nil {
		return nil
	}
	return &DBWaypointData{
		Id: entry.Id,
		Point: entry.Point,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Delay: entry.Delay,
		MoveType: entry.MoveType,
		Action: entry.Action,
		ActionChance: entry.ActionChance,
		Wpguid: entry.Wpguid,
	}
}

func (entry *DBWaypointData) ToWeb() *WaypointData {
	if entry == nil {
		return nil
	}
	return &WaypointData{
		Id: entry.Id,
		Point: entry.Point,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Delay: entry.Delay,
		MoveType: entry.MoveType,
		Action: entry.Action,
		ActionChance: entry.ActionChance,
		Wpguid: entry.Wpguid,
	}
}

func (data WaypointDataSlice) ToDB() DBWaypointDataSlice {
	result := make(DBWaypointDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBWaypointDataSlice) ToWeb() WaypointDataSlice {
	result := make(WaypointDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
