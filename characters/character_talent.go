package model

import "github.com/uptrace/bun"

type CharacterTalent struct {
	Guid int `json:"guid,omitempty"`
	Spell int32 `json:"spell,omitempty"`
	TalentGroup int8 `json:"talentgroup,omitempty"`
}

type CharacterTalentSlice []CharacterTalent

type DBCharacterTalent struct {
	bun.BaseModel `bun:"table:character_talent,alias:character_talent"`
	Guid int `bun:"guid"`
	Spell int32 `bun:"spell"`
	TalentGroup int8 `bun:"talentGroup"`
}

type DBCharacterTalentSlice []DBCharacterTalent

func (entry *CharacterTalent) ToDB() *DBCharacterTalent {
	if entry == nil {
		return nil
	}
	return &DBCharacterTalent{
		Guid: entry.Guid,
		Spell: entry.Spell,
		TalentGroup: entry.TalentGroup,
	}
}

func (entry *DBCharacterTalent) ToWeb() *CharacterTalent {
	if entry == nil {
		return nil
	}
	return &CharacterTalent{
		Guid: entry.Guid,
		Spell: entry.Spell,
		TalentGroup: entry.TalentGroup,
	}
}

func (data CharacterTalentSlice) ToDB() DBCharacterTalentSlice {
	result := make(DBCharacterTalentSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterTalentSlice) ToWeb() CharacterTalentSlice {
	result := make(CharacterTalentSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
