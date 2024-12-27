package model

import "github.com/uptrace/bun"

type QuestPoi struct {
	QuestID int `json:"questid,omitempty"`
	Id int `json:"id,omitempty"`
	ObjectiveIndex int `json:"objectiveindex,omitempty"`
	MapID int `json:"mapid,omitempty"`
	WorldMapAreaId int `json:"worldmapareaid,omitempty"`
	Floor int `json:"floor,omitempty"`
	Priority int `json:"priority,omitempty"`
	Flags int `json:"flags,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestPoiSlice []QuestPoi

type DBQuestPoi struct {
	bun.BaseModel `bun:"table:quest_poi,alias:quest_poi"`
	QuestID int `bun:"QuestID"`
	Id int `bun:"id"`
	ObjectiveIndex int `bun:"ObjectiveIndex"`
	MapID int `bun:"MapID"`
	WorldMapAreaId int `bun:"WorldMapAreaId"`
	Floor int `bun:"Floor"`
	Priority int `bun:"Priority"`
	Flags int `bun:"Flags"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestPoiSlice []DBQuestPoi

func (entry *QuestPoi) ToDB() *DBQuestPoi {
	if entry == nil {
		return nil
	}
	return &DBQuestPoi{
		QuestID: entry.QuestID,
		Id: entry.Id,
		ObjectiveIndex: entry.ObjectiveIndex,
		MapID: entry.MapID,
		WorldMapAreaId: entry.WorldMapAreaId,
		Floor: entry.Floor,
		Priority: entry.Priority,
		Flags: entry.Flags,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestPoi) ToWeb() *QuestPoi {
	if entry == nil {
		return nil
	}
	return &QuestPoi{
		QuestID: entry.QuestID,
		Id: entry.Id,
		ObjectiveIndex: entry.ObjectiveIndex,
		MapID: entry.MapID,
		WorldMapAreaId: entry.WorldMapAreaId,
		Floor: entry.Floor,
		Priority: entry.Priority,
		Flags: entry.Flags,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestPoiSlice) ToDB() DBQuestPoiSlice {
	result := make(DBQuestPoiSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestPoiSlice) ToWeb() QuestPoiSlice {
	result := make(QuestPoiSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
