package model

import "github.com/uptrace/bun"

type ScriptWaypoint struct {
	Entry int `json:"entry,omitempty"`
	Pointid int `json:"pointid,omitempty"`
	LocationX float64 `json:"location_x,omitempty"`
	LocationY float64 `json:"location_y,omitempty"`
	LocationZ float64 `json:"location_z,omitempty"`
	Waittime int `json:"waittime,omitempty"`
	PointComment string `json:"point_comment,omitempty"`
}

type ScriptWaypointSlice []ScriptWaypoint

type DBScriptWaypoint struct {
	bun.BaseModel `bun:"table:script_waypoint,alias:script_waypoint"`
	Entry int `bun:"entry"`
	Pointid int `bun:"pointid"`
	LocationX float64 `bun:"location_x"`
	LocationY float64 `bun:"location_y"`
	LocationZ float64 `bun:"location_z"`
	Waittime int `bun:"waittime"`
	PointComment string `bun:"point_comment"`
}

type DBScriptWaypointSlice []DBScriptWaypoint

func (entry *ScriptWaypoint) ToDB() *DBScriptWaypoint {
	if entry == nil {
		return nil
	}
	return &DBScriptWaypoint{
		Entry: entry.Entry,
		Pointid: entry.Pointid,
		LocationX: entry.LocationX,
		LocationY: entry.LocationY,
		LocationZ: entry.LocationZ,
		Waittime: entry.Waittime,
		PointComment: entry.PointComment,
	}
}

func (entry *DBScriptWaypoint) ToWeb() *ScriptWaypoint {
	if entry == nil {
		return nil
	}
	return &ScriptWaypoint{
		Entry: entry.Entry,
		Pointid: entry.Pointid,
		LocationX: entry.LocationX,
		LocationY: entry.LocationY,
		LocationZ: entry.LocationZ,
		Waittime: entry.Waittime,
		PointComment: entry.PointComment,
	}
}

func (data ScriptWaypointSlice) ToDB() DBScriptWaypointSlice {
	result := make(DBScriptWaypointSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBScriptWaypointSlice) ToWeb() ScriptWaypointSlice {
	result := make(ScriptWaypointSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
