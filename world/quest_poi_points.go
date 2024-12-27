package model

import "github.com/uptrace/bun"

type QuestPoiPoints struct {
	QuestID int `json:"questid,omitempty"`
	Idx1 int `json:"idx1,omitempty"`
	Idx2 int `json:"idx2,omitempty"`
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestPoiPointsSlice []QuestPoiPoints

type DBQuestPoiPoints struct {
	bun.BaseModel `bun:"table:quest_poi_points,alias:quest_poi_points"`
	QuestID int `bun:"QuestID"`
	Idx1 int `bun:"Idx1"`
	Idx2 int `bun:"Idx2"`
	X int `bun:"X"`
	Y int `bun:"Y"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestPoiPointsSlice []DBQuestPoiPoints

func (entry *QuestPoiPoints) ToDB() *DBQuestPoiPoints {
	if entry == nil {
		return nil
	}
	return &DBQuestPoiPoints{
		QuestID: entry.QuestID,
		Idx1: entry.Idx1,
		Idx2: entry.Idx2,
		X: entry.X,
		Y: entry.Y,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestPoiPoints) ToWeb() *QuestPoiPoints {
	if entry == nil {
		return nil
	}
	return &QuestPoiPoints{
		QuestID: entry.QuestID,
		Idx1: entry.Idx1,
		Idx2: entry.Idx2,
		X: entry.X,
		Y: entry.Y,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestPoiPointsSlice) ToDB() DBQuestPoiPointsSlice {
	result := make(DBQuestPoiPointsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestPoiPointsSlice) ToWeb() QuestPoiPointsSlice {
	result := make(QuestPoiPointsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
