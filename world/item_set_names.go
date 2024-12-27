package model

import "github.com/uptrace/bun"

type ItemSetNames struct {
	Entry int `json:"entry,omitempty"`
	Name string `json:"name,omitempty"`
	InventoryType int8 `json:"inventorytype,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type ItemSetNamesSlice []ItemSetNames

type DBItemSetNames struct {
	bun.BaseModel `bun:"table:item_set_names,alias:item_set_names"`
	Entry int `bun:"entry"`
	Name string `bun:"name"`
	InventoryType int8 `bun:"InventoryType"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBItemSetNamesSlice []DBItemSetNames

func (entry *ItemSetNames) ToDB() *DBItemSetNames {
	if entry == nil {
		return nil
	}
	return &DBItemSetNames{
		Entry: entry.Entry,
		Name: entry.Name,
		InventoryType: entry.InventoryType,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBItemSetNames) ToWeb() *ItemSetNames {
	if entry == nil {
		return nil
	}
	return &ItemSetNames{
		Entry: entry.Entry,
		Name: entry.Name,
		InventoryType: entry.InventoryType,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data ItemSetNamesSlice) ToDB() DBItemSetNamesSlice {
	result := make(DBItemSetNamesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemSetNamesSlice) ToWeb() ItemSetNamesSlice {
	result := make(ItemSetNamesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
