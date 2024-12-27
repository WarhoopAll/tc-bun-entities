package model

import "github.com/uptrace/bun"

type SkillExtraItemTemplate struct {
	SpellId int `json:"spellid,omitempty"`
	RequiredSpecialization int `json:"requiredspecialization,omitempty"`
	AdditionalCreateChance float64 `json:"additionalcreatechance,omitempty"`
	AdditionalMaxNum int8 `json:"additionalmaxnum,omitempty"`
}

type SkillExtraItemTemplateSlice []SkillExtraItemTemplate

type DBSkillExtraItemTemplate struct {
	bun.BaseModel `bun:"table:skill_extra_item_template,alias:skill_extra_item_template"`
	SpellId int `bun:"spellId"`
	RequiredSpecialization int `bun:"requiredSpecialization"`
	AdditionalCreateChance float64 `bun:"additionalCreateChance"`
	AdditionalMaxNum int8 `bun:"additionalMaxNum"`
}

type DBSkillExtraItemTemplateSlice []DBSkillExtraItemTemplate

func (entry *SkillExtraItemTemplate) ToDB() *DBSkillExtraItemTemplate {
	if entry == nil {
		return nil
	}
	return &DBSkillExtraItemTemplate{
		SpellId: entry.SpellId,
		RequiredSpecialization: entry.RequiredSpecialization,
		AdditionalCreateChance: entry.AdditionalCreateChance,
		AdditionalMaxNum: entry.AdditionalMaxNum,
	}
}

func (entry *DBSkillExtraItemTemplate) ToWeb() *SkillExtraItemTemplate {
	if entry == nil {
		return nil
	}
	return &SkillExtraItemTemplate{
		SpellId: entry.SpellId,
		RequiredSpecialization: entry.RequiredSpecialization,
		AdditionalCreateChance: entry.AdditionalCreateChance,
		AdditionalMaxNum: entry.AdditionalMaxNum,
	}
}

func (data SkillExtraItemTemplateSlice) ToDB() DBSkillExtraItemTemplateSlice {
	result := make(DBSkillExtraItemTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSkillExtraItemTemplateSlice) ToWeb() SkillExtraItemTemplateSlice {
	result := make(SkillExtraItemTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
