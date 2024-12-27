package model

import "github.com/uptrace/bun"

type ReputationSpilloverTemplate struct {
	Faction int16 `json:"faction,omitempty"`
	Faction1 int16 `json:"faction1,omitempty"`
	Rate1 float64 `json:"rate_1,omitempty"`
	Rank1 int8 `json:"rank_1,omitempty"`
	Faction2 int16 `json:"faction2,omitempty"`
	Rate2 float64 `json:"rate_2,omitempty"`
	Rank2 int8 `json:"rank_2,omitempty"`
	Faction3 int16 `json:"faction3,omitempty"`
	Rate3 float64 `json:"rate_3,omitempty"`
	Rank3 int8 `json:"rank_3,omitempty"`
	Faction4 int16 `json:"faction4,omitempty"`
	Rate4 float64 `json:"rate_4,omitempty"`
	Rank4 int8 `json:"rank_4,omitempty"`
}

type ReputationSpilloverTemplateSlice []ReputationSpilloverTemplate

type DBReputationSpilloverTemplate struct {
	bun.BaseModel `bun:"table:reputation_spillover_template,alias:reputation_spillover_template"`
	Faction int16 `bun:"faction"`
	Faction1 int16 `bun:"faction1"`
	Rate1 float64 `bun:"rate_1"`
	Rank1 int8 `bun:"rank_1"`
	Faction2 int16 `bun:"faction2"`
	Rate2 float64 `bun:"rate_2"`
	Rank2 int8 `bun:"rank_2"`
	Faction3 int16 `bun:"faction3"`
	Rate3 float64 `bun:"rate_3"`
	Rank3 int8 `bun:"rank_3"`
	Faction4 int16 `bun:"faction4"`
	Rate4 float64 `bun:"rate_4"`
	Rank4 int8 `bun:"rank_4"`
}

type DBReputationSpilloverTemplateSlice []DBReputationSpilloverTemplate

func (entry *ReputationSpilloverTemplate) ToDB() *DBReputationSpilloverTemplate {
	if entry == nil {
		return nil
	}
	return &DBReputationSpilloverTemplate{
		Faction: entry.Faction,
		Faction1: entry.Faction1,
		Rate1: entry.Rate1,
		Rank1: entry.Rank1,
		Faction2: entry.Faction2,
		Rate2: entry.Rate2,
		Rank2: entry.Rank2,
		Faction3: entry.Faction3,
		Rate3: entry.Rate3,
		Rank3: entry.Rank3,
		Faction4: entry.Faction4,
		Rate4: entry.Rate4,
		Rank4: entry.Rank4,
	}
}

func (entry *DBReputationSpilloverTemplate) ToWeb() *ReputationSpilloverTemplate {
	if entry == nil {
		return nil
	}
	return &ReputationSpilloverTemplate{
		Faction: entry.Faction,
		Faction1: entry.Faction1,
		Rate1: entry.Rate1,
		Rank1: entry.Rank1,
		Faction2: entry.Faction2,
		Rate2: entry.Rate2,
		Rank2: entry.Rank2,
		Faction3: entry.Faction3,
		Rate3: entry.Rate3,
		Rank3: entry.Rank3,
		Faction4: entry.Faction4,
		Rate4: entry.Rate4,
		Rank4: entry.Rank4,
	}
}

func (data ReputationSpilloverTemplateSlice) ToDB() DBReputationSpilloverTemplateSlice {
	result := make(DBReputationSpilloverTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBReputationSpilloverTemplateSlice) ToWeb() ReputationSpilloverTemplateSlice {
	result := make(ReputationSpilloverTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
