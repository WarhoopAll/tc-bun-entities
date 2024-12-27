package model

import "github.com/uptrace/bun"

type GuildBankRight struct {
	Guildid int `json:"guildid,omitempty"`
	TabId int8 `json:"tabid,omitempty"`
	Rid int8 `json:"rid,omitempty"`
	Gbright int8 `json:"gbright,omitempty"`
	SlotPerDay int `json:"slotperday,omitempty"`
}

type GuildBankRightSlice []GuildBankRight

type DBGuildBankRight struct {
	bun.BaseModel `bun:"table:guild_bank_right,alias:guild_bank_right"`
	Guildid int `bun:"guildid"`
	TabId int8 `bun:"TabId"`
	Rid int8 `bun:"rid"`
	Gbright int8 `bun:"gbright"`
	SlotPerDay int `bun:"SlotPerDay"`
}

type DBGuildBankRightSlice []DBGuildBankRight

func (entry *GuildBankRight) ToDB() *DBGuildBankRight {
	if entry == nil {
		return nil
	}
	return &DBGuildBankRight{
		Guildid: entry.Guildid,
		TabId: entry.TabId,
		Rid: entry.Rid,
		Gbright: entry.Gbright,
		SlotPerDay: entry.SlotPerDay,
	}
}

func (entry *DBGuildBankRight) ToWeb() *GuildBankRight {
	if entry == nil {
		return nil
	}
	return &GuildBankRight{
		Guildid: entry.Guildid,
		TabId: entry.TabId,
		Rid: entry.Rid,
		Gbright: entry.Gbright,
		SlotPerDay: entry.SlotPerDay,
	}
}

func (data GuildBankRightSlice) ToDB() DBGuildBankRightSlice {
	result := make(DBGuildBankRightSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildBankRightSlice) ToWeb() GuildBankRightSlice {
	result := make(GuildBankRightSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
