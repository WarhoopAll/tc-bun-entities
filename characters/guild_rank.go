package model

import "github.com/uptrace/bun"

type GuildRank struct {
	Guildid int `json:"guildid,omitempty"`
	Rid int8 `json:"rid,omitempty"`
	Rname string `json:"rname,omitempty"`
	Rights int32 `json:"rights,omitempty"`
	BankMoneyPerDay int `json:"bankmoneyperday,omitempty"`
}

type GuildRankSlice []GuildRank

type DBGuildRank struct {
	bun.BaseModel `bun:"table:guild_rank,alias:guild_rank"`
	Guildid int `bun:"guildid"`
	Rid int8 `bun:"rid"`
	Rname string `bun:"rname"`
	Rights int32 `bun:"rights"`
	BankMoneyPerDay int `bun:"BankMoneyPerDay"`
}

type DBGuildRankSlice []DBGuildRank

func (entry *GuildRank) ToDB() *DBGuildRank {
	if entry == nil {
		return nil
	}
	return &DBGuildRank{
		Guildid: entry.Guildid,
		Rid: entry.Rid,
		Rname: entry.Rname,
		Rights: entry.Rights,
		BankMoneyPerDay: entry.BankMoneyPerDay,
	}
}

func (entry *DBGuildRank) ToWeb() *GuildRank {
	if entry == nil {
		return nil
	}
	return &GuildRank{
		Guildid: entry.Guildid,
		Rid: entry.Rid,
		Rname: entry.Rname,
		Rights: entry.Rights,
		BankMoneyPerDay: entry.BankMoneyPerDay,
	}
}

func (data GuildRankSlice) ToDB() DBGuildRankSlice {
	result := make(DBGuildRankSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildRankSlice) ToWeb() GuildRankSlice {
	result := make(GuildRankSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
