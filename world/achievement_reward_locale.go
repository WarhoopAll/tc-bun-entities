package model

import "github.com/uptrace/bun"

type AchievementRewardLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body string `json:"body,omitempty"`
}

type AchievementRewardLocaleSlice []AchievementRewardLocale

type DBAchievementRewardLocale struct {
	bun.BaseModel `bun:"table:achievement_reward_locale,alias:achievement_reward_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"Locale"`
	Subject string `bun:"Subject"`
	Body string `bun:"Body"`
}

type DBAchievementRewardLocaleSlice []DBAchievementRewardLocale

func (entry *AchievementRewardLocale) ToDB() *DBAchievementRewardLocale {
	if entry == nil {
		return nil
	}
	return &DBAchievementRewardLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Subject: entry.Subject,
		Body: entry.Body,
	}
}

func (entry *DBAchievementRewardLocale) ToWeb() *AchievementRewardLocale {
	if entry == nil {
		return nil
	}
	return &AchievementRewardLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Subject: entry.Subject,
		Body: entry.Body,
	}
}

func (data AchievementRewardLocaleSlice) ToDB() DBAchievementRewardLocaleSlice {
	result := make(DBAchievementRewardLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAchievementRewardLocaleSlice) ToWeb() AchievementRewardLocaleSlice {
	result := make(AchievementRewardLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
