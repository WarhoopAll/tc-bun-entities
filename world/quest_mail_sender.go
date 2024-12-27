package model

import "github.com/uptrace/bun"

type QuestMailSender struct {
	QuestId int `json:"questid,omitempty"`
	RewardMailSenderEntry int `json:"rewardmailsenderentry,omitempty"`
}

type QuestMailSenderSlice []QuestMailSender

type DBQuestMailSender struct {
	bun.BaseModel `bun:"table:quest_mail_sender,alias:quest_mail_sender"`
	QuestId int `bun:"QuestId"`
	RewardMailSenderEntry int `bun:"RewardMailSenderEntry"`
}

type DBQuestMailSenderSlice []DBQuestMailSender

func (entry *QuestMailSender) ToDB() *DBQuestMailSender {
	if entry == nil {
		return nil
	}
	return &DBQuestMailSender{
		QuestId: entry.QuestId,
		RewardMailSenderEntry: entry.RewardMailSenderEntry,
	}
}

func (entry *DBQuestMailSender) ToWeb() *QuestMailSender {
	if entry == nil {
		return nil
	}
	return &QuestMailSender{
		QuestId: entry.QuestId,
		RewardMailSenderEntry: entry.RewardMailSenderEntry,
	}
}

func (data QuestMailSenderSlice) ToDB() DBQuestMailSenderSlice {
	result := make(DBQuestMailSenderSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestMailSenderSlice) ToWeb() QuestMailSenderSlice {
	result := make(QuestMailSenderSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
