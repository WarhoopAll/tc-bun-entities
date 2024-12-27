package model

import "github.com/uptrace/bun"

type AchievementCriteriaData struct {
	CriteriaId int `json:"criteria_id,omitempty"`
	Type int8 `json:"type,omitempty"`
	Value1 int `json:"value1,omitempty"`
	Value2 int `json:"value2,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
}

type AchievementCriteriaDataSlice []AchievementCriteriaData

type DBAchievementCriteriaData struct {
	bun.BaseModel `bun:"table:achievement_criteria_data,alias:achievement_criteria_data"`
	CriteriaId int `bun:"criteria_id"`
	Type int8 `bun:"type"`
	Value1 int `bun:"value1"`
	Value2 int `bun:"value2"`
	ScriptName string `bun:"ScriptName"`
}

type DBAchievementCriteriaDataSlice []DBAchievementCriteriaData

func (entry *AchievementCriteriaData) ToDB() *DBAchievementCriteriaData {
	if entry == nil {
		return nil
	}
	return &DBAchievementCriteriaData{
		CriteriaId: entry.CriteriaId,
		Type: entry.Type,
		Value1: entry.Value1,
		Value2: entry.Value2,
		ScriptName: entry.ScriptName,
	}
}

func (entry *DBAchievementCriteriaData) ToWeb() *AchievementCriteriaData {
	if entry == nil {
		return nil
	}
	return &AchievementCriteriaData{
		CriteriaId: entry.CriteriaId,
		Type: entry.Type,
		Value1: entry.Value1,
		Value2: entry.Value2,
		ScriptName: entry.ScriptName,
	}
}

func (data AchievementCriteriaDataSlice) ToDB() DBAchievementCriteriaDataSlice {
	result := make(DBAchievementCriteriaDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAchievementCriteriaDataSlice) ToWeb() AchievementCriteriaDataSlice {
	result := make(AchievementCriteriaDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
