package model

import "github.com/uptrace/bun"

type ItemSetNamesLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Name string `json:"name,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type ItemSetNamesLocaleSlice []ItemSetNamesLocale

type DBItemSetNamesLocale struct {
	bun.BaseModel `bun:"table:item_set_names_locale,alias:item_set_names_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"locale"`
	Name string `bun:"Name"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBItemSetNamesLocaleSlice []DBItemSetNamesLocale

func (entry *ItemSetNamesLocale) ToDB() *DBItemSetNamesLocale {
	if entry == nil {
		return nil
	}
	return &DBItemSetNamesLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Name: entry.Name,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBItemSetNamesLocale) ToWeb() *ItemSetNamesLocale {
	if entry == nil {
		return nil
	}
	return &ItemSetNamesLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Name: entry.Name,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data ItemSetNamesLocaleSlice) ToDB() DBItemSetNamesLocaleSlice {
	result := make(DBItemSetNamesLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemSetNamesLocaleSlice) ToWeb() ItemSetNamesLocaleSlice {
	result := make(ItemSetNamesLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
