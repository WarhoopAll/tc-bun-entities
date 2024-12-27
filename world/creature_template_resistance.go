package model

import "github.com/uptrace/bun"

type CreatureTemplateResistance struct {
	CreatureID int `json:"creatureid,omitempty"`
	School int8 `json:"school,omitempty"`
	Resistance int16 `json:"resistance,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type CreatureTemplateResistanceSlice []CreatureTemplateResistance

type DBCreatureTemplateResistance struct {
	bun.BaseModel `bun:"table:creature_template_resistance,alias:creature_template_resistance"`
	CreatureID int `bun:"CreatureID"`
	School int8 `bun:"School"`
	Resistance int16 `bun:"Resistance"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBCreatureTemplateResistanceSlice []DBCreatureTemplateResistance

func (entry *CreatureTemplateResistance) ToDB() *DBCreatureTemplateResistance {
	if entry == nil {
		return nil
	}
	return &DBCreatureTemplateResistance{
		CreatureID: entry.CreatureID,
		School: entry.School,
		Resistance: entry.Resistance,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBCreatureTemplateResistance) ToWeb() *CreatureTemplateResistance {
	if entry == nil {
		return nil
	}
	return &CreatureTemplateResistance{
		CreatureID: entry.CreatureID,
		School: entry.School,
		Resistance: entry.Resistance,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data CreatureTemplateResistanceSlice) ToDB() DBCreatureTemplateResistanceSlice {
	result := make(DBCreatureTemplateResistanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureTemplateResistanceSlice) ToWeb() CreatureTemplateResistanceSlice {
	result := make(CreatureTemplateResistanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
