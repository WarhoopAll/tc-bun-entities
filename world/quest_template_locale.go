package model

import "github.com/uptrace/bun"

type QuestTemplateLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Title string `json:"title,omitempty"`
	Details string `json:"details,omitempty"`
	Objectives string `json:"objectives,omitempty"`
	EndText string `json:"endtext,omitempty"`
	CompletedText string `json:"completedtext,omitempty"`
	ObjectiveText1 string `json:"objectivetext1,omitempty"`
	ObjectiveText2 string `json:"objectivetext2,omitempty"`
	ObjectiveText3 string `json:"objectivetext3,omitempty"`
	ObjectiveText4 string `json:"objectivetext4,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestTemplateLocaleSlice []QuestTemplateLocale

type DBQuestTemplateLocale struct {
	bun.BaseModel `bun:"table:quest_template_locale,alias:quest_template_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"locale"`
	Title string `bun:"Title"`
	Details string `bun:"Details"`
	Objectives string `bun:"Objectives"`
	EndText string `bun:"EndText"`
	CompletedText string `bun:"CompletedText"`
	ObjectiveText1 string `bun:"ObjectiveText1"`
	ObjectiveText2 string `bun:"ObjectiveText2"`
	ObjectiveText3 string `bun:"ObjectiveText3"`
	ObjectiveText4 string `bun:"ObjectiveText4"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestTemplateLocaleSlice []DBQuestTemplateLocale

func (entry *QuestTemplateLocale) ToDB() *DBQuestTemplateLocale {
	if entry == nil {
		return nil
	}
	return &DBQuestTemplateLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Title: entry.Title,
		Details: entry.Details,
		Objectives: entry.Objectives,
		EndText: entry.EndText,
		CompletedText: entry.CompletedText,
		ObjectiveText1: entry.ObjectiveText1,
		ObjectiveText2: entry.ObjectiveText2,
		ObjectiveText3: entry.ObjectiveText3,
		ObjectiveText4: entry.ObjectiveText4,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestTemplateLocale) ToWeb() *QuestTemplateLocale {
	if entry == nil {
		return nil
	}
	return &QuestTemplateLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Title: entry.Title,
		Details: entry.Details,
		Objectives: entry.Objectives,
		EndText: entry.EndText,
		CompletedText: entry.CompletedText,
		ObjectiveText1: entry.ObjectiveText1,
		ObjectiveText2: entry.ObjectiveText2,
		ObjectiveText3: entry.ObjectiveText3,
		ObjectiveText4: entry.ObjectiveText4,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestTemplateLocaleSlice) ToDB() DBQuestTemplateLocaleSlice {
	result := make(DBQuestTemplateLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestTemplateLocaleSlice) ToWeb() QuestTemplateLocaleSlice {
	result := make(QuestTemplateLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
