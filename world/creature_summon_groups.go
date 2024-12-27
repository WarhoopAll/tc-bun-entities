package model

import "github.com/uptrace/bun"

type CreatureSummonGroups struct {
	SummonerId int `json:"summonerid,omitempty"`
	SummonerType int8 `json:"summonertype,omitempty"`
	GroupId int8 `json:"groupid,omitempty"`
	Entry int `json:"entry,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	SummonType int8 `json:"summontype,omitempty"`
	SummonTime int `json:"summontime,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type CreatureSummonGroupsSlice []CreatureSummonGroups

type DBCreatureSummonGroups struct {
	bun.BaseModel `bun:"table:creature_summon_groups,alias:creature_summon_groups"`
	SummonerId int `bun:"summonerId"`
	SummonerType int8 `bun:"summonerType"`
	GroupId int8 `bun:"groupId"`
	Entry int `bun:"entry"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Orientation float64 `bun:"orientation"`
	SummonType int8 `bun:"summonType"`
	SummonTime int `bun:"summonTime"`
	Comment string `bun:"Comment"`
}

type DBCreatureSummonGroupsSlice []DBCreatureSummonGroups

func (entry *CreatureSummonGroups) ToDB() *DBCreatureSummonGroups {
	if entry == nil {
		return nil
	}
	return &DBCreatureSummonGroups{
		SummonerId: entry.SummonerId,
		SummonerType: entry.SummonerType,
		GroupId: entry.GroupId,
		Entry: entry.Entry,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		SummonType: entry.SummonType,
		SummonTime: entry.SummonTime,
		Comment: entry.Comment,
	}
}

func (entry *DBCreatureSummonGroups) ToWeb() *CreatureSummonGroups {
	if entry == nil {
		return nil
	}
	return &CreatureSummonGroups{
		SummonerId: entry.SummonerId,
		SummonerType: entry.SummonerType,
		GroupId: entry.GroupId,
		Entry: entry.Entry,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		SummonType: entry.SummonType,
		SummonTime: entry.SummonTime,
		Comment: entry.Comment,
	}
}

func (data CreatureSummonGroupsSlice) ToDB() DBCreatureSummonGroupsSlice {
	result := make(DBCreatureSummonGroupsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureSummonGroupsSlice) ToWeb() CreatureSummonGroupsSlice {
	result := make(CreatureSummonGroupsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
