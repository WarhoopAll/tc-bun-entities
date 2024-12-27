package model

import "github.com/uptrace/bun"

type PageTextLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Text string `json:"text,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type PageTextLocaleSlice []PageTextLocale

type DBPageTextLocale struct {
	bun.BaseModel `bun:"table:page_text_locale,alias:page_text_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"locale"`
	Text string `bun:"Text"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBPageTextLocaleSlice []DBPageTextLocale

func (entry *PageTextLocale) ToDB() *DBPageTextLocale {
	if entry == nil {
		return nil
	}
	return &DBPageTextLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Text: entry.Text,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBPageTextLocale) ToWeb() *PageTextLocale {
	if entry == nil {
		return nil
	}
	return &PageTextLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Text: entry.Text,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data PageTextLocaleSlice) ToDB() DBPageTextLocaleSlice {
	result := make(DBPageTextLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPageTextLocaleSlice) ToWeb() PageTextLocaleSlice {
	result := make(PageTextLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
