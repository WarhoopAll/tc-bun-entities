package model

import "github.com/uptrace/bun"

type BattlegroundTemplate struct {
	ID int `json:"id,omitempty"`
	MinPlayersPerTeam int16 `json:"minplayersperteam,omitempty"`
	MaxPlayersPerTeam int16 `json:"maxplayersperteam,omitempty"`
	MinLvl int8 `json:"minlvl,omitempty"`
	MaxLvl int8 `json:"maxlvl,omitempty"`
	AllianceStartLoc int `json:"alliancestartloc,omitempty"`
	AllianceStartO float64 `json:"alliancestarto,omitempty"`
	HordeStartLoc int `json:"hordestartloc,omitempty"`
	HordeStartO float64 `json:"hordestarto,omitempty"`
	StartMaxDist float64 `json:"startmaxdist,omitempty"`
	Weight int8 `json:"weight,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type BattlegroundTemplateSlice []BattlegroundTemplate

type DBBattlegroundTemplate struct {
	bun.BaseModel `bun:"table:battleground_template,alias:battleground_template"`
	ID int `bun:"ID"`
	MinPlayersPerTeam int16 `bun:"MinPlayersPerTeam"`
	MaxPlayersPerTeam int16 `bun:"MaxPlayersPerTeam"`
	MinLvl int8 `bun:"MinLvl"`
	MaxLvl int8 `bun:"MaxLvl"`
	AllianceStartLoc int `bun:"AllianceStartLoc"`
	AllianceStartO float64 `bun:"AllianceStartO"`
	HordeStartLoc int `bun:"HordeStartLoc"`
	HordeStartO float64 `bun:"HordeStartO"`
	StartMaxDist float64 `bun:"StartMaxDist"`
	Weight int8 `bun:"Weight"`
	ScriptName string `bun:"ScriptName"`
	Comment string `bun:"Comment"`
}

type DBBattlegroundTemplateSlice []DBBattlegroundTemplate

func (entry *BattlegroundTemplate) ToDB() *DBBattlegroundTemplate {
	if entry == nil {
		return nil
	}
	return &DBBattlegroundTemplate{
		ID: entry.ID,
		MinPlayersPerTeam: entry.MinPlayersPerTeam,
		MaxPlayersPerTeam: entry.MaxPlayersPerTeam,
		MinLvl: entry.MinLvl,
		MaxLvl: entry.MaxLvl,
		AllianceStartLoc: entry.AllianceStartLoc,
		AllianceStartO: entry.AllianceStartO,
		HordeStartLoc: entry.HordeStartLoc,
		HordeStartO: entry.HordeStartO,
		StartMaxDist: entry.StartMaxDist,
		Weight: entry.Weight,
		ScriptName: entry.ScriptName,
		Comment: entry.Comment,
	}
}

func (entry *DBBattlegroundTemplate) ToWeb() *BattlegroundTemplate {
	if entry == nil {
		return nil
	}
	return &BattlegroundTemplate{
		ID: entry.ID,
		MinPlayersPerTeam: entry.MinPlayersPerTeam,
		MaxPlayersPerTeam: entry.MaxPlayersPerTeam,
		MinLvl: entry.MinLvl,
		MaxLvl: entry.MaxLvl,
		AllianceStartLoc: entry.AllianceStartLoc,
		AllianceStartO: entry.AllianceStartO,
		HordeStartLoc: entry.HordeStartLoc,
		HordeStartO: entry.HordeStartO,
		StartMaxDist: entry.StartMaxDist,
		Weight: entry.Weight,
		ScriptName: entry.ScriptName,
		Comment: entry.Comment,
	}
}

func (data BattlegroundTemplateSlice) ToDB() DBBattlegroundTemplateSlice {
	result := make(DBBattlegroundTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBattlegroundTemplateSlice) ToWeb() BattlegroundTemplateSlice {
	result := make(BattlegroundTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
