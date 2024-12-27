package model

import "github.com/uptrace/bun"

type ItemTemplateLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type ItemTemplateLocaleSlice []ItemTemplateLocale

type DBItemTemplateLocale struct {
	bun.BaseModel `bun:"table:item_template_locale,alias:item_template_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"locale"`
	Name string `bun:"Name"`
	Description string `bun:"Description"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBItemTemplateLocaleSlice []DBItemTemplateLocale

func (entry *ItemTemplateLocale) ToDB() *DBItemTemplateLocale {
	if entry == nil {
		return nil
	}
	return &DBItemTemplateLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Name: entry.Name,
		Description: entry.Description,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBItemTemplateLocale) ToWeb() *ItemTemplateLocale {
	if entry == nil {
		return nil
	}
	return &ItemTemplateLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Name: entry.Name,
		Description: entry.Description,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data ItemTemplateLocaleSlice) ToDB() DBItemTemplateLocaleSlice {
	result := make(DBItemTemplateLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemTemplateLocaleSlice) ToWeb() ItemTemplateLocaleSlice {
	result := make(ItemTemplateLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
