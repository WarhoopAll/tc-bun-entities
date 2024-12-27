package model

import "github.com/uptrace/bun"

type LfgDungeonTemplate struct {
	DungeonId int `json:"dungeonid,omitempty"`
	Name string `json:"name,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type LfgDungeonTemplateSlice []LfgDungeonTemplate

type DBLfgDungeonTemplate struct {
	bun.BaseModel `bun:"table:lfg_dungeon_template,alias:lfg_dungeon_template"`
	DungeonId int `bun:"dungeonId"`
	Name string `bun:"name"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Orientation float64 `bun:"orientation"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBLfgDungeonTemplateSlice []DBLfgDungeonTemplate

func (entry *LfgDungeonTemplate) ToDB() *DBLfgDungeonTemplate {
	if entry == nil {
		return nil
	}
	return &DBLfgDungeonTemplate{
		DungeonId: entry.DungeonId,
		Name: entry.Name,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBLfgDungeonTemplate) ToWeb() *LfgDungeonTemplate {
	if entry == nil {
		return nil
	}
	return &LfgDungeonTemplate{
		DungeonId: entry.DungeonId,
		Name: entry.Name,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data LfgDungeonTemplateSlice) ToDB() DBLfgDungeonTemplateSlice {
	result := make(DBLfgDungeonTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBLfgDungeonTemplateSlice) ToWeb() LfgDungeonTemplateSlice {
	result := make(LfgDungeonTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
