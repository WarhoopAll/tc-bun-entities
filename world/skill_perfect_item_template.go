package model

import "github.com/uptrace/bun"

type SkillPerfectItemTemplate struct {
	SpellId int `json:"spellid,omitempty"`
	RequiredSpecialization int `json:"requiredspecialization,omitempty"`
	PerfectCreateChance float64 `json:"perfectcreatechance,omitempty"`
	PerfectItemType int `json:"perfectitemtype,omitempty"`
}

type SkillPerfectItemTemplateSlice []SkillPerfectItemTemplate

type DBSkillPerfectItemTemplate struct {
	bun.BaseModel `bun:"table:skill_perfect_item_template,alias:skill_perfect_item_template"`
	SpellId int `bun:"spellId"`
	RequiredSpecialization int `bun:"requiredSpecialization"`
	PerfectCreateChance float64 `bun:"perfectCreateChance"`
	PerfectItemType int `bun:"perfectItemType"`
}

type DBSkillPerfectItemTemplateSlice []DBSkillPerfectItemTemplate

func (entry *SkillPerfectItemTemplate) ToDB() *DBSkillPerfectItemTemplate {
	if entry == nil {
		return nil
	}
	return &DBSkillPerfectItemTemplate{
		SpellId: entry.SpellId,
		RequiredSpecialization: entry.RequiredSpecialization,
		PerfectCreateChance: entry.PerfectCreateChance,
		PerfectItemType: entry.PerfectItemType,
	}
}

func (entry *DBSkillPerfectItemTemplate) ToWeb() *SkillPerfectItemTemplate {
	if entry == nil {
		return nil
	}
	return &SkillPerfectItemTemplate{
		SpellId: entry.SpellId,
		RequiredSpecialization: entry.RequiredSpecialization,
		PerfectCreateChance: entry.PerfectCreateChance,
		PerfectItemType: entry.PerfectItemType,
	}
}

func (data SkillPerfectItemTemplateSlice) ToDB() DBSkillPerfectItemTemplateSlice {
	result := make(DBSkillPerfectItemTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSkillPerfectItemTemplateSlice) ToWeb() SkillPerfectItemTemplateSlice {
	result := make(SkillPerfectItemTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
