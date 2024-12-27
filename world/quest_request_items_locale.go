package model

import "github.com/uptrace/bun"

type QuestRequestItemsLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	CompletionText string `json:"completiontext,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestRequestItemsLocaleSlice []QuestRequestItemsLocale

type DBQuestRequestItemsLocale struct {
	bun.BaseModel `bun:"table:quest_request_items_locale,alias:quest_request_items_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"locale"`
	CompletionText string `bun:"CompletionText"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestRequestItemsLocaleSlice []DBQuestRequestItemsLocale

func (entry *QuestRequestItemsLocale) ToDB() *DBQuestRequestItemsLocale {
	if entry == nil {
		return nil
	}
	return &DBQuestRequestItemsLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		CompletionText: entry.CompletionText,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestRequestItemsLocale) ToWeb() *QuestRequestItemsLocale {
	if entry == nil {
		return nil
	}
	return &QuestRequestItemsLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		CompletionText: entry.CompletionText,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestRequestItemsLocaleSlice) ToDB() DBQuestRequestItemsLocaleSlice {
	result := make(DBQuestRequestItemsLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestRequestItemsLocaleSlice) ToWeb() QuestRequestItemsLocaleSlice {
	result := make(QuestRequestItemsLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
