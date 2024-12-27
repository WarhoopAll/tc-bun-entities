package model

import "github.com/uptrace/bun"

type GuildBankEventlog struct {
	Guildid int `json:"guildid,omitempty"`
	LogGuid int `json:"logguid,omitempty"`
	TabId int8 `json:"tabid,omitempty"`
	EventType int8 `json:"eventtype,omitempty"`
	PlayerGuid int `json:"playerguid,omitempty"`
	ItemOrMoney int `json:"itemormoney,omitempty"`
	ItemStackCount int16 `json:"itemstackcount,omitempty"`
	DestTabId int8 `json:"desttabid,omitempty"`
	TimeStamp int `json:"timestamp,omitempty"`
}

type GuildBankEventlogSlice []GuildBankEventlog

type DBGuildBankEventlog struct {
	bun.BaseModel `bun:"table:guild_bank_eventlog,alias:guild_bank_eventlog"`
	Guildid int `bun:"guildid"`
	LogGuid int `bun:"LogGuid"`
	TabId int8 `bun:"TabId"`
	EventType int8 `bun:"EventType"`
	PlayerGuid int `bun:"PlayerGuid"`
	ItemOrMoney int `bun:"ItemOrMoney"`
	ItemStackCount int16 `bun:"ItemStackCount"`
	DestTabId int8 `bun:"DestTabId"`
	TimeStamp int `bun:"TimeStamp"`
}

type DBGuildBankEventlogSlice []DBGuildBankEventlog

func (entry *GuildBankEventlog) ToDB() *DBGuildBankEventlog {
	if entry == nil {
		return nil
	}
	return &DBGuildBankEventlog{
		Guildid: entry.Guildid,
		LogGuid: entry.LogGuid,
		TabId: entry.TabId,
		EventType: entry.EventType,
		PlayerGuid: entry.PlayerGuid,
		ItemOrMoney: entry.ItemOrMoney,
		ItemStackCount: entry.ItemStackCount,
		DestTabId: entry.DestTabId,
		TimeStamp: entry.TimeStamp,
	}
}

func (entry *DBGuildBankEventlog) ToWeb() *GuildBankEventlog {
	if entry == nil {
		return nil
	}
	return &GuildBankEventlog{
		Guildid: entry.Guildid,
		LogGuid: entry.LogGuid,
		TabId: entry.TabId,
		EventType: entry.EventType,
		PlayerGuid: entry.PlayerGuid,
		ItemOrMoney: entry.ItemOrMoney,
		ItemStackCount: entry.ItemStackCount,
		DestTabId: entry.DestTabId,
		TimeStamp: entry.TimeStamp,
	}
}

func (data GuildBankEventlogSlice) ToDB() DBGuildBankEventlogSlice {
	result := make(DBGuildBankEventlogSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildBankEventlogSlice) ToWeb() GuildBankEventlogSlice {
	result := make(GuildBankEventlogSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
