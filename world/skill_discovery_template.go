package model

import "github.com/uptrace/bun"

type SkillDiscoveryTemplate struct {
	SpellId int `json:"spellid,omitempty"`
	ReqSpell int `json:"reqspell,omitempty"`
	ReqSkillValue int16 `json:"reqskillvalue,omitempty"`
	Chance float64 `json:"chance,omitempty"`
}

type SkillDiscoveryTemplateSlice []SkillDiscoveryTemplate

type DBSkillDiscoveryTemplate struct {
	bun.BaseModel `bun:"table:skill_discovery_template,alias:skill_discovery_template"`
	SpellId int `bun:"spellId"`
	ReqSpell int `bun:"reqSpell"`
	ReqSkillValue int16 `bun:"reqSkillValue"`
	Chance float64 `bun:"chance"`
}

type DBSkillDiscoveryTemplateSlice []DBSkillDiscoveryTemplate

func (entry *SkillDiscoveryTemplate) ToDB() *DBSkillDiscoveryTemplate {
	if entry == nil {
		return nil
	}
	return &DBSkillDiscoveryTemplate{
		SpellId: entry.SpellId,
		ReqSpell: entry.ReqSpell,
		ReqSkillValue: entry.ReqSkillValue,
		Chance: entry.Chance,
	}
}

func (entry *DBSkillDiscoveryTemplate) ToWeb() *SkillDiscoveryTemplate {
	if entry == nil {
		return nil
	}
	return &SkillDiscoveryTemplate{
		SpellId: entry.SpellId,
		ReqSpell: entry.ReqSpell,
		ReqSkillValue: entry.ReqSkillValue,
		Chance: entry.Chance,
	}
}

func (data SkillDiscoveryTemplateSlice) ToDB() DBSkillDiscoveryTemplateSlice {
	result := make(DBSkillDiscoveryTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSkillDiscoveryTemplateSlice) ToWeb() SkillDiscoveryTemplateSlice {
	result := make(SkillDiscoveryTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
