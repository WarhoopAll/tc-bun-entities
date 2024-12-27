package model

import "github.com/uptrace/bun"

type WaypointScripts struct {
	Id int `json:"id,omitempty"`
	Delay int `json:"delay,omitempty"`
	Command int `json:"command,omitempty"`
	Datalong int `json:"datalong,omitempty"`
	Datalong2 int `json:"datalong2,omitempty"`
	Dataint int `json:"dataint,omitempty"`
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
	O float64 `json:"o,omitempty"`
	Guid int `json:"guid,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type WaypointScriptsSlice []WaypointScripts

type DBWaypointScripts struct {
	bun.BaseModel `bun:"table:waypoint_scripts,alias:waypoint_scripts"`
	Id int `bun:"id"`
	Delay int `bun:"delay"`
	Command int `bun:"command"`
	Datalong int `bun:"datalong"`
	Datalong2 int `bun:"datalong2"`
	Dataint int `bun:"dataint"`
	X float64 `bun:"x"`
	Y float64 `bun:"y"`
	Z float64 `bun:"z"`
	O float64 `bun:"o"`
	Guid int `bun:"guid"`
	Comment string `bun:"Comment"`
}

type DBWaypointScriptsSlice []DBWaypointScripts

func (entry *WaypointScripts) ToDB() *DBWaypointScripts {
	if entry == nil {
		return nil
	}
	return &DBWaypointScripts{
		Id: entry.Id,
		Delay: entry.Delay,
		Command: entry.Command,
		Datalong: entry.Datalong,
		Datalong2: entry.Datalong2,
		Dataint: entry.Dataint,
		X: entry.X,
		Y: entry.Y,
		Z: entry.Z,
		O: entry.O,
		Guid: entry.Guid,
		Comment: entry.Comment,
	}
}

func (entry *DBWaypointScripts) ToWeb() *WaypointScripts {
	if entry == nil {
		return nil
	}
	return &WaypointScripts{
		Id: entry.Id,
		Delay: entry.Delay,
		Command: entry.Command,
		Datalong: entry.Datalong,
		Datalong2: entry.Datalong2,
		Dataint: entry.Dataint,
		X: entry.X,
		Y: entry.Y,
		Z: entry.Z,
		O: entry.O,
		Guid: entry.Guid,
		Comment: entry.Comment,
	}
}

func (data WaypointScriptsSlice) ToDB() DBWaypointScriptsSlice {
	result := make(DBWaypointScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBWaypointScriptsSlice) ToWeb() WaypointScriptsSlice {
	result := make(WaypointScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
