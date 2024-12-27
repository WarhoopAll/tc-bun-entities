package model

import "github.com/uptrace/bun"

type CharacterReputation struct {
	Guid int `json:"guid,omitempty"`
	Faction int16 `json:"faction,omitempty"`
	Standing int `json:"standing,omitempty"`
	Flags int16 `json:"flags,omitempty"`
}

type CharacterReputationSlice []CharacterReputation

type DBCharacterReputation struct {
	bun.BaseModel `bun:"table:character_reputation,alias:character_reputation"`
	Guid int `bun:"guid"`
	Faction int16 `bun:"faction"`
	Standing int `bun:"standing"`
	Flags int16 `bun:"flags"`
}

type DBCharacterReputationSlice []DBCharacterReputation

func (entry *CharacterReputation) ToDB() *DBCharacterReputation {
	if entry == nil {
		return nil
	}
	return &DBCharacterReputation{
		Guid: entry.Guid,
		Faction: entry.Faction,
		Standing: entry.Standing,
		Flags: entry.Flags,
	}
}

func (entry *DBCharacterReputation) ToWeb() *CharacterReputation {
	if entry == nil {
		return nil
	}
	return &CharacterReputation{
		Guid: entry.Guid,
		Faction: entry.Faction,
		Standing: entry.Standing,
		Flags: entry.Flags,
	}
}

func (data CharacterReputationSlice) ToDB() DBCharacterReputationSlice {
	result := make(DBCharacterReputationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterReputationSlice) ToWeb() CharacterReputationSlice {
	result := make(CharacterReputationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
