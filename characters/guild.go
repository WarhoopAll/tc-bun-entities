package model

import "github.com/uptrace/bun"

type Guild struct {
	Guildid int `json:"guildid,omitempty"`
	Name string `json:"name,omitempty"`
	Leaderguid int `json:"leaderguid,omitempty"`
	EmblemStyle int8 `json:"emblemstyle,omitempty"`
	EmblemColor int8 `json:"emblemcolor,omitempty"`
	BorderStyle int8 `json:"borderstyle,omitempty"`
	BorderColor int8 `json:"bordercolor,omitempty"`
	BackgroundColor int8 `json:"backgroundcolor,omitempty"`
	Info string `json:"info,omitempty"`
	Motd string `json:"motd,omitempty"`
	Createdate int `json:"createdate,omitempty"`
	BankMoney int `json:"bankmoney,omitempty"`
}

type GuildSlice []Guild

type DBGuild struct {
	bun.BaseModel `bun:"table:guild,alias:guild"`
	Guildid int `bun:"guildid"`
	Name string `bun:"name"`
	Leaderguid int `bun:"leaderguid"`
	EmblemStyle int8 `bun:"EmblemStyle"`
	EmblemColor int8 `bun:"EmblemColor"`
	BorderStyle int8 `bun:"BorderStyle"`
	BorderColor int8 `bun:"BorderColor"`
	BackgroundColor int8 `bun:"BackgroundColor"`
	Info string `bun:"info"`
	Motd string `bun:"motd"`
	Createdate int `bun:"createdate"`
	BankMoney int `bun:"BankMoney"`
}

type DBGuildSlice []DBGuild

func (entry *Guild) ToDB() *DBGuild {
	if entry == nil {
		return nil
	}
	return &DBGuild{
		Guildid: entry.Guildid,
		Name: entry.Name,
		Leaderguid: entry.Leaderguid,
		EmblemStyle: entry.EmblemStyle,
		EmblemColor: entry.EmblemColor,
		BorderStyle: entry.BorderStyle,
		BorderColor: entry.BorderColor,
		BackgroundColor: entry.BackgroundColor,
		Info: entry.Info,
		Motd: entry.Motd,
		Createdate: entry.Createdate,
		BankMoney: entry.BankMoney,
	}
}

func (entry *DBGuild) ToWeb() *Guild {
	if entry == nil {
		return nil
	}
	return &Guild{
		Guildid: entry.Guildid,
		Name: entry.Name,
		Leaderguid: entry.Leaderguid,
		EmblemStyle: entry.EmblemStyle,
		EmblemColor: entry.EmblemColor,
		BorderStyle: entry.BorderStyle,
		BorderColor: entry.BorderColor,
		BackgroundColor: entry.BackgroundColor,
		Info: entry.Info,
		Motd: entry.Motd,
		Createdate: entry.Createdate,
		BankMoney: entry.BankMoney,
	}
}

func (data GuildSlice) ToDB() DBGuildSlice {
	result := make(DBGuildSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildSlice) ToWeb() GuildSlice {
	result := make(GuildSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
