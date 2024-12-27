package model

import "github.com/uptrace/bun"

type CreatureModelInfo struct {
	DisplayID int `json:"displayid,omitempty"`
	BoundingRadius float64 `json:"boundingradius,omitempty"`
	CombatReach float64 `json:"combatreach,omitempty"`
	Gender int8 `json:"gender,omitempty"`
	DisplayIDOtherGender int `json:"displayid_other_gender,omitempty"`
}

type CreatureModelInfoSlice []CreatureModelInfo

type DBCreatureModelInfo struct {
	bun.BaseModel `bun:"table:creature_model_info,alias:creature_model_info"`
	DisplayID int `bun:"DisplayID"`
	BoundingRadius float64 `bun:"BoundingRadius"`
	CombatReach float64 `bun:"CombatReach"`
	Gender int8 `bun:"Gender"`
	DisplayIDOtherGender int `bun:"DisplayID_Other_Gender"`
}

type DBCreatureModelInfoSlice []DBCreatureModelInfo

func (entry *CreatureModelInfo) ToDB() *DBCreatureModelInfo {
	if entry == nil {
		return nil
	}
	return &DBCreatureModelInfo{
		DisplayID: entry.DisplayID,
		BoundingRadius: entry.BoundingRadius,
		CombatReach: entry.CombatReach,
		Gender: entry.Gender,
		DisplayIDOtherGender: entry.DisplayIDOtherGender,
	}
}

func (entry *DBCreatureModelInfo) ToWeb() *CreatureModelInfo {
	if entry == nil {
		return nil
	}
	return &CreatureModelInfo{
		DisplayID: entry.DisplayID,
		BoundingRadius: entry.BoundingRadius,
		CombatReach: entry.CombatReach,
		Gender: entry.Gender,
		DisplayIDOtherGender: entry.DisplayIDOtherGender,
	}
}

func (data CreatureModelInfoSlice) ToDB() DBCreatureModelInfoSlice {
	result := make(DBCreatureModelInfoSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureModelInfoSlice) ToWeb() CreatureModelInfoSlice {
	result := make(CreatureModelInfoSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
