package model

import "github.com/uptrace/bun"

type GuildBankItem struct {
	Guildid int `json:"guildid,omitempty"`
	TabId int8 `json:"tabid,omitempty"`
	SlotId int8 `json:"slotid,omitempty"`
	ItemGuid int `json:"item_guid,omitempty"`
}

type GuildBankItemSlice []GuildBankItem

type DBGuildBankItem struct {
	bun.BaseModel `bun:"table:guild_bank_item,alias:guild_bank_item"`
	Guildid int `bun:"guildid"`
	TabId int8 `bun:"TabId"`
	SlotId int8 `bun:"SlotId"`
	ItemGuid int `bun:"item_guid"`
}

type DBGuildBankItemSlice []DBGuildBankItem

func (entry *GuildBankItem) ToDB() *DBGuildBankItem {
	if entry == nil {
		return nil
	}
	return &DBGuildBankItem{
		Guildid: entry.Guildid,
		TabId: entry.TabId,
		SlotId: entry.SlotId,
		ItemGuid: entry.ItemGuid,
	}
}

func (entry *DBGuildBankItem) ToWeb() *GuildBankItem {
	if entry == nil {
		return nil
	}
	return &GuildBankItem{
		Guildid: entry.Guildid,
		TabId: entry.TabId,
		SlotId: entry.SlotId,
		ItemGuid: entry.ItemGuid,
	}
}

func (data GuildBankItemSlice) ToDB() DBGuildBankItemSlice {
	result := make(DBGuildBankItemSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGuildBankItemSlice) ToWeb() GuildBankItemSlice {
	result := make(GuildBankItemSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
