package model

import "github.com/uptrace/bun"

type CharacterFishingsteps struct {
	Guid int `json:"guid,omitempty"`
	FishingSteps int8 `json:"fishingsteps,omitempty"`
}

type CharacterFishingstepsSlice []CharacterFishingsteps

type DBCharacterFishingsteps struct {
	bun.BaseModel `bun:"table:character_fishingsteps,alias:character_fishingsteps"`
	Guid int `bun:"guid"`
	FishingSteps int8 `bun:"fishingSteps"`
}

type DBCharacterFishingstepsSlice []DBCharacterFishingsteps

func (entry *CharacterFishingsteps) ToDB() *DBCharacterFishingsteps {
	if entry == nil {
		return nil
	}
	return &DBCharacterFishingsteps{
		Guid: entry.Guid,
		FishingSteps: entry.FishingSteps,
	}
}

func (entry *DBCharacterFishingsteps) ToWeb() *CharacterFishingsteps {
	if entry == nil {
		return nil
	}
	return &CharacterFishingsteps{
		Guid: entry.Guid,
		FishingSteps: entry.FishingSteps,
	}
}

func (data CharacterFishingstepsSlice) ToDB() DBCharacterFishingstepsSlice {
	result := make(DBCharacterFishingstepsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterFishingstepsSlice) ToWeb() CharacterFishingstepsSlice {
	result := make(CharacterFishingstepsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
