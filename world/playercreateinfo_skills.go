package model

import "github.com/uptrace/bun"

type PlayercreateinfoSkills struct {
	RaceMask int `json:"racemask,omitempty"`
	ClassMask int `json:"classmask,omitempty"`
	Skill int16 `json:"skill,omitempty"`
	Rank int16 `json:"rank,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type PlayercreateinfoSkillsSlice []PlayercreateinfoSkills

type DBPlayercreateinfoSkills struct {
	bun.BaseModel `bun:"table:playercreateinfo_skills,alias:playercreateinfo_skills"`
	RaceMask int `bun:"raceMask"`
	ClassMask int `bun:"classMask"`
	Skill int16 `bun:"skill"`
	Rank int16 `bun:"rank"`
	Comment string `bun:"comment"`
}

type DBPlayercreateinfoSkillsSlice []DBPlayercreateinfoSkills

func (entry *PlayercreateinfoSkills) ToDB() *DBPlayercreateinfoSkills {
	if entry == nil {
		return nil
	}
	return &DBPlayercreateinfoSkills{
		RaceMask: entry.RaceMask,
		ClassMask: entry.ClassMask,
		Skill: entry.Skill,
		Rank: entry.Rank,
		Comment: entry.Comment,
	}
}

func (entry *DBPlayercreateinfoSkills) ToWeb() *PlayercreateinfoSkills {
	if entry == nil {
		return nil
	}
	return &PlayercreateinfoSkills{
		RaceMask: entry.RaceMask,
		ClassMask: entry.ClassMask,
		Skill: entry.Skill,
		Rank: entry.Rank,
		Comment: entry.Comment,
	}
}

func (data PlayercreateinfoSkillsSlice) ToDB() DBPlayercreateinfoSkillsSlice {
	result := make(DBPlayercreateinfoSkillsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayercreateinfoSkillsSlice) ToWeb() PlayercreateinfoSkillsSlice {
	result := make(PlayercreateinfoSkillsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
