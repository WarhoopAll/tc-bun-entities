package model

import "github.com/uptrace/bun"

type GameobjectTemplateLocale struct {
	Entry int `json:"entry,omitempty"`
	Locale string `json:"locale,omitempty"`
	Name string `json:"name,omitempty"`
	CastBarCaption string `json:"castbarcaption,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type GameobjectTemplateLocaleSlice []GameobjectTemplateLocale

type DBGameobjectTemplateLocale struct {
	bun.BaseModel `bun:"table:gameobject_template_locale,alias:gameobject_template_locale"`
	Entry int `bun:"entry"`
	Locale string `bun:"locale"`
	Name string `bun:"name"`
	CastBarCaption string `bun:"castBarCaption"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBGameobjectTemplateLocaleSlice []DBGameobjectTemplateLocale

func (entry *GameobjectTemplateLocale) ToDB() *DBGameobjectTemplateLocale {
	if entry == nil {
		return nil
	}
	return &DBGameobjectTemplateLocale{
		Entry: entry.Entry,
		Locale: entry.Locale,
		Name: entry.Name,
		CastBarCaption: entry.CastBarCaption,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBGameobjectTemplateLocale) ToWeb() *GameobjectTemplateLocale {
	if entry == nil {
		return nil
	}
	return &GameobjectTemplateLocale{
		Entry: entry.Entry,
		Locale: entry.Locale,
		Name: entry.Name,
		CastBarCaption: entry.CastBarCaption,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data GameobjectTemplateLocaleSlice) ToDB() DBGameobjectTemplateLocaleSlice {
	result := make(DBGameobjectTemplateLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectTemplateLocaleSlice) ToWeb() GameobjectTemplateLocaleSlice {
	result := make(GameobjectTemplateLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
