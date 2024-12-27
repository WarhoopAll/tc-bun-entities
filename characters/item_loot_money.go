package model

import "github.com/uptrace/bun"

type ItemLootMoney struct {
	ContainerId int `json:"container_id,omitempty"`
	Money int `json:"money,omitempty"`
}

type ItemLootMoneySlice []ItemLootMoney

type DBItemLootMoney struct {
	bun.BaseModel `bun:"table:item_loot_money,alias:item_loot_money"`
	ContainerId int `bun:"container_id"`
	Money int `bun:"money"`
}

type DBItemLootMoneySlice []DBItemLootMoney

func (entry *ItemLootMoney) ToDB() *DBItemLootMoney {
	if entry == nil {
		return nil
	}
	return &DBItemLootMoney{
		ContainerId: entry.ContainerId,
		Money: entry.Money,
	}
}

func (entry *DBItemLootMoney) ToWeb() *ItemLootMoney {
	if entry == nil {
		return nil
	}
	return &ItemLootMoney{
		ContainerId: entry.ContainerId,
		Money: entry.Money,
	}
}

func (data ItemLootMoneySlice) ToDB() DBItemLootMoneySlice {
	result := make(DBItemLootMoneySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemLootMoneySlice) ToWeb() ItemLootMoneySlice {
	result := make(ItemLootMoneySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
