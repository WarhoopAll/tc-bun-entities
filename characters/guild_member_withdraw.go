package model

import "github.com/uptrace/bun"

type GuildMemberWithdraw struct {
	Guid int `json:"guid,omitempty"`
	Tab0 int `json:"tab0,omitempty"`
	Tab1 int `json:"tab1,omitempty"`
	Tab2 int `json:"tab2,omitempty"`
	Tab3 int `json:"tab3,omitempty"`
	Tab4 int `json:"tab4,omitempty"`
	Tab5 int `json:"tab5,omitempty"`
	Money int `json:"money,omitempty"`
}

type GuildMemberWithdrawSlice []GuildMemberWithdraw

type DBGuildMemberWithdraw struct {
	bun.BaseModel `bun:"table:guild_member_withdraw,alias:guild_member_withdraw"`
	Guid int `bun:"guid"`
	Tab0 int `bun:"tab0"`
	Tab1 int `bun:"tab1"`
	Tab2 int `bun:"tab2"`
	Tab3 int `bun:"tab3"`
	Tab4 int `bun:"tab4"`
	Tab5 int `bun:"tab5"`
	Money int `bun:"money"`
}

type DBGuildMemberWithdrawSlice []DBGuildMemberWithdraw

func (entry *GuildMemberWithdraw) ToDB() *DBGuildMemberWithdraw {
	if entry == nil {
		return nil
	}
	return &DBGuildMemberWithdraw{
		Guid: entry.Guid,
		Tab0: entry.Tab0,
		Tab1: entry.Tab1,
		Tab2: entry.Tab2,
		Tab3: entry.Tab3,
		Tab4: entry.Tab4,
		Tab5: entry.Tab5,
		Money: entry.Money,
	}
}

func (entry *DBGuildMemberWithdraw) ToWeb() *GuildMemberWithdraw {
	if entry == nil {
		return nil
	}
	return &GuildMemberWithdraw{
		Guid: entry.Guid,
		Tab0: entry.Tab0,
		Tab1: entry.Tab1,
		Tab2: entry.Tab2,
		Tab3: entry.Tab3,
		Tab4: entry.Tab4,
		Tab5: entry.Tab5,
		Money: entry.Money,
	}
}

func (data GuildMemberWithdrawSlice) ToDB() DBGuildMemberWithdrawSlice {
	result := make(DBGuildMemberWithdrawSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildMemberWithdrawSlice) ToWeb() GuildMemberWithdrawSlice {
	result := make(GuildMemberWithdrawSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
