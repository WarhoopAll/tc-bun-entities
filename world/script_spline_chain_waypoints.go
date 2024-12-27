package model

import "github.com/uptrace/bun"

type ScriptSplineChainWaypoints struct {
	Entry int `json:"entry,omitempty"`
	ChainId int16 `json:"chainid,omitempty"`
	SplineId int8 `json:"splineid,omitempty"`
	WpId int8 `json:"wpid,omitempty"`
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
}

type ScriptSplineChainWaypointsSlice []ScriptSplineChainWaypoints

type DBScriptSplineChainWaypoints struct {
	bun.BaseModel `bun:"table:script_spline_chain_waypoints,alias:script_spline_chain_waypoints"`
	Entry int `bun:"entry"`
	ChainId int16 `bun:"chainId"`
	SplineId int8 `bun:"splineId"`
	WpId int8 `bun:"wpId"`
	X float64 `bun:"x"`
	Y float64 `bun:"y"`
	Z float64 `bun:"z"`
}

type DBScriptSplineChainWaypointsSlice []DBScriptSplineChainWaypoints

func (entry *ScriptSplineChainWaypoints) ToDB() *DBScriptSplineChainWaypoints {
	if entry == nil {
		return nil
	}
	return &DBScriptSplineChainWaypoints{
		Entry: entry.Entry,
		ChainId: entry.ChainId,
		SplineId: entry.SplineId,
		WpId: entry.WpId,
		X: entry.X,
		Y: entry.Y,
		Z: entry.Z,
	}
}

func (entry *DBScriptSplineChainWaypoints) ToWeb() *ScriptSplineChainWaypoints {
	if entry == nil {
		return nil
	}
	return &ScriptSplineChainWaypoints{
		Entry: entry.Entry,
		ChainId: entry.ChainId,
		SplineId: entry.SplineId,
		WpId: entry.WpId,
		X: entry.X,
		Y: entry.Y,
		Z: entry.Z,
	}
}

func (data ScriptSplineChainWaypointsSlice) ToDB() DBScriptSplineChainWaypointsSlice {
	result := make(DBScriptSplineChainWaypointsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBScriptSplineChainWaypointsSlice) ToWeb() ScriptSplineChainWaypointsSlice {
	result := make(ScriptSplineChainWaypointsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
