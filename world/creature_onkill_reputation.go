package model

import "github.com/uptrace/bun"

type CreatureOnkillReputation struct {
	CreatureId int `json:"creature_id,omitempty"`
	RewOnKillRepFaction1 int16 `json:"rewonkillrepfaction1,omitempty"`
	RewOnKillRepFaction2 int16 `json:"rewonkillrepfaction2,omitempty"`
	MaxStanding1 int8 `json:"maxstanding1,omitempty"`
	IsTeamAward1 int8 `json:"isteamaward1,omitempty"`
	RewOnKillRepValue1 int `json:"rewonkillrepvalue1,omitempty"`
	MaxStanding2 int8 `json:"maxstanding2,omitempty"`
	IsTeamAward2 int8 `json:"isteamaward2,omitempty"`
	RewOnKillRepValue2 int `json:"rewonkillrepvalue2,omitempty"`
	TeamDependent int8 `json:"teamdependent,omitempty"`
}

type CreatureOnkillReputationSlice []CreatureOnkillReputation

type DBCreatureOnkillReputation struct {
	bun.BaseModel `bun:"table:creature_onkill_reputation,alias:creature_onkill_reputation"`
	CreatureId int `bun:"creature_id"`
	RewOnKillRepFaction1 int16 `bun:"RewOnKillRepFaction1"`
	RewOnKillRepFaction2 int16 `bun:"RewOnKillRepFaction2"`
	MaxStanding1 int8 `bun:"MaxStanding1"`
	IsTeamAward1 int8 `bun:"IsTeamAward1"`
	RewOnKillRepValue1 int `bun:"RewOnKillRepValue1"`
	MaxStanding2 int8 `bun:"MaxStanding2"`
	IsTeamAward2 int8 `bun:"IsTeamAward2"`
	RewOnKillRepValue2 int `bun:"RewOnKillRepValue2"`
	TeamDependent int8 `bun:"TeamDependent"`
}

type DBCreatureOnkillReputationSlice []DBCreatureOnkillReputation

func (entry *CreatureOnkillReputation) ToDB() *DBCreatureOnkillReputation {
	if entry == nil {
		return nil
	}
	return &DBCreatureOnkillReputation{
		CreatureId: entry.CreatureId,
		RewOnKillRepFaction1: entry.RewOnKillRepFaction1,
		RewOnKillRepFaction2: entry.RewOnKillRepFaction2,
		MaxStanding1: entry.MaxStanding1,
		IsTeamAward1: entry.IsTeamAward1,
		RewOnKillRepValue1: entry.RewOnKillRepValue1,
		MaxStanding2: entry.MaxStanding2,
		IsTeamAward2: entry.IsTeamAward2,
		RewOnKillRepValue2: entry.RewOnKillRepValue2,
		TeamDependent: entry.TeamDependent,
	}
}

func (entry *DBCreatureOnkillReputation) ToWeb() *CreatureOnkillReputation {
	if entry == nil {
		return nil
	}
	return &CreatureOnkillReputation{
		CreatureId: entry.CreatureId,
		RewOnKillRepFaction1: entry.RewOnKillRepFaction1,
		RewOnKillRepFaction2: entry.RewOnKillRepFaction2,
		MaxStanding1: entry.MaxStanding1,
		IsTeamAward1: entry.IsTeamAward1,
		RewOnKillRepValue1: entry.RewOnKillRepValue1,
		MaxStanding2: entry.MaxStanding2,
		IsTeamAward2: entry.IsTeamAward2,
		RewOnKillRepValue2: entry.RewOnKillRepValue2,
		TeamDependent: entry.TeamDependent,
	}
}

func (data CreatureOnkillReputationSlice) ToDB() DBCreatureOnkillReputationSlice {
	result := make(DBCreatureOnkillReputationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureOnkillReputationSlice) ToWeb() CreatureOnkillReputationSlice {
	result := make(CreatureOnkillReputationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
