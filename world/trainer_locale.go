package model

import "github.com/uptrace/bun"

type TrainerLocale struct {
	Id int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	GreetingLang string `json:"greeting_lang,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type TrainerLocaleSlice []TrainerLocale

type DBTrainerLocale struct {
	bun.BaseModel `bun:"table:trainer_locale,alias:trainer_locale"`
	Id int `bun:"Id"`
	Locale string `bun:"locale"`
	GreetingLang string `bun:"Greeting_lang"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBTrainerLocaleSlice []DBTrainerLocale

func (entry *TrainerLocale) ToDB() *DBTrainerLocale {
	if entry == nil {
		return nil
	}
	return &DBTrainerLocale{
		Id: entry.Id,
		Locale: entry.Locale,
		GreetingLang: entry.GreetingLang,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBTrainerLocale) ToWeb() *TrainerLocale {
	if entry == nil {
		return nil
	}
	return &TrainerLocale{
		Id: entry.Id,
		Locale: entry.Locale,
		GreetingLang: entry.GreetingLang,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data TrainerLocaleSlice) ToDB() DBTrainerLocaleSlice {
	result := make(DBTrainerLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBTrainerLocaleSlice) ToWeb() TrainerLocaleSlice {
	result := make(TrainerLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
