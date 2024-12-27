package model

import "github.com/uptrace/bun"

type QuestOfferRewardLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	RewardText string `json:"rewardtext,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestOfferRewardLocaleSlice []QuestOfferRewardLocale

type DBQuestOfferRewardLocale struct {
	bun.BaseModel `bun:"table:quest_offer_reward_locale,alias:quest_offer_reward_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"locale"`
	RewardText string `bun:"RewardText"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestOfferRewardLocaleSlice []DBQuestOfferRewardLocale

func (entry *QuestOfferRewardLocale) ToDB() *DBQuestOfferRewardLocale {
	if entry == nil {
		return nil
	}
	return &DBQuestOfferRewardLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		RewardText: entry.RewardText,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestOfferRewardLocale) ToWeb() *QuestOfferRewardLocale {
	if entry == nil {
		return nil
	}
	return &QuestOfferRewardLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		RewardText: entry.RewardText,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestOfferRewardLocaleSlice) ToDB() DBQuestOfferRewardLocaleSlice {
	result := make(DBQuestOfferRewardLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestOfferRewardLocaleSlice) ToWeb() QuestOfferRewardLocaleSlice {
	result := make(QuestOfferRewardLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
