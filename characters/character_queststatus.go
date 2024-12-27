package model

import "github.com/uptrace/bun"

type CharacterQueststatus struct {
	Guid int `json:"guid,omitempty"`
	Quest int `json:"quest,omitempty"`
	Status int8 `json:"status,omitempty"`
	Explored int8 `json:"explored,omitempty"`
	Timer int `json:"timer,omitempty"`
	Mobcount1 int16 `json:"mobcount1,omitempty"`
	Mobcount2 int16 `json:"mobcount2,omitempty"`
	Mobcount3 int16 `json:"mobcount3,omitempty"`
	Mobcount4 int16 `json:"mobcount4,omitempty"`
	Itemcount1 int16 `json:"itemcount1,omitempty"`
	Itemcount2 int16 `json:"itemcount2,omitempty"`
	Itemcount3 int16 `json:"itemcount3,omitempty"`
	Itemcount4 int16 `json:"itemcount4,omitempty"`
	Itemcount5 int16 `json:"itemcount5,omitempty"`
	Itemcount6 int16 `json:"itemcount6,omitempty"`
	Playercount int16 `json:"playercount,omitempty"`
}

type CharacterQueststatusSlice []CharacterQueststatus

type DBCharacterQueststatus struct {
	bun.BaseModel `bun:"table:character_queststatus,alias:character_queststatus"`
	Guid int `bun:"guid"`
	Quest int `bun:"quest"`
	Status int8 `bun:"status"`
	Explored int8 `bun:"explored"`
	Timer int `bun:"timer"`
	Mobcount1 int16 `bun:"mobcount1"`
	Mobcount2 int16 `bun:"mobcount2"`
	Mobcount3 int16 `bun:"mobcount3"`
	Mobcount4 int16 `bun:"mobcount4"`
	Itemcount1 int16 `bun:"itemcount1"`
	Itemcount2 int16 `bun:"itemcount2"`
	Itemcount3 int16 `bun:"itemcount3"`
	Itemcount4 int16 `bun:"itemcount4"`
	Itemcount5 int16 `bun:"itemcount5"`
	Itemcount6 int16 `bun:"itemcount6"`
	Playercount int16 `bun:"playercount"`
}

type DBCharacterQueststatusSlice []DBCharacterQueststatus

func (entry *CharacterQueststatus) ToDB() *DBCharacterQueststatus {
	if entry == nil {
		return nil
	}
	return &DBCharacterQueststatus{
		Guid: entry.Guid,
		Quest: entry.Quest,
		Status: entry.Status,
		Explored: entry.Explored,
		Timer: entry.Timer,
		Mobcount1: entry.Mobcount1,
		Mobcount2: entry.Mobcount2,
		Mobcount3: entry.Mobcount3,
		Mobcount4: entry.Mobcount4,
		Itemcount1: entry.Itemcount1,
		Itemcount2: entry.Itemcount2,
		Itemcount3: entry.Itemcount3,
		Itemcount4: entry.Itemcount4,
		Itemcount5: entry.Itemcount5,
		Itemcount6: entry.Itemcount6,
		Playercount: entry.Playercount,
	}
}

func (entry *DBCharacterQueststatus) ToWeb() *CharacterQueststatus {
	if entry == nil {
		return nil
	}
	return &CharacterQueststatus{
		Guid: entry.Guid,
		Quest: entry.Quest,
		Status: entry.Status,
		Explored: entry.Explored,
		Timer: entry.Timer,
		Mobcount1: entry.Mobcount1,
		Mobcount2: entry.Mobcount2,
		Mobcount3: entry.Mobcount3,
		Mobcount4: entry.Mobcount4,
		Itemcount1: entry.Itemcount1,
		Itemcount2: entry.Itemcount2,
		Itemcount3: entry.Itemcount3,
		Itemcount4: entry.Itemcount4,
		Itemcount5: entry.Itemcount5,
		Itemcount6: entry.Itemcount6,
		Playercount: entry.Playercount,
	}
}

func (data CharacterQueststatusSlice) ToDB() DBCharacterQueststatusSlice {
	result := make(DBCharacterQueststatusSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterQueststatusSlice) ToWeb() CharacterQueststatusSlice {
	result := make(CharacterQueststatusSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
