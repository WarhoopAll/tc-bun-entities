package model

import "github.com/uptrace/bun"

type AchievementReward struct {
	ID int `json:"id,omitempty"`
	TitleA int `json:"titlea,omitempty"`
	TitleH int `json:"titleh,omitempty"`
	ItemID int `json:"itemid,omitempty"`
	Sender int `json:"sender,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body string `json:"body,omitempty"`
	MailTemplateID int `json:"mailtemplateid,omitempty"`
}

type AchievementRewardSlice []AchievementReward

type DBAchievementReward struct {
	bun.BaseModel `bun:"table:achievement_reward,alias:achievement_reward"`
	ID int `bun:"ID"`
	TitleA int `bun:"TitleA"`
	TitleH int `bun:"TitleH"`
	ItemID int `bun:"ItemID"`
	Sender int `bun:"Sender"`
	Subject string `bun:"Subject"`
	Body string `bun:"Body"`
	MailTemplateID int `bun:"MailTemplateID"`
}

type DBAchievementRewardSlice []DBAchievementReward

func (entry *AchievementReward) ToDB() *DBAchievementReward {
	if entry == nil {
		return nil
	}
	return &DBAchievementReward{
		ID: entry.ID,
		TitleA: entry.TitleA,
		TitleH: entry.TitleH,
		ItemID: entry.ItemID,
		Sender: entry.Sender,
		Subject: entry.Subject,
		Body: entry.Body,
		MailTemplateID: entry.MailTemplateID,
	}
}

func (entry *DBAchievementReward) ToWeb() *AchievementReward {
	if entry == nil {
		return nil
	}
	return &AchievementReward{
		ID: entry.ID,
		TitleA: entry.TitleA,
		TitleH: entry.TitleH,
		ItemID: entry.ItemID,
		Sender: entry.Sender,
		Subject: entry.Subject,
		Body: entry.Body,
		MailTemplateID: entry.MailTemplateID,
	}
}

func (data AchievementRewardSlice) ToDB() DBAchievementRewardSlice {
	result := make(DBAchievementRewardSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAchievementRewardSlice) ToWeb() AchievementRewardSlice {
	result := make(AchievementRewardSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
