package model

import "github.com/uptrace/bun"

type QuestDetails struct {
	ID int `json:"id,omitempty"`
	Emote1 int16 `json:"emote1,omitempty"`
	Emote2 int16 `json:"emote2,omitempty"`
	Emote3 int16 `json:"emote3,omitempty"`
	Emote4 int16 `json:"emote4,omitempty"`
	EmoteDelay1 int `json:"emotedelay1,omitempty"`
	EmoteDelay2 int `json:"emotedelay2,omitempty"`
	EmoteDelay3 int `json:"emotedelay3,omitempty"`
	EmoteDelay4 int `json:"emotedelay4,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestDetailsSlice []QuestDetails

type DBQuestDetails struct {
	bun.BaseModel `bun:"table:quest_details,alias:quest_details"`
	ID int `bun:"ID"`
	Emote1 int16 `bun:"Emote1"`
	Emote2 int16 `bun:"Emote2"`
	Emote3 int16 `bun:"Emote3"`
	Emote4 int16 `bun:"Emote4"`
	EmoteDelay1 int `bun:"EmoteDelay1"`
	EmoteDelay2 int `bun:"EmoteDelay2"`
	EmoteDelay3 int `bun:"EmoteDelay3"`
	EmoteDelay4 int `bun:"EmoteDelay4"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestDetailsSlice []DBQuestDetails

func (entry *QuestDetails) ToDB() *DBQuestDetails {
	if entry == nil {
		return nil
	}
	return &DBQuestDetails{
		ID: entry.ID,
		Emote1: entry.Emote1,
		Emote2: entry.Emote2,
		Emote3: entry.Emote3,
		Emote4: entry.Emote4,
		EmoteDelay1: entry.EmoteDelay1,
		EmoteDelay2: entry.EmoteDelay2,
		EmoteDelay3: entry.EmoteDelay3,
		EmoteDelay4: entry.EmoteDelay4,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestDetails) ToWeb() *QuestDetails {
	if entry == nil {
		return nil
	}
	return &QuestDetails{
		ID: entry.ID,
		Emote1: entry.Emote1,
		Emote2: entry.Emote2,
		Emote3: entry.Emote3,
		Emote4: entry.Emote4,
		EmoteDelay1: entry.EmoteDelay1,
		EmoteDelay2: entry.EmoteDelay2,
		EmoteDelay3: entry.EmoteDelay3,
		EmoteDelay4: entry.EmoteDelay4,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestDetailsSlice) ToDB() DBQuestDetailsSlice {
	result := make(DBQuestDetailsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestDetailsSlice) ToWeb() QuestDetailsSlice {
	result := make(QuestDetailsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
