package model

import "github.com/uptrace/bun"

type QuestGreetingLocale struct {
	ID int `json:"id,omitempty"`
	Type int8 `json:"type,omitempty"`
	Locale string `json:"locale,omitempty"`
	Greeting string `json:"greeting,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestGreetingLocaleSlice []QuestGreetingLocale

type DBQuestGreetingLocale struct {
	bun.BaseModel `bun:"table:quest_greeting_locale,alias:quest_greeting_locale"`
	ID int `bun:"ID"`
	Type int8 `bun:"Type"`
	Locale string `bun:"locale"`
	Greeting string `bun:"Greeting"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestGreetingLocaleSlice []DBQuestGreetingLocale

func (entry *QuestGreetingLocale) ToDB() *DBQuestGreetingLocale {
	if entry == nil {
		return nil
	}
	return &DBQuestGreetingLocale{
		ID: entry.ID,
		Type: entry.Type,
		Locale: entry.Locale,
		Greeting: entry.Greeting,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestGreetingLocale) ToWeb() *QuestGreetingLocale {
	if entry == nil {
		return nil
	}
	return &QuestGreetingLocale{
		ID: entry.ID,
		Type: entry.Type,
		Locale: entry.Locale,
		Greeting: entry.Greeting,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestGreetingLocaleSlice) ToDB() DBQuestGreetingLocaleSlice {
	result := make(DBQuestGreetingLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestGreetingLocaleSlice) ToWeb() QuestGreetingLocaleSlice {
	result := make(QuestGreetingLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
