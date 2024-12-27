package model

import "github.com/uptrace/bun"

type CharacterAchievementProgress struct {
	Guid int `json:"guid,omitempty"`
	Criteria int16 `json:"criteria,omitempty"`
	Counter int `json:"counter,omitempty"`
	Date int `json:"date,omitempty"`
}

type CharacterAchievementProgressSlice []CharacterAchievementProgress

type DBCharacterAchievementProgress struct {
	bun.BaseModel `bun:"table:character_achievement_progress,alias:character_achievement_progress"`
	Guid int `bun:"guid"`
	Criteria int16 `bun:"criteria"`
	Counter int `bun:"counter"`
	Date int `bun:"date"`
}

type DBCharacterAchievementProgressSlice []DBCharacterAchievementProgress

func (entry *CharacterAchievementProgress) ToDB() *DBCharacterAchievementProgress {
	if entry == nil {
		return nil
	}
	return &DBCharacterAchievementProgress{
		Guid: entry.Guid,
		Criteria: entry.Criteria,
		Counter: entry.Counter,
		Date: entry.Date,
	}
}

func (entry *DBCharacterAchievementProgress) ToWeb() *CharacterAchievementProgress {
	if entry == nil {
		return nil
	}
	return &CharacterAchievementProgress{
		Guid: entry.Guid,
		Criteria: entry.Criteria,
		Counter: entry.Counter,
		Date: entry.Date,
	}
}

func (data CharacterAchievementProgressSlice) ToDB() DBCharacterAchievementProgressSlice {
	result := make(DBCharacterAchievementProgressSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterAchievementProgressSlice) ToWeb() CharacterAchievementProgressSlice {
	result := make(CharacterAchievementProgressSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
