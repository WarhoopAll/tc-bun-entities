package model

import "github.com/uptrace/bun"

type NpcTextLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Text00 string `json:"text0_0,omitempty"`
	Text01 string `json:"text0_1,omitempty"`
	Text10 string `json:"text1_0,omitempty"`
	Text11 string `json:"text1_1,omitempty"`
	Text20 string `json:"text2_0,omitempty"`
	Text21 string `json:"text2_1,omitempty"`
	Text30 string `json:"text3_0,omitempty"`
	Text31 string `json:"text3_1,omitempty"`
	Text40 string `json:"text4_0,omitempty"`
	Text41 string `json:"text4_1,omitempty"`
	Text50 string `json:"text5_0,omitempty"`
	Text51 string `json:"text5_1,omitempty"`
	Text60 string `json:"text6_0,omitempty"`
	Text61 string `json:"text6_1,omitempty"`
	Text70 string `json:"text7_0,omitempty"`
	Text71 string `json:"text7_1,omitempty"`
}

type NpcTextLocaleSlice []NpcTextLocale

type DBNpcTextLocale struct {
	bun.BaseModel `bun:"table:npc_text_locale,alias:npc_text_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"Locale"`
	Text00 string `bun:"Text0_0"`
	Text01 string `bun:"Text0_1"`
	Text10 string `bun:"Text1_0"`
	Text11 string `bun:"Text1_1"`
	Text20 string `bun:"Text2_0"`
	Text21 string `bun:"Text2_1"`
	Text30 string `bun:"Text3_0"`
	Text31 string `bun:"Text3_1"`
	Text40 string `bun:"Text4_0"`
	Text41 string `bun:"Text4_1"`
	Text50 string `bun:"Text5_0"`
	Text51 string `bun:"Text5_1"`
	Text60 string `bun:"Text6_0"`
	Text61 string `bun:"Text6_1"`
	Text70 string `bun:"Text7_0"`
	Text71 string `bun:"Text7_1"`
}

type DBNpcTextLocaleSlice []DBNpcTextLocale

func (entry *NpcTextLocale) ToDB() *DBNpcTextLocale {
	if entry == nil {
		return nil
	}
	return &DBNpcTextLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Text00: entry.Text00,
		Text01: entry.Text01,
		Text10: entry.Text10,
		Text11: entry.Text11,
		Text20: entry.Text20,
		Text21: entry.Text21,
		Text30: entry.Text30,
		Text31: entry.Text31,
		Text40: entry.Text40,
		Text41: entry.Text41,
		Text50: entry.Text50,
		Text51: entry.Text51,
		Text60: entry.Text60,
		Text61: entry.Text61,
		Text70: entry.Text70,
		Text71: entry.Text71,
	}
}

func (entry *DBNpcTextLocale) ToWeb() *NpcTextLocale {
	if entry == nil {
		return nil
	}
	return &NpcTextLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Text00: entry.Text00,
		Text01: entry.Text01,
		Text10: entry.Text10,
		Text11: entry.Text11,
		Text20: entry.Text20,
		Text21: entry.Text21,
		Text30: entry.Text30,
		Text31: entry.Text31,
		Text40: entry.Text40,
		Text41: entry.Text41,
		Text50: entry.Text50,
		Text51: entry.Text51,
		Text60: entry.Text60,
		Text61: entry.Text61,
		Text70: entry.Text70,
		Text71: entry.Text71,
	}
}

func (data NpcTextLocaleSlice) ToDB() DBNpcTextLocaleSlice {
	result := make(DBNpcTextLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBNpcTextLocaleSlice) ToWeb() NpcTextLocaleSlice {
	result := make(NpcTextLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
