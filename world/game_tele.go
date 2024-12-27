package model

import "github.com/uptrace/bun"

type GameTele struct {
	Id int `json:"id,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	Map int16 `json:"map,omitempty"`
	Name string `json:"name,omitempty"`
}

type GameTeleSlice []GameTele

type DBGameTele struct {
	bun.BaseModel `bun:"table:game_tele,alias:game_tele"`
	Id int `bun:"id"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Orientation float64 `bun:"orientation"`
	Map int16 `bun:"map"`
	Name string `bun:"name"`
}

type DBGameTeleSlice []DBGameTele

func (entry *GameTele) ToDB() *DBGameTele {
	if entry == nil {
		return nil
	}
	return &DBGameTele{
		Id: entry.Id,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Map: entry.Map,
		Name: entry.Name,
	}
}

func (entry *DBGameTele) ToWeb() *GameTele {
	if entry == nil {
		return nil
	}
	return &GameTele{
		Id: entry.Id,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Map: entry.Map,
		Name: entry.Name,
	}
}

func (data GameTeleSlice) ToDB() DBGameTeleSlice {
	result := make(DBGameTeleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameTeleSlice) ToWeb() GameTeleSlice {
	result := make(GameTeleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
