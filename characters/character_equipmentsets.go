package model

import "github.com/uptrace/bun"

type CharacterEquipmentsets struct {
	Guid int `json:"guid,omitempty"`
	Setguid int `json:"setguid,omitempty"`
	Setindex int8 `json:"setindex,omitempty"`
	Name string `json:"name,omitempty"`
	Iconname string `json:"iconname,omitempty"`
	IgnoreMask int `json:"ignore_mask,omitempty"`
	Item0 int `json:"item0,omitempty"`
	Item1 int `json:"item1,omitempty"`
	Item2 int `json:"item2,omitempty"`
	Item3 int `json:"item3,omitempty"`
	Item4 int `json:"item4,omitempty"`
	Item5 int `json:"item5,omitempty"`
	Item6 int `json:"item6,omitempty"`
	Item7 int `json:"item7,omitempty"`
	Item8 int `json:"item8,omitempty"`
	Item9 int `json:"item9,omitempty"`
	Item10 int `json:"item10,omitempty"`
	Item11 int `json:"item11,omitempty"`
	Item12 int `json:"item12,omitempty"`
	Item13 int `json:"item13,omitempty"`
	Item14 int `json:"item14,omitempty"`
	Item15 int `json:"item15,omitempty"`
	Item16 int `json:"item16,omitempty"`
	Item17 int `json:"item17,omitempty"`
	Item18 int `json:"item18,omitempty"`
}

type CharacterEquipmentsetsSlice []CharacterEquipmentsets

type DBCharacterEquipmentsets struct {
	bun.BaseModel `bun:"table:character_equipmentsets,alias:character_equipmentsets"`
	Guid int `bun:"guid"`
	Setguid int `bun:"setguid"`
	Setindex int8 `bun:"setindex"`
	Name string `bun:"name"`
	Iconname string `bun:"iconname"`
	IgnoreMask int `bun:"ignore_mask"`
	Item0 int `bun:"item0"`
	Item1 int `bun:"item1"`
	Item2 int `bun:"item2"`
	Item3 int `bun:"item3"`
	Item4 int `bun:"item4"`
	Item5 int `bun:"item5"`
	Item6 int `bun:"item6"`
	Item7 int `bun:"item7"`
	Item8 int `bun:"item8"`
	Item9 int `bun:"item9"`
	Item10 int `bun:"item10"`
	Item11 int `bun:"item11"`
	Item12 int `bun:"item12"`
	Item13 int `bun:"item13"`
	Item14 int `bun:"item14"`
	Item15 int `bun:"item15"`
	Item16 int `bun:"item16"`
	Item17 int `bun:"item17"`
	Item18 int `bun:"item18"`
}

type DBCharacterEquipmentsetsSlice []DBCharacterEquipmentsets

func (entry *CharacterEquipmentsets) ToDB() *DBCharacterEquipmentsets {
	if entry == nil {
		return nil
	}
	return &DBCharacterEquipmentsets{
		Guid: entry.Guid,
		Setguid: entry.Setguid,
		Setindex: entry.Setindex,
		Name: entry.Name,
		Iconname: entry.Iconname,
		IgnoreMask: entry.IgnoreMask,
		Item0: entry.Item0,
		Item1: entry.Item1,
		Item2: entry.Item2,
		Item3: entry.Item3,
		Item4: entry.Item4,
		Item5: entry.Item5,
		Item6: entry.Item6,
		Item7: entry.Item7,
		Item8: entry.Item8,
		Item9: entry.Item9,
		Item10: entry.Item10,
		Item11: entry.Item11,
		Item12: entry.Item12,
		Item13: entry.Item13,
		Item14: entry.Item14,
		Item15: entry.Item15,
		Item16: entry.Item16,
		Item17: entry.Item17,
		Item18: entry.Item18,
	}
}

func (entry *DBCharacterEquipmentsets) ToWeb() *CharacterEquipmentsets {
	if entry == nil {
		return nil
	}
	return &CharacterEquipmentsets{
		Guid: entry.Guid,
		Setguid: entry.Setguid,
		Setindex: entry.Setindex,
		Name: entry.Name,
		Iconname: entry.Iconname,
		IgnoreMask: entry.IgnoreMask,
		Item0: entry.Item0,
		Item1: entry.Item1,
		Item2: entry.Item2,
		Item3: entry.Item3,
		Item4: entry.Item4,
		Item5: entry.Item5,
		Item6: entry.Item6,
		Item7: entry.Item7,
		Item8: entry.Item8,
		Item9: entry.Item9,
		Item10: entry.Item10,
		Item11: entry.Item11,
		Item12: entry.Item12,
		Item13: entry.Item13,
		Item14: entry.Item14,
		Item15: entry.Item15,
		Item16: entry.Item16,
		Item17: entry.Item17,
		Item18: entry.Item18,
	}
}

func (data CharacterEquipmentsetsSlice) ToDB() DBCharacterEquipmentsetsSlice {
	result := make(DBCharacterEquipmentsetsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterEquipmentsetsSlice) ToWeb() CharacterEquipmentsetsSlice {
	result := make(CharacterEquipmentsetsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
