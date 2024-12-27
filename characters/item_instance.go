package model

import "github.com/uptrace/bun"

type ItemInstance struct {
	Guid int `json:"guid,omitempty"`
	ItemEntry int32 `json:"itementry,omitempty"`
	OwnerGuid int `json:"owner_guid,omitempty"`
	CreatorGuid int `json:"creatorguid,omitempty"`
	GiftCreatorGuid int `json:"giftcreatorguid,omitempty"`
	Count int `json:"count,omitempty"`
	Duration int `json:"duration,omitempty"`
	Charges string `json:"charges,omitempty"`
	Flags int32 `json:"flags,omitempty"`
	Enchantments string `json:"enchantments,omitempty"`
	RandomPropertyId int16 `json:"randompropertyid,omitempty"`
	Durability int16 `json:"durability,omitempty"`
	PlayedTime int `json:"playedtime,omitempty"`
	Text string `json:"text,omitempty"`
}

type ItemInstanceSlice []ItemInstance

type DBItemInstance struct {
	bun.BaseModel `bun:"table:item_instance,alias:item_instance"`
	Guid int `bun:"guid"`
	ItemEntry int32 `bun:"itemEntry"`
	OwnerGuid int `bun:"owner_guid"`
	CreatorGuid int `bun:"creatorGuid"`
	GiftCreatorGuid int `bun:"giftCreatorGuid"`
	Count int `bun:"count"`
	Duration int `bun:"duration"`
	Charges string `bun:"charges"`
	Flags int32 `bun:"flags"`
	Enchantments string `bun:"enchantments"`
	RandomPropertyId int16 `bun:"randomPropertyId"`
	Durability int16 `bun:"durability"`
	PlayedTime int `bun:"playedTime"`
	Text string `bun:"text"`
}

type DBItemInstanceSlice []DBItemInstance

func (entry *ItemInstance) ToDB() *DBItemInstance {
	if entry == nil {
		return nil
	}
	return &DBItemInstance{
		Guid: entry.Guid,
		ItemEntry: entry.ItemEntry,
		OwnerGuid: entry.OwnerGuid,
		CreatorGuid: entry.CreatorGuid,
		GiftCreatorGuid: entry.GiftCreatorGuid,
		Count: entry.Count,
		Duration: entry.Duration,
		Charges: entry.Charges,
		Flags: entry.Flags,
		Enchantments: entry.Enchantments,
		RandomPropertyId: entry.RandomPropertyId,
		Durability: entry.Durability,
		PlayedTime: entry.PlayedTime,
		Text: entry.Text,
	}
}

func (entry *DBItemInstance) ToWeb() *ItemInstance {
	if entry == nil {
		return nil
	}
	return &ItemInstance{
		Guid: entry.Guid,
		ItemEntry: entry.ItemEntry,
		OwnerGuid: entry.OwnerGuid,
		CreatorGuid: entry.CreatorGuid,
		GiftCreatorGuid: entry.GiftCreatorGuid,
		Count: entry.Count,
		Duration: entry.Duration,
		Charges: entry.Charges,
		Flags: entry.Flags,
		Enchantments: entry.Enchantments,
		RandomPropertyId: entry.RandomPropertyId,
		Durability: entry.Durability,
		PlayedTime: entry.PlayedTime,
		Text: entry.Text,
	}
}

func (data ItemInstanceSlice) ToDB() DBItemInstanceSlice {
	result := make(DBItemInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemInstanceSlice) ToWeb() ItemInstanceSlice {
	result := make(ItemInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
