package model

import "github.com/uptrace/bun"

type AchievementDbc struct {
	ID int `json:"id,omitempty"`
	RequiredFaction int `json:"requiredfaction,omitempty"`
	MapID int `json:"mapid,omitempty"`
	Points int `json:"points,omitempty"`
	Flags int `json:"flags,omitempty"`
	Count int `json:"count,omitempty"`
	RefAchievement int `json:"refachievement,omitempty"`
}

type AchievementDbcSlice []AchievementDbc

type DBAchievementDbc struct {
	bun.BaseModel `bun:"table:achievement_dbc,alias:achievement_dbc"`
	ID int `bun:"ID"`
	RequiredFaction int `bun:"requiredFaction"`
	MapID int `bun:"mapID"`
	Points int `bun:"points"`
	Flags int `bun:"flags"`
	Count int `bun:"count"`
	RefAchievement int `bun:"refAchievement"`
}

type DBAchievementDbcSlice []DBAchievementDbc

func (entry *AchievementDbc) ToDB() *DBAchievementDbc {
	if entry == nil {
		return nil
	}
	return &DBAchievementDbc{
		ID: entry.ID,
		RequiredFaction: entry.RequiredFaction,
		MapID: entry.MapID,
		Points: entry.Points,
		Flags: entry.Flags,
		Count: entry.Count,
		RefAchievement: entry.RefAchievement,
	}
}

func (entry *DBAchievementDbc) ToWeb() *AchievementDbc {
	if entry == nil {
		return nil
	}
	return &AchievementDbc{
		ID: entry.ID,
		RequiredFaction: entry.RequiredFaction,
		MapID: entry.MapID,
		Points: entry.Points,
		Flags: entry.Flags,
		Count: entry.Count,
		RefAchievement: entry.RefAchievement,
	}
}

func (data AchievementDbcSlice) ToDB() DBAchievementDbcSlice {
	result := make(DBAchievementDbcSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAchievementDbcSlice) ToWeb() AchievementDbcSlice {
	result := make(AchievementDbcSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
