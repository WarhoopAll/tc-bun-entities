package model

import "github.com/uptrace/bun"

type GameobjectAddon struct {
	Guid int `json:"guid,omitempty"`
	ParentRotation0 float64 `json:"parent_rotation0,omitempty"`
	ParentRotation1 float64 `json:"parent_rotation1,omitempty"`
	ParentRotation2 float64 `json:"parent_rotation2,omitempty"`
	ParentRotation3 float64 `json:"parent_rotation3,omitempty"`
	InvisibilityType int8 `json:"invisibilitytype,omitempty"`
	InvisibilityValue int `json:"invisibilityvalue,omitempty"`
}

type GameobjectAddonSlice []GameobjectAddon

type DBGameobjectAddon struct {
	bun.BaseModel `bun:"table:gameobject_addon,alias:gameobject_addon"`
	Guid int `bun:"guid"`
	ParentRotation0 float64 `bun:"parent_rotation0"`
	ParentRotation1 float64 `bun:"parent_rotation1"`
	ParentRotation2 float64 `bun:"parent_rotation2"`
	ParentRotation3 float64 `bun:"parent_rotation3"`
	InvisibilityType int8 `bun:"invisibilityType"`
	InvisibilityValue int `bun:"invisibilityValue"`
}

type DBGameobjectAddonSlice []DBGameobjectAddon

func (entry *GameobjectAddon) ToDB() *DBGameobjectAddon {
	if entry == nil {
		return nil
	}
	return &DBGameobjectAddon{
		Guid: entry.Guid,
		ParentRotation0: entry.ParentRotation0,
		ParentRotation1: entry.ParentRotation1,
		ParentRotation2: entry.ParentRotation2,
		ParentRotation3: entry.ParentRotation3,
		InvisibilityType: entry.InvisibilityType,
		InvisibilityValue: entry.InvisibilityValue,
	}
}

func (entry *DBGameobjectAddon) ToWeb() *GameobjectAddon {
	if entry == nil {
		return nil
	}
	return &GameobjectAddon{
		Guid: entry.Guid,
		ParentRotation0: entry.ParentRotation0,
		ParentRotation1: entry.ParentRotation1,
		ParentRotation2: entry.ParentRotation2,
		ParentRotation3: entry.ParentRotation3,
		InvisibilityType: entry.InvisibilityType,
		InvisibilityValue: entry.InvisibilityValue,
	}
}

func (data GameobjectAddonSlice) ToDB() DBGameobjectAddonSlice {
	result := make(DBGameobjectAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectAddonSlice) ToWeb() GameobjectAddonSlice {
	result := make(GameobjectAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
