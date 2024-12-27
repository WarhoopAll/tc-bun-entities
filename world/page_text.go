package model

import "github.com/uptrace/bun"

type PageText struct {
	ID int `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	NextPageID int `json:"nextpageid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type PageTextSlice []PageText

type DBPageText struct {
	bun.BaseModel `bun:"table:page_text,alias:page_text"`
	ID int `bun:"ID"`
	Text string `bun:"Text"`
	NextPageID int `bun:"NextPageID"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBPageTextSlice []DBPageText

func (entry *PageText) ToDB() *DBPageText {
	if entry == nil {
		return nil
	}
	return &DBPageText{
		ID: entry.ID,
		Text: entry.Text,
		NextPageID: entry.NextPageID,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBPageText) ToWeb() *PageText {
	if entry == nil {
		return nil
	}
	return &PageText{
		ID: entry.ID,
		Text: entry.Text,
		NextPageID: entry.NextPageID,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data PageTextSlice) ToDB() DBPageTextSlice {
	result := make(DBPageTextSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPageTextSlice) ToWeb() PageTextSlice {
	result := make(PageTextSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
