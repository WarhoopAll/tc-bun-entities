package model

import "github.com/uptrace/bun"

type CreatureEquipTemplate struct {
	CreatureID int `json:"creatureid,omitempty"`
	ID int8 `json:"id,omitempty"`
	ItemID1 int `json:"itemid1,omitempty"`
	ItemID2 int `json:"itemid2,omitempty"`
	ItemID3 int `json:"itemid3,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type CreatureEquipTemplateSlice []CreatureEquipTemplate

type DBCreatureEquipTemplate struct {
	bun.BaseModel `bun:"table:creature_equip_template,alias:creature_equip_template"`
	CreatureID int `bun:"CreatureID"`
	ID int8 `bun:"ID"`
	ItemID1 int `bun:"ItemID1"`
	ItemID2 int `bun:"ItemID2"`
	ItemID3 int `bun:"ItemID3"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBCreatureEquipTemplateSlice []DBCreatureEquipTemplate

func (entry *CreatureEquipTemplate) ToDB() *DBCreatureEquipTemplate {
	if entry == nil {
		return nil
	}
	return &DBCreatureEquipTemplate{
		CreatureID: entry.CreatureID,
		ID: entry.ID,
		ItemID1: entry.ItemID1,
		ItemID2: entry.ItemID2,
		ItemID3: entry.ItemID3,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBCreatureEquipTemplate) ToWeb() *CreatureEquipTemplate {
	if entry == nil {
		return nil
	}
	return &CreatureEquipTemplate{
		CreatureID: entry.CreatureID,
		ID: entry.ID,
		ItemID1: entry.ItemID1,
		ItemID2: entry.ItemID2,
		ItemID3: entry.ItemID3,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data CreatureEquipTemplateSlice) ToDB() DBCreatureEquipTemplateSlice {
	result := make(DBCreatureEquipTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureEquipTemplateSlice) ToWeb() CreatureEquipTemplateSlice {
	result := make(CreatureEquipTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
