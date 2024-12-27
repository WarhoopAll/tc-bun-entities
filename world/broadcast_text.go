package model

import "github.com/uptrace/bun"

type BroadcastText struct {
	ID int `json:"id,omitempty"`
	LanguageID int `json:"languageid,omitempty"`
	Text string `json:"text,omitempty"`
	Text1 string `json:"text1,omitempty"`
	EmoteID1 int `json:"emoteid1,omitempty"`
	EmoteID2 int `json:"emoteid2,omitempty"`
	EmoteID3 int `json:"emoteid3,omitempty"`
	EmoteDelay1 int `json:"emotedelay1,omitempty"`
	EmoteDelay2 int `json:"emotedelay2,omitempty"`
	EmoteDelay3 int `json:"emotedelay3,omitempty"`
	SoundEntriesID int `json:"soundentriesid,omitempty"`
	EmotesID int `json:"emotesid,omitempty"`
	Flags int `json:"flags,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type BroadcastTextSlice []BroadcastText

type DBBroadcastText struct {
	bun.BaseModel `bun:"table:broadcast_text,alias:broadcast_text"`
	ID int `bun:"ID"`
	LanguageID int `bun:"LanguageID"`
	Text string `bun:"Text"`
	Text1 string `bun:"Text1"`
	EmoteID1 int `bun:"EmoteID1"`
	EmoteID2 int `bun:"EmoteID2"`
	EmoteID3 int `bun:"EmoteID3"`
	EmoteDelay1 int `bun:"EmoteDelay1"`
	EmoteDelay2 int `bun:"EmoteDelay2"`
	EmoteDelay3 int `bun:"EmoteDelay3"`
	SoundEntriesID int `bun:"SoundEntriesID"`
	EmotesID int `bun:"EmotesID"`
	Flags int `bun:"Flags"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBBroadcastTextSlice []DBBroadcastText

func (entry *BroadcastText) ToDB() *DBBroadcastText {
	if entry == nil {
		return nil
	}
	return &DBBroadcastText{
		ID: entry.ID,
		LanguageID: entry.LanguageID,
		Text: entry.Text,
		Text1: entry.Text1,
		EmoteID1: entry.EmoteID1,
		EmoteID2: entry.EmoteID2,
		EmoteID3: entry.EmoteID3,
		EmoteDelay1: entry.EmoteDelay1,
		EmoteDelay2: entry.EmoteDelay2,
		EmoteDelay3: entry.EmoteDelay3,
		SoundEntriesID: entry.SoundEntriesID,
		EmotesID: entry.EmotesID,
		Flags: entry.Flags,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBBroadcastText) ToWeb() *BroadcastText {
	if entry == nil {
		return nil
	}
	return &BroadcastText{
		ID: entry.ID,
		LanguageID: entry.LanguageID,
		Text: entry.Text,
		Text1: entry.Text1,
		EmoteID1: entry.EmoteID1,
		EmoteID2: entry.EmoteID2,
		EmoteID3: entry.EmoteID3,
		EmoteDelay1: entry.EmoteDelay1,
		EmoteDelay2: entry.EmoteDelay2,
		EmoteDelay3: entry.EmoteDelay3,
		SoundEntriesID: entry.SoundEntriesID,
		EmotesID: entry.EmotesID,
		Flags: entry.Flags,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data BroadcastTextSlice) ToDB() DBBroadcastTextSlice {
	result := make(DBBroadcastTextSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBroadcastTextSlice) ToWeb() BroadcastTextSlice {
	result := make(BroadcastTextSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
