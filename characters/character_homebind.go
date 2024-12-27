package model

import "github.com/uptrace/bun"

type CharacterHomebind struct {
	Guid int `json:"guid,omitempty"`
	MapId int16 `json:"mapid,omitempty"`
	ZoneId int16 `json:"zoneid,omitempty"`
	PosX float64 `json:"posx,omitempty"`
	PosY float64 `json:"posy,omitempty"`
	PosZ float64 `json:"posz,omitempty"`
}

type CharacterHomebindSlice []CharacterHomebind

type DBCharacterHomebind struct {
	bun.BaseModel `bun:"table:character_homebind,alias:character_homebind"`
	Guid int `bun:"guid"`
	MapId int16 `bun:"mapId"`
	ZoneId int16 `bun:"zoneId"`
	PosX float64 `bun:"posX"`
	PosY float64 `bun:"posY"`
	PosZ float64 `bun:"posZ"`
}

type DBCharacterHomebindSlice []DBCharacterHomebind

func (entry *CharacterHomebind) ToDB() *DBCharacterHomebind {
	if entry == nil {
		return nil
	}
	return &DBCharacterHomebind{
		Guid: entry.Guid,
		MapId: entry.MapId,
		ZoneId: entry.ZoneId,
		PosX: entry.PosX,
		PosY: entry.PosY,
		PosZ: entry.PosZ,
	}
}

func (entry *DBCharacterHomebind) ToWeb() *CharacterHomebind {
	if entry == nil {
		return nil
	}
	return &CharacterHomebind{
		Guid: entry.Guid,
		MapId: entry.MapId,
		ZoneId: entry.ZoneId,
		PosX: entry.PosX,
		PosY: entry.PosY,
		PosZ: entry.PosZ,
	}
}

func (data CharacterHomebindSlice) ToDB() DBCharacterHomebindSlice {
	result := make(DBCharacterHomebindSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterHomebindSlice) ToWeb() CharacterHomebindSlice {
	result := make(CharacterHomebindSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
