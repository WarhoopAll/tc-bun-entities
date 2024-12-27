package model

import "github.com/uptrace/bun"

type TrainerSpell struct {
	TrainerId int `json:"trainerid,omitempty"`
	SpellId int `json:"spellid,omitempty"`
	MoneyCost int `json:"moneycost,omitempty"`
	ReqSkillLine int `json:"reqskillline,omitempty"`
	ReqSkillRank int `json:"reqskillrank,omitempty"`
	ReqAbility1 int `json:"reqability1,omitempty"`
	ReqAbility2 int `json:"reqability2,omitempty"`
	ReqAbility3 int `json:"reqability3,omitempty"`
	ReqLevel int8 `json:"reqlevel,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type TrainerSpellSlice []TrainerSpell

type DBTrainerSpell struct {
	bun.BaseModel `bun:"table:trainer_spell,alias:trainer_spell"`
	TrainerId int `bun:"TrainerId"`
	SpellId int `bun:"SpellId"`
	MoneyCost int `bun:"MoneyCost"`
	ReqSkillLine int `bun:"ReqSkillLine"`
	ReqSkillRank int `bun:"ReqSkillRank"`
	ReqAbility1 int `bun:"ReqAbility1"`
	ReqAbility2 int `bun:"ReqAbility2"`
	ReqAbility3 int `bun:"ReqAbility3"`
	ReqLevel int8 `bun:"ReqLevel"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBTrainerSpellSlice []DBTrainerSpell

func (entry *TrainerSpell) ToDB() *DBTrainerSpell {
	if entry == nil {
		return nil
	}
	return &DBTrainerSpell{
		TrainerId: entry.TrainerId,
		SpellId: entry.SpellId,
		MoneyCost: entry.MoneyCost,
		ReqSkillLine: entry.ReqSkillLine,
		ReqSkillRank: entry.ReqSkillRank,
		ReqAbility1: entry.ReqAbility1,
		ReqAbility2: entry.ReqAbility2,
		ReqAbility3: entry.ReqAbility3,
		ReqLevel: entry.ReqLevel,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBTrainerSpell) ToWeb() *TrainerSpell {
	if entry == nil {
		return nil
	}
	return &TrainerSpell{
		TrainerId: entry.TrainerId,
		SpellId: entry.SpellId,
		MoneyCost: entry.MoneyCost,
		ReqSkillLine: entry.ReqSkillLine,
		ReqSkillRank: entry.ReqSkillRank,
		ReqAbility1: entry.ReqAbility1,
		ReqAbility2: entry.ReqAbility2,
		ReqAbility3: entry.ReqAbility3,
		ReqLevel: entry.ReqLevel,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data TrainerSpellSlice) ToDB() DBTrainerSpellSlice {
	result := make(DBTrainerSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBTrainerSpellSlice) ToWeb() TrainerSpellSlice {
	result := make(TrainerSpellSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
