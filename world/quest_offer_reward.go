package model

import "github.com/uptrace/bun"

type QuestOfferReward struct {
	ID int `json:"id,omitempty"`
	Emote1 int16 `json:"emote1,omitempty"`
	Emote2 int16 `json:"emote2,omitempty"`
	Emote3 int16 `json:"emote3,omitempty"`
	Emote4 int16 `json:"emote4,omitempty"`
	EmoteDelay1 int `json:"emotedelay1,omitempty"`
	EmoteDelay2 int `json:"emotedelay2,omitempty"`
	EmoteDelay3 int `json:"emotedelay3,omitempty"`
	EmoteDelay4 int `json:"emotedelay4,omitempty"`
	RewardText string `json:"rewardtext,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestOfferRewardSlice []QuestOfferReward

type DBQuestOfferReward struct {
	bun.BaseModel `bun:"table:quest_offer_reward,alias:quest_offer_reward"`
	ID int `bun:"ID"`
	Emote1 int16 `bun:"Emote1"`
	Emote2 int16 `bun:"Emote2"`
	Emote3 int16 `bun:"Emote3"`
	Emote4 int16 `bun:"Emote4"`
	EmoteDelay1 int `bun:"EmoteDelay1"`
	EmoteDelay2 int `bun:"EmoteDelay2"`
	EmoteDelay3 int `bun:"EmoteDelay3"`
	EmoteDelay4 int `bun:"EmoteDelay4"`
	RewardText string `bun:"RewardText"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestOfferRewardSlice []DBQuestOfferReward

func (entry *QuestOfferReward) ToDB() *DBQuestOfferReward {
	if entry == nil {
		return nil
	}
	return &DBQuestOfferReward{
		ID: entry.ID,
		Emote1: entry.Emote1,
		Emote2: entry.Emote2,
		Emote3: entry.Emote3,
		Emote4: entry.Emote4,
		EmoteDelay1: entry.EmoteDelay1,
		EmoteDelay2: entry.EmoteDelay2,
		EmoteDelay3: entry.EmoteDelay3,
		EmoteDelay4: entry.EmoteDelay4,
		RewardText: entry.RewardText,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestOfferReward) ToWeb() *QuestOfferReward {
	if entry == nil {
		return nil
	}
	return &QuestOfferReward{
		ID: entry.ID,
		Emote1: entry.Emote1,
		Emote2: entry.Emote2,
		Emote3: entry.Emote3,
		Emote4: entry.Emote4,
		EmoteDelay1: entry.EmoteDelay1,
		EmoteDelay2: entry.EmoteDelay2,
		EmoteDelay3: entry.EmoteDelay3,
		EmoteDelay4: entry.EmoteDelay4,
		RewardText: entry.RewardText,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestOfferRewardSlice) ToDB() DBQuestOfferRewardSlice {
	result := make(DBQuestOfferRewardSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestOfferRewardSlice) ToWeb() QuestOfferRewardSlice {
	result := make(QuestOfferRewardSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
