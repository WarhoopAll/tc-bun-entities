package model

import "github.com/uptrace/bun"

type QuestRequestItems struct {
	ID int `json:"id,omitempty"`
	EmoteOnComplete int16 `json:"emoteoncomplete,omitempty"`
	EmoteOnIncomplete int16 `json:"emoteonincomplete,omitempty"`
	CompletionText string `json:"completiontext,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestRequestItemsSlice []QuestRequestItems

type DBQuestRequestItems struct {
	bun.BaseModel `bun:"table:quest_request_items,alias:quest_request_items"`
	ID int `bun:"ID"`
	EmoteOnComplete int16 `bun:"EmoteOnComplete"`
	EmoteOnIncomplete int16 `bun:"EmoteOnIncomplete"`
	CompletionText string `bun:"CompletionText"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestRequestItemsSlice []DBQuestRequestItems

func (entry *QuestRequestItems) ToDB() *DBQuestRequestItems {
	if entry == nil {
		return nil
	}
	return &DBQuestRequestItems{
		ID: entry.ID,
		EmoteOnComplete: entry.EmoteOnComplete,
		EmoteOnIncomplete: entry.EmoteOnIncomplete,
		CompletionText: entry.CompletionText,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestRequestItems) ToWeb() *QuestRequestItems {
	if entry == nil {
		return nil
	}
	return &QuestRequestItems{
		ID: entry.ID,
		EmoteOnComplete: entry.EmoteOnComplete,
		EmoteOnIncomplete: entry.EmoteOnIncomplete,
		CompletionText: entry.CompletionText,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestRequestItemsSlice) ToDB() DBQuestRequestItemsSlice {
	result := make(DBQuestRequestItemsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestRequestItemsSlice) ToWeb() QuestRequestItemsSlice {
	result := make(QuestRequestItemsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
