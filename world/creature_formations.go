package model

import "github.com/uptrace/bun"

type CreatureFormations struct {
	LeaderGUID int `json:"leaderguid,omitempty"`
	MemberGUID int `json:"memberguid,omitempty"`
	Dist float64 `json:"dist,omitempty"`
	Angle float64 `json:"angle,omitempty"`
	GroupAI int `json:"groupai,omitempty"`
	Point1 int16 `json:"point_1,omitempty"`
	Point2 int16 `json:"point_2,omitempty"`
}

type CreatureFormationsSlice []CreatureFormations

type DBCreatureFormations struct {
	bun.BaseModel `bun:"table:creature_formations,alias:creature_formations"`
	LeaderGUID int `bun:"leaderGUID"`
	MemberGUID int `bun:"memberGUID"`
	Dist float64 `bun:"dist"`
	Angle float64 `bun:"angle"`
	GroupAI int `bun:"groupAI"`
	Point1 int16 `bun:"point_1"`
	Point2 int16 `bun:"point_2"`
}

type DBCreatureFormationsSlice []DBCreatureFormations

func (entry *CreatureFormations) ToDB() *DBCreatureFormations {
	if entry == nil {
		return nil
	}
	return &DBCreatureFormations{
		LeaderGUID: entry.LeaderGUID,
		MemberGUID: entry.MemberGUID,
		Dist: entry.Dist,
		Angle: entry.Angle,
		GroupAI: entry.GroupAI,
		Point1: entry.Point1,
		Point2: entry.Point2,
	}
}

func (entry *DBCreatureFormations) ToWeb() *CreatureFormations {
	if entry == nil {
		return nil
	}
	return &CreatureFormations{
		LeaderGUID: entry.LeaderGUID,
		MemberGUID: entry.MemberGUID,
		Dist: entry.Dist,
		Angle: entry.Angle,
		GroupAI: entry.GroupAI,
		Point1: entry.Point1,
		Point2: entry.Point2,
	}
}

func (data CreatureFormationsSlice) ToDB() DBCreatureFormationsSlice {
	result := make(DBCreatureFormationsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureFormationsSlice) ToWeb() CreatureFormationsSlice {
	result := make(CreatureFormationsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
