package model

import "github.com/uptrace/bun"

type MailLevelReward struct {
	Level int8 `json:"level,omitempty"`
	RaceMask int `json:"racemask,omitempty"`
	MailTemplateId int `json:"mailtemplateid,omitempty"`
	SenderEntry int `json:"senderentry,omitempty"`
}

type MailLevelRewardSlice []MailLevelReward

type DBMailLevelReward struct {
	bun.BaseModel `bun:"table:mail_level_reward,alias:mail_level_reward"`
	Level int8 `bun:"level"`
	RaceMask int `bun:"raceMask"`
	MailTemplateId int `bun:"mailTemplateId"`
	SenderEntry int `bun:"senderEntry"`
}

type DBMailLevelRewardSlice []DBMailLevelReward

func (entry *MailLevelReward) ToDB() *DBMailLevelReward {
	if entry == nil {
		return nil
	}
	return &DBMailLevelReward{
		Level: entry.Level,
		RaceMask: entry.RaceMask,
		MailTemplateId: entry.MailTemplateId,
		SenderEntry: entry.SenderEntry,
	}
}

func (entry *DBMailLevelReward) ToWeb() *MailLevelReward {
	if entry == nil {
		return nil
	}
	return &MailLevelReward{
		Level: entry.Level,
		RaceMask: entry.RaceMask,
		MailTemplateId: entry.MailTemplateId,
		SenderEntry: entry.SenderEntry,
	}
}

func (data MailLevelRewardSlice) ToDB() DBMailLevelRewardSlice {
	result := make(DBMailLevelRewardSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBMailLevelRewardSlice) ToWeb() MailLevelRewardSlice {
	result := make(MailLevelRewardSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
