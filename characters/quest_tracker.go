package model

import "github.com/uptrace/bun"

type QuestTracker struct {
	Id int32 `json:"id,omitempty"`
	CharacterGuid int `json:"character_guid,omitempty"`
	QuestAcceptTime time.Time `json:"quest_accept_time,omitempty"`
	QuestCompleteTime time.Time `json:"quest_complete_time,omitempty"`
	QuestAbandonTime time.Time `json:"quest_abandon_time,omitempty"`
	CompletedByGm bool `json:"completed_by_gm,omitempty"`
	CoreHash string `json:"core_hash,omitempty"`
	CoreRevision string `json:"core_revision,omitempty"`
}

type QuestTrackerSlice []QuestTracker

type DBQuestTracker struct {
	bun.BaseModel `bun:"table:quest_tracker,alias:quest_tracker"`
	Id int32 `bun:"id"`
	CharacterGuid int `bun:"character_guid"`
	QuestAcceptTime time.Time `bun:"quest_accept_time"`
	QuestCompleteTime time.Time `bun:"quest_complete_time"`
	QuestAbandonTime time.Time `bun:"quest_abandon_time"`
	CompletedByGm bool `bun:"completed_by_gm"`
	CoreHash string `bun:"core_hash"`
	CoreRevision string `bun:"core_revision"`
}

type DBQuestTrackerSlice []DBQuestTracker

func (entry *QuestTracker) ToDB() *DBQuestTracker {
	if entry == nil {
		return nil
	}
	return &DBQuestTracker{
		Id: entry.Id,
		CharacterGuid: entry.CharacterGuid,
		QuestAcceptTime: entry.QuestAcceptTime,
		QuestCompleteTime: entry.QuestCompleteTime,
		QuestAbandonTime: entry.QuestAbandonTime,
		CompletedByGm: entry.CompletedByGm,
		CoreHash: entry.CoreHash,
		CoreRevision: entry.CoreRevision,
	}
}

func (entry *DBQuestTracker) ToWeb() *QuestTracker {
	if entry == nil {
		return nil
	}
	return &QuestTracker{
		Id: entry.Id,
		CharacterGuid: entry.CharacterGuid,
		QuestAcceptTime: entry.QuestAcceptTime,
		QuestCompleteTime: entry.QuestCompleteTime,
		QuestAbandonTime: entry.QuestAbandonTime,
		CompletedByGm: entry.CompletedByGm,
		CoreHash: entry.CoreHash,
		CoreRevision: entry.CoreRevision,
	}
}

func (data QuestTrackerSlice) ToDB() DBQuestTrackerSlice {
	result := make(DBQuestTrackerSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestTrackerSlice) ToWeb() QuestTrackerSlice {
	result := make(QuestTrackerSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
