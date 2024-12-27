package model

import "github.com/uptrace/bun"

type SkillFishingBaseLevel struct {
	Entry int `json:"entry,omitempty"`
	Skill int16 `json:"skill,omitempty"`
}

type SkillFishingBaseLevelSlice []SkillFishingBaseLevel

type DBSkillFishingBaseLevel struct {
	bun.BaseModel `bun:"table:skill_fishing_base_level,alias:skill_fishing_base_level"`
	Entry int `bun:"entry"`
	Skill int16 `bun:"skill"`
}

type DBSkillFishingBaseLevelSlice []DBSkillFishingBaseLevel

func (entry *SkillFishingBaseLevel) ToDB() *DBSkillFishingBaseLevel {
	if entry == nil {
		return nil
	}
	return &DBSkillFishingBaseLevel{
		Entry: entry.Entry,
		Skill: entry.Skill,
	}
}

func (entry *DBSkillFishingBaseLevel) ToWeb() *SkillFishingBaseLevel {
	if entry == nil {
		return nil
	}
	return &SkillFishingBaseLevel{
		Entry: entry.Entry,
		Skill: entry.Skill,
	}
}

func (data SkillFishingBaseLevelSlice) ToDB() DBSkillFishingBaseLevelSlice {
	result := make(DBSkillFishingBaseLevelSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSkillFishingBaseLevelSlice) ToWeb() SkillFishingBaseLevelSlice {
	result := make(SkillFishingBaseLevelSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
