package model

import "github.com/uptrace/bun"

type CharacterSkills struct {
	Guid int `json:"guid,omitempty"`
	Skill int16 `json:"skill,omitempty"`
	Value int16 `json:"value,omitempty"`
	Max int16 `json:"max,omitempty"`
}

type CharacterSkillsSlice []CharacterSkills

type DBCharacterSkills struct {
	bun.BaseModel `bun:"table:character_skills,alias:character_skills"`
	Guid int `bun:"guid"`
	Skill int16 `bun:"skill"`
	Value int16 `bun:"value"`
	Max int16 `bun:"max"`
}

type DBCharacterSkillsSlice []DBCharacterSkills

func (entry *CharacterSkills) ToDB() *DBCharacterSkills {
	if entry == nil {
		return nil
	}
	return &DBCharacterSkills{
		Guid: entry.Guid,
		Skill: entry.Skill,
		Value: entry.Value,
		Max: entry.Max,
	}
}

func (entry *DBCharacterSkills) ToWeb() *CharacterSkills {
	if entry == nil {
		return nil
	}
	return &CharacterSkills{
		Guid: entry.Guid,
		Skill: entry.Skill,
		Value: entry.Value,
		Max: entry.Max,
	}
}

func (data CharacterSkillsSlice) ToDB() DBCharacterSkillsSlice {
	result := make(DBCharacterSkillsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterSkillsSlice) ToWeb() CharacterSkillsSlice {
	result := make(CharacterSkillsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
