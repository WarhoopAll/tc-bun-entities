package model

import "github.com/uptrace/bun"

type GraveyardZone struct {
	ID int `json:"id,omitempty"`
	GhostZone int `json:"ghostzone,omitempty"`
	Faction int16 `json:"faction,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type GraveyardZoneSlice []GraveyardZone

type DBGraveyardZone struct {
	bun.BaseModel `bun:"table:graveyard_zone,alias:graveyard_zone"`
	ID int `bun:"ID"`
	GhostZone int `bun:"GhostZone"`
	Faction int16 `bun:"Faction"`
	Comment string `bun:"Comment"`
}

type DBGraveyardZoneSlice []DBGraveyardZone

func (entry *GraveyardZone) ToDB() *DBGraveyardZone {
	if entry == nil {
		return nil
	}
	return &DBGraveyardZone{
		ID: entry.ID,
		GhostZone: entry.GhostZone,
		Faction: entry.Faction,
		Comment: entry.Comment,
	}
}

func (entry *DBGraveyardZone) ToWeb() *GraveyardZone {
	if entry == nil {
		return nil
	}
	return &GraveyardZone{
		ID: entry.ID,
		GhostZone: entry.GhostZone,
		Faction: entry.Faction,
		Comment: entry.Comment,
	}
}

func (data GraveyardZoneSlice) ToDB() DBGraveyardZoneSlice {
	result := make(DBGraveyardZoneSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGraveyardZoneSlice) ToWeb() GraveyardZoneSlice {
	result := make(GraveyardZoneSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
