package model

import "github.com/uptrace/bun"

type BroadcastTextLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Text string `json:"text,omitempty"`
	Text1 string `json:"text1,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type BroadcastTextLocaleSlice []BroadcastTextLocale

type DBBroadcastTextLocale struct {
	bun.BaseModel `bun:"table:broadcast_text_locale,alias:broadcast_text_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"locale"`
	Text string `bun:"Text"`
	Text1 string `bun:"Text1"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBBroadcastTextLocaleSlice []DBBroadcastTextLocale

func (entry *BroadcastTextLocale) ToDB() *DBBroadcastTextLocale {
	if entry == nil {
		return nil
	}
	return &DBBroadcastTextLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Text: entry.Text,
		Text1: entry.Text1,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBBroadcastTextLocale) ToWeb() *BroadcastTextLocale {
	if entry == nil {
		return nil
	}
	return &BroadcastTextLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Text: entry.Text,
		Text1: entry.Text1,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data BroadcastTextLocaleSlice) ToDB() DBBroadcastTextLocaleSlice {
	result := make(DBBroadcastTextLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBroadcastTextLocaleSlice) ToWeb() BroadcastTextLocaleSlice {
	result := make(BroadcastTextLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
