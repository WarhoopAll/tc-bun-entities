package model

import "github.com/uptrace/bun"

type CreatureTemplateLocale struct {
	Entry int `json:"entry,omitempty"`
	Locale string `json:"locale,omitempty"`
	Name string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type CreatureTemplateLocaleSlice []CreatureTemplateLocale

type DBCreatureTemplateLocale struct {
	bun.BaseModel `bun:"table:creature_template_locale,alias:creature_template_locale"`
	Entry int `bun:"entry"`
	Locale string `bun:"locale"`
	Name string `bun:"Name"`
	Title string `bun:"Title"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBCreatureTemplateLocaleSlice []DBCreatureTemplateLocale

func (entry *CreatureTemplateLocale) ToDB() *DBCreatureTemplateLocale {
	if entry == nil {
		return nil
	}
	return &DBCreatureTemplateLocale{
		Entry: entry.Entry,
		Locale: entry.Locale,
		Name: entry.Name,
		Title: entry.Title,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBCreatureTemplateLocale) ToWeb() *CreatureTemplateLocale {
	if entry == nil {
		return nil
	}
	return &CreatureTemplateLocale{
		Entry: entry.Entry,
		Locale: entry.Locale,
		Name: entry.Name,
		Title: entry.Title,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data CreatureTemplateLocaleSlice) ToDB() DBCreatureTemplateLocaleSlice {
	result := make(DBCreatureTemplateLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureTemplateLocaleSlice) ToWeb() CreatureTemplateLocaleSlice {
	result := make(CreatureTemplateLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
