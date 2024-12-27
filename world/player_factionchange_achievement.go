package model

import "github.com/uptrace/bun"

type PlayerFactionchangeAchievement struct {
	AllianceId int `json:"alliance_id,omitempty"`
	HordeId int `json:"horde_id,omitempty"`
}

type PlayerFactionchangeAchievementSlice []PlayerFactionchangeAchievement

type DBPlayerFactionchangeAchievement struct {
	bun.BaseModel `bun:"table:player_factionchange_achievement,alias:player_factionchange_achievement"`
	AllianceId int `bun:"alliance_id"`
	HordeId int `bun:"horde_id"`
}

type DBPlayerFactionchangeAchievementSlice []DBPlayerFactionchangeAchievement

func (entry *PlayerFactionchangeAchievement) ToDB() *DBPlayerFactionchangeAchievement {
	if entry == nil {
		return nil
	}
	return &DBPlayerFactionchangeAchievement{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (entry *DBPlayerFactionchangeAchievement) ToWeb() *PlayerFactionchangeAchievement {
	if entry == nil {
		return nil
	}
	return &PlayerFactionchangeAchievement{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (data PlayerFactionchangeAchievementSlice) ToDB() DBPlayerFactionchangeAchievementSlice {
	result := make(DBPlayerFactionchangeAchievementSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerFactionchangeAchievementSlice) ToWeb() PlayerFactionchangeAchievementSlice {
	result := make(PlayerFactionchangeAchievementSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
