package model

import "github.com/uptrace/bun"

type CreatureClasslevelstats struct {
	Level int8 `json:"level,omitempty"`
	Class int8 `json:"class,omitempty"`
	Basehp0 int16 `json:"basehp0,omitempty"`
	Basehp1 int16 `json:"basehp1,omitempty"`
	Basehp2 int16 `json:"basehp2,omitempty"`
	Basemana int16 `json:"basemana,omitempty"`
	Basearmor int16 `json:"basearmor,omitempty"`
	Attackpower int16 `json:"attackpower,omitempty"`
	Rangedattackpower int16 `json:"rangedattackpower,omitempty"`
	DamageBase float64 `json:"damage_base,omitempty"`
	DamageExp1 float64 `json:"damage_exp1,omitempty"`
	DamageExp2 float64 `json:"damage_exp2,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type CreatureClasslevelstatsSlice []CreatureClasslevelstats

type DBCreatureClasslevelstats struct {
	bun.BaseModel `bun:"table:creature_classlevelstats,alias:creature_classlevelstats"`
	Level int8 `bun:"level"`
	Class int8 `bun:"class"`
	Basehp0 int16 `bun:"basehp0"`
	Basehp1 int16 `bun:"basehp1"`
	Basehp2 int16 `bun:"basehp2"`
	Basemana int16 `bun:"basemana"`
	Basearmor int16 `bun:"basearmor"`
	Attackpower int16 `bun:"attackpower"`
	Rangedattackpower int16 `bun:"rangedattackpower"`
	DamageBase float64 `bun:"damage_base"`
	DamageExp1 float64 `bun:"damage_exp1"`
	DamageExp2 float64 `bun:"damage_exp2"`
	Comment string `bun:"comment"`
}

type DBCreatureClasslevelstatsSlice []DBCreatureClasslevelstats

func (entry *CreatureClasslevelstats) ToDB() *DBCreatureClasslevelstats {
	if entry == nil {
		return nil
	}
	return &DBCreatureClasslevelstats{
		Level: entry.Level,
		Class: entry.Class,
		Basehp0: entry.Basehp0,
		Basehp1: entry.Basehp1,
		Basehp2: entry.Basehp2,
		Basemana: entry.Basemana,
		Basearmor: entry.Basearmor,
		Attackpower: entry.Attackpower,
		Rangedattackpower: entry.Rangedattackpower,
		DamageBase: entry.DamageBase,
		DamageExp1: entry.DamageExp1,
		DamageExp2: entry.DamageExp2,
		Comment: entry.Comment,
	}
}

func (entry *DBCreatureClasslevelstats) ToWeb() *CreatureClasslevelstats {
	if entry == nil {
		return nil
	}
	return &CreatureClasslevelstats{
		Level: entry.Level,
		Class: entry.Class,
		Basehp0: entry.Basehp0,
		Basehp1: entry.Basehp1,
		Basehp2: entry.Basehp2,
		Basemana: entry.Basemana,
		Basearmor: entry.Basearmor,
		Attackpower: entry.Attackpower,
		Rangedattackpower: entry.Rangedattackpower,
		DamageBase: entry.DamageBase,
		DamageExp1: entry.DamageExp1,
		DamageExp2: entry.DamageExp2,
		Comment: entry.Comment,
	}
}

func (data CreatureClasslevelstatsSlice) ToDB() DBCreatureClasslevelstatsSlice {
	result := make(DBCreatureClasslevelstatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureClasslevelstatsSlice) ToWeb() CreatureClasslevelstatsSlice {
	result := make(CreatureClasslevelstatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
