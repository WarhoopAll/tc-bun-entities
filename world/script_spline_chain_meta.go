package model

import "github.com/uptrace/bun"

type ScriptSplineChainMeta struct {
	Entry int `json:"entry,omitempty"`
	ChainId int16 `json:"chainid,omitempty"`
	SplineId int8 `json:"splineid,omitempty"`
	ExpectedDuration int `json:"expectedduration,omitempty"`
	MsUntilNext int `json:"msuntilnext,omitempty"`
	Velocity float64 `json:"velocity,omitempty"`
}

type ScriptSplineChainMetaSlice []ScriptSplineChainMeta

type DBScriptSplineChainMeta struct {
	bun.BaseModel `bun:"table:script_spline_chain_meta,alias:script_spline_chain_meta"`
	Entry int `bun:"entry"`
	ChainId int16 `bun:"chainId"`
	SplineId int8 `bun:"splineId"`
	ExpectedDuration int `bun:"expectedDuration"`
	MsUntilNext int `bun:"msUntilNext"`
	Velocity float64 `bun:"velocity"`
}

type DBScriptSplineChainMetaSlice []DBScriptSplineChainMeta

func (entry *ScriptSplineChainMeta) ToDB() *DBScriptSplineChainMeta {
	if entry == nil {
		return nil
	}
	return &DBScriptSplineChainMeta{
		Entry: entry.Entry,
		ChainId: entry.ChainId,
		SplineId: entry.SplineId,
		ExpectedDuration: entry.ExpectedDuration,
		MsUntilNext: entry.MsUntilNext,
		Velocity: entry.Velocity,
	}
}

func (entry *DBScriptSplineChainMeta) ToWeb() *ScriptSplineChainMeta {
	if entry == nil {
		return nil
	}
	return &ScriptSplineChainMeta{
		Entry: entry.Entry,
		ChainId: entry.ChainId,
		SplineId: entry.SplineId,
		ExpectedDuration: entry.ExpectedDuration,
		MsUntilNext: entry.MsUntilNext,
		Velocity: entry.Velocity,
	}
}

func (data ScriptSplineChainMetaSlice) ToDB() DBScriptSplineChainMetaSlice {
	result := make(DBScriptSplineChainMetaSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBScriptSplineChainMetaSlice) ToWeb() ScriptSplineChainMetaSlice {
	result := make(ScriptSplineChainMetaSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
