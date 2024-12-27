package model

import "github.com/uptrace/bun"

type CharacterInventory struct {
	Guid int `json:"guid,omitempty"`
	Bag int `json:"bag,omitempty"`
	Slot int8 `json:"slot,omitempty"`
	Item int `json:"item,omitempty"`
}

type CharacterInventorySlice []CharacterInventory

type DBCharacterInventory struct {
	bun.BaseModel `bun:"table:character_inventory,alias:character_inventory"`
	Guid int `bun:"guid"`
	Bag int `bun:"bag"`
	Slot int8 `bun:"slot"`
	Item int `bun:"item"`
}

type DBCharacterInventorySlice []DBCharacterInventory

func (entry *CharacterInventory) ToDB() *DBCharacterInventory {
	if entry == nil {
		return nil
	}
	return &DBCharacterInventory{
		Guid: entry.Guid,
		Bag: entry.Bag,
		Slot: entry.Slot,
		Item: entry.Item,
	}
}

func (entry *DBCharacterInventory) ToWeb() *CharacterInventory {
	if entry == nil {
		return nil
	}
	return &CharacterInventory{
		Guid: entry.Guid,
		Bag: entry.Bag,
		Slot: entry.Slot,
		Item: entry.Item,
	}
}

func (data CharacterInventorySlice) ToDB() DBCharacterInventorySlice {
	result := make(DBCharacterInventorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterInventorySlice) ToWeb() CharacterInventorySlice {
	result := make(CharacterInventorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
