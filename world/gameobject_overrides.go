package model

import "github.com/uptrace/bun"

type GameobjectOverrides struct {
	SpawnId int `json:"spawnid,omitempty"`
	Faction int16 `json:"faction,omitempty"`
	Flags int `json:"flags,omitempty"`
}

type GameobjectOverridesSlice []GameobjectOverrides

type DBGameobjectOverrides struct {
	bun.BaseModel `bun:"table:gameobject_overrides,alias:gameobject_overrides"`
	SpawnId int `bun:"spawnId"`
	Faction int16 `bun:"faction"`
	Flags int `bun:"flags"`
}

type DBGameobjectOverridesSlice []DBGameobjectOverrides

func (entry *GameobjectOverrides) ToDB() *DBGameobjectOverrides {
	if entry == nil {
		return nil
	}
	return &DBGameobjectOverrides{
		SpawnId: entry.SpawnId,
		Faction: entry.Faction,
		Flags: entry.Flags,
	}
}

func (entry *DBGameobjectOverrides) ToWeb() *GameobjectOverrides {
	if entry == nil {
		return nil
	}
	return &GameobjectOverrides{
		SpawnId: entry.SpawnId,
		Faction: entry.Faction,
		Flags: entry.Flags,
	}
}

func (data GameobjectOverridesSlice) ToDB() DBGameobjectOverridesSlice {
	result := make(DBGameobjectOverridesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectOverridesSlice) ToWeb() GameobjectOverridesSlice {
	result := make(GameobjectOverridesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
