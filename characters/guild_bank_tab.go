package model

import "github.com/uptrace/bun"

type GuildBankTab struct {
	Guildid int `json:"guildid,omitempty"`
	TabId int8 `json:"tabid,omitempty"`
	TabName string `json:"tabname,omitempty"`
	TabIcon string `json:"tabicon,omitempty"`
	TabText string `json:"tabtext,omitempty"`
}

type GuildBankTabSlice []GuildBankTab

type DBGuildBankTab struct {
	bun.BaseModel `bun:"table:guild_bank_tab,alias:guild_bank_tab"`
	Guildid int `bun:"guildid"`
	TabId int8 `bun:"TabId"`
	TabName string `bun:"TabName"`
	TabIcon string `bun:"TabIcon"`
	TabText string `bun:"TabText"`
}

type DBGuildBankTabSlice []DBGuildBankTab

func (entry *GuildBankTab) ToDB() *DBGuildBankTab {
	if entry == nil {
		return nil
	}
	return &DBGuildBankTab{
		Guildid: entry.Guildid,
		TabId: entry.TabId,
		TabName: entry.TabName,
		TabIcon: entry.TabIcon,
		TabText: entry.TabText,
	}
}

func (entry *DBGuildBankTab) ToWeb() *GuildBankTab {
	if entry == nil {
		return nil
	}
	return &GuildBankTab{
		Guildid: entry.Guildid,
		TabId: entry.TabId,
		TabName: entry.TabName,
		TabIcon: entry.TabIcon,
		TabText: entry.TabText,
	}
}

func (data GuildBankTabSlice) ToDB() DBGuildBankTabSlice {
	result := make(DBGuildBankTabSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildBankTabSlice) ToWeb() GuildBankTabSlice {
	result := make(GuildBankTabSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
