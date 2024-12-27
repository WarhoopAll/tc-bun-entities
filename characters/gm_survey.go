package model

import "github.com/uptrace/bun"

type GmSurvey struct {
	SurveyId int `json:"surveyid,omitempty"`
	Guid int `json:"guid,omitempty"`
	MainSurvey int `json:"mainsurvey,omitempty"`
	Comment string `json:"comment,omitempty"`
	CreateTime int `json:"createtime,omitempty"`
}

type GmSurveySlice []GmSurvey

type DBGmSurvey struct {
	bun.BaseModel `bun:"table:gm_survey,alias:gm_survey"`
	SurveyId int `bun:"surveyId"`
	Guid int `bun:"guid"`
	MainSurvey int `bun:"mainSurvey"`
	Comment string `bun:"comment"`
	CreateTime int `bun:"createTime"`
}

type DBGmSurveySlice []DBGmSurvey

func (entry *GmSurvey) ToDB() *DBGmSurvey {
	if entry == nil {
		return nil
	}
	return &DBGmSurvey{
		SurveyId: entry.SurveyId,
		Guid: entry.Guid,
		MainSurvey: entry.MainSurvey,
		Comment: entry.Comment,
		CreateTime: entry.CreateTime,
	}
}

func (entry *DBGmSurvey) ToWeb() *GmSurvey {
	if entry == nil {
		return nil
	}
	return &GmSurvey{
		SurveyId: entry.SurveyId,
		Guid: entry.Guid,
		MainSurvey: entry.MainSurvey,
		Comment: entry.Comment,
		CreateTime: entry.CreateTime,
	}
}

func (data GmSurveySlice) ToDB() DBGmSurveySlice {
	result := make(DBGmSurveySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGmSurveySlice) ToWeb() GmSurveySlice {
	result := make(GmSurveySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
