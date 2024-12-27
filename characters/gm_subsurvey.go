package model

import "github.com/uptrace/bun"

type GmSubsurvey struct {
	SurveyId int `json:"surveyid,omitempty"`
	QuestionId int `json:"questionid,omitempty"`
	Answer int `json:"answer,omitempty"`
	AnswerComment string `json:"answercomment,omitempty"`
}

type GmSubsurveySlice []GmSubsurvey

type DBGmSubsurvey struct {
	bun.BaseModel `bun:"table:gm_subsurvey,alias:gm_subsurvey"`
	SurveyId int `bun:"surveyId"`
	QuestionId int `bun:"questionId"`
	Answer int `bun:"answer"`
	AnswerComment string `bun:"answerComment"`
}

type DBGmSubsurveySlice []DBGmSubsurvey

func (entry *GmSubsurvey) ToDB() *DBGmSubsurvey {
	if entry == nil {
		return nil
	}
	return &DBGmSubsurvey{
		SurveyId: entry.SurveyId,
		QuestionId: entry.QuestionId,
		Answer: entry.Answer,
		AnswerComment: entry.AnswerComment,
	}
}

func (entry *DBGmSubsurvey) ToWeb() *GmSubsurvey {
	if entry == nil {
		return nil
	}
	return &GmSubsurvey{
		SurveyId: entry.SurveyId,
		QuestionId: entry.QuestionId,
		Answer: entry.Answer,
		AnswerComment: entry.AnswerComment,
	}
}

func (data GmSubsurveySlice) ToDB() DBGmSubsurveySlice {
	result := make(DBGmSubsurveySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGmSubsurveySlice) ToWeb() GmSubsurveySlice {
	result := make(GmSubsurveySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
