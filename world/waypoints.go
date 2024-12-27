package model

import "github.com/uptrace/bun"

type Waypoints struct {
	Entry int `json:"entry,omitempty"`
	Pointid int `json:"pointid,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	Delay int `json:"delay,omitempty"`
	PointComment string `json:"point_comment,omitempty"`
}

type WaypointsSlice []Waypoints

type DBWaypoints struct {
	bun.BaseModel `bun:"table:waypoints,alias:waypoints"`
	Entry int `bun:"entry"`
	Pointid int `bun:"pointid"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Orientation float64 `bun:"orientation"`
	Delay int `bun:"delay"`
	PointComment string `bun:"point_comment"`
}

type DBWaypointsSlice []DBWaypoints

func (entry *Waypoints) ToDB() *DBWaypoints {
	if entry == nil {
		return nil
	}
	return &DBWaypoints{
		Entry: entry.Entry,
		Pointid: entry.Pointid,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Delay: entry.Delay,
		PointComment: entry.PointComment,
	}
}

func (entry *DBWaypoints) ToWeb() *Waypoints {
	if entry == nil {
		return nil
	}
	return &Waypoints{
		Entry: entry.Entry,
		Pointid: entry.Pointid,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Delay: entry.Delay,
		PointComment: entry.PointComment,
	}
}

func (data WaypointsSlice) ToDB() DBWaypointsSlice {
	result := make(DBWaypointsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBWaypointsSlice) ToWeb() WaypointsSlice {
	result := make(WaypointsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
