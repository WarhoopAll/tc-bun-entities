package model

import "github.com/uptrace/bun"

type ItemLootItems struct {
	ContainerId int `json:"container_id,omitempty"`
	ItemId int `json:"item_id,omitempty"`
	ItemCount int `json:"item_count,omitempty"`
	ItemIndex int `json:"item_index,omitempty"`
	FollowRules bool `json:"follow_rules,omitempty"`
	Ffa bool `json:"ffa,omitempty"`
	Blocked bool `json:"blocked,omitempty"`
	Counted bool `json:"counted,omitempty"`
	UnderThreshold bool `json:"under_threshold,omitempty"`
	NeedsQuest bool `json:"needs_quest,omitempty"`
	RndProp int `json:"rnd_prop,omitempty"`
	RndSuffix int `json:"rnd_suffix,omitempty"`
}

type ItemLootItemsSlice []ItemLootItems

type DBItemLootItems struct {
	bun.BaseModel `bun:"table:item_loot_items,alias:item_loot_items"`
	ContainerId int `bun:"container_id"`
	ItemId int `bun:"item_id"`
	ItemCount int `bun:"item_count"`
	ItemIndex int `bun:"item_index"`
	FollowRules bool `bun:"follow_rules"`
	Ffa bool `bun:"ffa"`
	Blocked bool `bun:"blocked"`
	Counted bool `bun:"counted"`
	UnderThreshold bool `bun:"under_threshold"`
	NeedsQuest bool `bun:"needs_quest"`
	RndProp int `bun:"rnd_prop"`
	RndSuffix int `bun:"rnd_suffix"`
}

type DBItemLootItemsSlice []DBItemLootItems

func (entry *ItemLootItems) ToDB() *DBItemLootItems {
	if entry == nil {
		return nil
	}
	return &DBItemLootItems{
		ContainerId: entry.ContainerId,
		ItemId: entry.ItemId,
		ItemCount: entry.ItemCount,
		ItemIndex: entry.ItemIndex,
		FollowRules: entry.FollowRules,
		Ffa: entry.Ffa,
		Blocked: entry.Blocked,
		Counted: entry.Counted,
		UnderThreshold: entry.UnderThreshold,
		NeedsQuest: entry.NeedsQuest,
		RndProp: entry.RndProp,
		RndSuffix: entry.RndSuffix,
	}
}

func (entry *DBItemLootItems) ToWeb() *ItemLootItems {
	if entry == nil {
		return nil
	}
	return &ItemLootItems{
		ContainerId: entry.ContainerId,
		ItemId: entry.ItemId,
		ItemCount: entry.ItemCount,
		ItemIndex: entry.ItemIndex,
		FollowRules: entry.FollowRules,
		Ffa: entry.Ffa,
		Blocked: entry.Blocked,
		Counted: entry.Counted,
		UnderThreshold: entry.UnderThreshold,
		NeedsQuest: entry.NeedsQuest,
		RndProp: entry.RndProp,
		RndSuffix: entry.RndSuffix,
	}
}

func (data ItemLootItemsSlice) ToDB() DBItemLootItemsSlice {
	result := make(DBItemLootItemsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemLootItemsSlice) ToWeb() ItemLootItemsSlice {
	result := make(ItemLootItemsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
