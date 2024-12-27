package model

import "github.com/uptrace/bun"

type CreatureTextLocale struct {
	CreatureID int `json:"creatureid,omitempty"`
	GroupID int8 `json:"groupid,omitempty"`
	ID int8 `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Text string `json:"text,omitempty"`
}

type CreatureTextLocaleSlice []CreatureTextLocale

type DBCreatureTextLocale struct {
	bun.BaseModel `bun:"table:creature_text_locale,alias:creature_text_locale"`
	CreatureID int `bun:"CreatureID"`
	GroupID int8 `bun:"GroupID"`
	ID int8 `bun:"ID"`
	Locale string `bun:"Locale"`
	Text string `bun:"Text"`
}

type DBCreatureTextLocaleSlice []DBCreatureTextLocale

func (entry *CreatureTextLocale) ToDB() *DBCreatureTextLocale {
	if entry == nil {
		return nil
	}
	return &DBCreatureTextLocale{
		CreatureID: entry.CreatureID,
		GroupID: entry.GroupID,
		ID: entry.ID,
		Locale: entry.Locale,
		Text: entry.Text,
	}
}

func (entry *DBCreatureTextLocale) ToWeb() *CreatureTextLocale {
	if entry == nil {
		return nil
	}
	return &CreatureTextLocale{
		CreatureID: entry.CreatureID,
		GroupID: entry.GroupID,
		ID: entry.ID,
		Locale: entry.Locale,
		Text: entry.Text,
	}
}

func (data CreatureTextLocaleSlice) ToDB() DBCreatureTextLocaleSlice {
	result := make(DBCreatureTextLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureTextLocaleSlice) ToWeb() CreatureTextLocaleSlice {
	result := make(CreatureTextLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
