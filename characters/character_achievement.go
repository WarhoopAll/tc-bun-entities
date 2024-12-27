package model

import "github.com/uptrace/bun"

type CharacterAchievement struct {
	Guid int `json:"guid,omitempty"`
	Achievement int16 `json:"achievement,omitempty"`
	Date int `json:"date,omitempty"`
}

type CharacterAchievementSlice []CharacterAchievement

type DBCharacterAchievement struct {
	bun.BaseModel `bun:"table:character_achievement,alias:character_achievement"`
	Guid int `bun:"guid"`
	Achievement int16 `bun:"achievement"`
	Date int `bun:"date"`
}

type DBCharacterAchievementSlice []DBCharacterAchievement

func (entry *CharacterAchievement) ToDB() *DBCharacterAchievement {
	if entry == nil {
		return nil
	}
	return &DBCharacterAchievement{
		Guid: entry.Guid,
		Achievement: entry.Achievement,
		Date: entry.Date,
	}
}

func (entry *DBCharacterAchievement) ToWeb() *CharacterAchievement {
	if entry == nil {
		return nil
	}
	return &CharacterAchievement{
		Guid: entry.Guid,
		Achievement: entry.Achievement,
		Date: entry.Date,
	}
}

func (data CharacterAchievementSlice) ToDB() DBCharacterAchievementSlice {
	result := make(DBCharacterAchievementSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterAchievementSlice) ToWeb() CharacterAchievementSlice {
	result := make(CharacterAchievementSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
