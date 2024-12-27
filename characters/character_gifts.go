package model

import "github.com/uptrace/bun"

type CharacterGifts struct {
	Guid int `json:"guid,omitempty"`
	ItemGuid int `json:"item_guid,omitempty"`
	Entry int `json:"entry,omitempty"`
	Flags int `json:"flags,omitempty"`
}

type CharacterGiftsSlice []CharacterGifts

type DBCharacterGifts struct {
	bun.BaseModel `bun:"table:character_gifts,alias:character_gifts"`
	Guid int `bun:"guid"`
	ItemGuid int `bun:"item_guid"`
	Entry int `bun:"entry"`
	Flags int `bun:"flags"`
}

type DBCharacterGiftsSlice []DBCharacterGifts

func (entry *CharacterGifts) ToDB() *DBCharacterGifts {
	if entry == nil {
		return nil
	}
	return &DBCharacterGifts{
		Guid: entry.Guid,
		ItemGuid: entry.ItemGuid,
		Entry: entry.Entry,
		Flags: entry.Flags,
	}
}

func (entry *DBCharacterGifts) ToWeb() *CharacterGifts {
	if entry == nil {
		return nil
	}
	return &CharacterGifts{
		Guid: entry.Guid,
		ItemGuid: entry.ItemGuid,
		Entry: entry.Entry,
		Flags: entry.Flags,
	}
}

func (data CharacterGiftsSlice) ToDB() DBCharacterGiftsSlice {
	result := make(DBCharacterGiftsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterGiftsSlice) ToWeb() CharacterGiftsSlice {
	result := make(CharacterGiftsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
