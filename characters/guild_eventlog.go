package model

import "github.com/uptrace/bun"

type GuildEventlog struct {
	Guildid int `json:"guildid,omitempty"`
	LogGuid int `json:"logguid,omitempty"`
	EventType int8 `json:"eventtype,omitempty"`
	PlayerGuid1 int `json:"playerguid1,omitempty"`
	PlayerGuid2 int `json:"playerguid2,omitempty"`
	NewRank int8 `json:"newrank,omitempty"`
	TimeStamp int `json:"timestamp,omitempty"`
}

type GuildEventlogSlice []GuildEventlog

type DBGuildEventlog struct {
	bun.BaseModel `bun:"table:guild_eventlog,alias:guild_eventlog"`
	Guildid int `bun:"guildid"`
	LogGuid int `bun:"LogGuid"`
	EventType int8 `bun:"EventType"`
	PlayerGuid1 int `bun:"PlayerGuid1"`
	PlayerGuid2 int `bun:"PlayerGuid2"`
	NewRank int8 `bun:"NewRank"`
	TimeStamp int `bun:"TimeStamp"`
}

type DBGuildEventlogSlice []DBGuildEventlog

func (entry *GuildEventlog) ToDB() *DBGuildEventlog {
	if entry == nil {
		return nil
	}
	return &DBGuildEventlog{
		Guildid: entry.Guildid,
		LogGuid: entry.LogGuid,
		EventType: entry.EventType,
		PlayerGuid1: entry.PlayerGuid1,
		PlayerGuid2: entry.PlayerGuid2,
		NewRank: entry.NewRank,
		TimeStamp: entry.TimeStamp,
	}
}

func (entry *DBGuildEventlog) ToWeb() *GuildEventlog {
	if entry == nil {
		return nil
	}
	return &GuildEventlog{
		Guildid: entry.Guildid,
		LogGuid: entry.LogGuid,
		EventType: entry.EventType,
		PlayerGuid1: entry.PlayerGuid1,
		PlayerGuid2: entry.PlayerGuid2,
		NewRank: entry.NewRank,
		TimeStamp: entry.TimeStamp,
	}
}

func (data GuildEventlogSlice) ToDB() DBGuildEventlogSlice {
	result := make(DBGuildEventlogSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildEventlogSlice) ToWeb() GuildEventlogSlice {
	result := make(GuildEventlogSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
