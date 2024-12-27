package model

import "github.com/uptrace/bun"

type CharacterBattlegroundRandom struct {
	Guid int `json:"guid,omitempty"`
}

type CharacterBattlegroundRandomSlice []CharacterBattlegroundRandom

type DBCharacterBattlegroundRandom struct {
	bun.BaseModel `bun:"table:character_battleground_random,alias:character_battleground_random"`
	Guid int `bun:"guid"`
}

type DBCharacterBattlegroundRandomSlice []DBCharacterBattlegroundRandom

func (entry *CharacterBattlegroundRandom) ToDB() *DBCharacterBattlegroundRandom {
	if entry == nil {
		return nil
	}
	return &DBCharacterBattlegroundRandom{
		Guid: entry.Guid,
	}
}

func (entry *DBCharacterBattlegroundRandom) ToWeb() *CharacterBattlegroundRandom {
	if entry == nil {
		return nil
	}
	return &CharacterBattlegroundRandom{
		Guid: entry.Guid,
	}
}

func (data CharacterBattlegroundRandomSlice) ToDB() DBCharacterBattlegroundRandomSlice {
	result := make(DBCharacterBattlegroundRandomSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterBattlegroundRandomSlice) ToWeb() CharacterBattlegroundRandomSlice {
	result := make(CharacterBattlegroundRandomSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
