package model

import "github.com/uptrace/bun"

type Playercreateinfo struct {
	Race int8 `json:"race,omitempty"`
	Class int8 `json:"class,omitempty"`
	Map int16 `json:"map,omitempty"`
	Zone int `json:"zone,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
}

type PlayercreateinfoSlice []Playercreateinfo

type DBPlayercreateinfo struct {
	bun.BaseModel `bun:"table:playercreateinfo,alias:playercreateinfo"`
	Race int8 `bun:"race"`
	Class int8 `bun:"class"`
	Map int16 `bun:"map"`
	Zone int `bun:"zone"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Orientation float64 `bun:"orientation"`
}

type DBPlayercreateinfoSlice []DBPlayercreateinfo

func (entry *Playercreateinfo) ToDB() *DBPlayercreateinfo {
	if entry == nil {
		return nil
	}
	return &DBPlayercreateinfo{
		Race: entry.Race,
		Class: entry.Class,
		Map: entry.Map,
		Zone: entry.Zone,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
	}
}

func (entry *DBPlayercreateinfo) ToWeb() *Playercreateinfo {
	if entry == nil {
		return nil
	}
	return &Playercreateinfo{
		Race: entry.Race,
		Class: entry.Class,
		Map: entry.Map,
		Zone: entry.Zone,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
	}
}

func (data PlayercreateinfoSlice) ToDB() DBPlayercreateinfoSlice {
	result := make(DBPlayercreateinfoSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayercreateinfoSlice) ToWeb() PlayercreateinfoSlice {
	result := make(PlayercreateinfoSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
