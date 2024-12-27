package model

import "github.com/uptrace/bun"

type MillingLootTemplate struct {
	Entry int `json:"entry,omitempty"`
	Item int `json:"item,omitempty"`
	Reference int `json:"reference,omitempty"`
	Chance float64 `json:"chance,omitempty"`
	QuestRequired bool `json:"questrequired,omitempty"`
	LootMode int16 `json:"lootmode,omitempty"`
	GroupId int8 `json:"groupid,omitempty"`
	MinCount int8 `json:"mincount,omitempty"`
	MaxCount int8 `json:"maxcount,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type MillingLootTemplateSlice []MillingLootTemplate

type DBMillingLootTemplate struct {
	bun.BaseModel `bun:"table:milling_loot_template,alias:milling_loot_template"`
	Entry int `bun:"Entry"`
	Item int `bun:"Item"`
	Reference int `bun:"Reference"`
	Chance float64 `bun:"Chance"`
	QuestRequired bool `bun:"QuestRequired"`
	LootMode int16 `bun:"LootMode"`
	GroupId int8 `bun:"GroupId"`
	MinCount int8 `bun:"MinCount"`
	MaxCount int8 `bun:"MaxCount"`
	Comment string `bun:"Comment"`
}

type DBMillingLootTemplateSlice []DBMillingLootTemplate

func (entry *MillingLootTemplate) ToDB() *DBMillingLootTemplate {
	if entry == nil {
		return nil
	}
	return &DBMillingLootTemplate{
		Entry: entry.Entry,
		Item: entry.Item,
		Reference: entry.Reference,
		Chance: entry.Chance,
		QuestRequired: entry.QuestRequired,
		LootMode: entry.LootMode,
		GroupId: entry.GroupId,
		MinCount: entry.MinCount,
		MaxCount: entry.MaxCount,
		Comment: entry.Comment,
	}
}

func (entry *DBMillingLootTemplate) ToWeb() *MillingLootTemplate {
	if entry == nil {
		return nil
	}
	return &MillingLootTemplate{
		Entry: entry.Entry,
		Item: entry.Item,
		Reference: entry.Reference,
		Chance: entry.Chance,
		QuestRequired: entry.QuestRequired,
		LootMode: entry.LootMode,
		GroupId: entry.GroupId,
		MinCount: entry.MinCount,
		MaxCount: entry.MaxCount,
		Comment: entry.Comment,
	}
}

func (data MillingLootTemplateSlice) ToDB() DBMillingLootTemplateSlice {
	result := make(DBMillingLootTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBMillingLootTemplateSlice) ToWeb() MillingLootTemplateSlice {
	result := make(MillingLootTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
