package model

import "github.com/uptrace/bun"

type LagReports struct {
	ReportId int `json:"reportid,omitempty"`
	Guid int `json:"guid,omitempty"`
	LagType int8 `json:"lagtype,omitempty"`
	MapId int16 `json:"mapid,omitempty"`
	PosX float64 `json:"posx,omitempty"`
	PosY float64 `json:"posy,omitempty"`
	PosZ float64 `json:"posz,omitempty"`
	Latency int `json:"latency,omitempty"`
	CreateTime int `json:"createtime,omitempty"`
}

type LagReportsSlice []LagReports

type DBLagReports struct {
	bun.BaseModel `bun:"table:lag_reports,alias:lag_reports"`
	ReportId int `bun:"reportId"`
	Guid int `bun:"guid"`
	LagType int8 `bun:"lagType"`
	MapId int16 `bun:"mapId"`
	PosX float64 `bun:"posX"`
	PosY float64 `bun:"posY"`
	PosZ float64 `bun:"posZ"`
	Latency int `bun:"latency"`
	CreateTime int `bun:"createTime"`
}

type DBLagReportsSlice []DBLagReports

func (entry *LagReports) ToDB() *DBLagReports {
	if entry == nil {
		return nil
	}
	return &DBLagReports{
		ReportId: entry.ReportId,
		Guid: entry.Guid,
		LagType: entry.LagType,
		MapId: entry.MapId,
		PosX: entry.PosX,
		PosY: entry.PosY,
		PosZ: entry.PosZ,
		Latency: entry.Latency,
		CreateTime: entry.CreateTime,
	}
}

func (entry *DBLagReports) ToWeb() *LagReports {
	if entry == nil {
		return nil
	}
	return &LagReports{
		ReportId: entry.ReportId,
		Guid: entry.Guid,
		LagType: entry.LagType,
		MapId: entry.MapId,
		PosX: entry.PosX,
		PosY: entry.PosY,
		PosZ: entry.PosZ,
		Latency: entry.Latency,
		CreateTime: entry.CreateTime,
	}
}

func (data LagReportsSlice) ToDB() DBLagReportsSlice {
	result := make(DBLagReportsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBLagReportsSlice) ToWeb() LagReportsSlice {
	result := make(LagReportsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
