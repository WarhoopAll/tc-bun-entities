package model

import "github.com/uptrace/bun"

type GuildMember struct {
	Guildid int `json:"guildid,omitempty"`
	Guid int `json:"guid,omitempty"`
	Rank int8 `json:"rank,omitempty"`
	Pnote string `json:"pnote,omitempty"`
	Offnote string `json:"offnote,omitempty"`
}

type GuildMemberSlice []GuildMember

type DBGuildMember struct {
	bun.BaseModel `bun:"table:guild_member,alias:guild_member"`
	Guildid int `bun:"guildid"`
	Guid int `bun:"guid"`
	Rank int8 `bun:"rank"`
	Pnote string `bun:"pnote"`
	Offnote string `bun:"offnote"`
}

type DBGuildMemberSlice []DBGuildMember

func (entry *GuildMember) ToDB() *DBGuildMember {
	if entry == nil {
		return nil
	}
	return &DBGuildMember{
		Guildid: entry.Guildid,
		Guid: entry.Guid,
		Rank: entry.Rank,
		Pnote: entry.Pnote,
		Offnote: entry.Offnote,
	}
}

func (entry *DBGuildMember) ToWeb() *GuildMember {
	if entry == nil {
		return nil
	}
	return &GuildMember{
		Guildid: entry.Guildid,
		Guid: entry.Guid,
		Rank: entry.Rank,
		Pnote: entry.Pnote,
		Offnote: entry.Offnote,
	}
}

func (data GuildMemberSlice) ToDB() DBGuildMemberSlice {
	result := make(DBGuildMemberSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildMemberSlice) ToWeb() GuildMemberSlice {
	result := make(GuildMemberSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
