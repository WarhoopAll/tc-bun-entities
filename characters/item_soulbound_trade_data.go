package model

import "github.com/uptrace/bun"

type ItemSoulboundTradeData struct {
	ItemGuid int `json:"itemguid,omitempty"`
	AllowedPlayers string `json:"allowedplayers,omitempty"`
}

type ItemSoulboundTradeDataSlice []ItemSoulboundTradeData

type DBItemSoulboundTradeData struct {
	bun.BaseModel `bun:"table:item_soulbound_trade_data,alias:item_soulbound_trade_data"`
	ItemGuid int `bun:"itemGuid"`
	AllowedPlayers string `bun:"allowedPlayers"`
}

type DBItemSoulboundTradeDataSlice []DBItemSoulboundTradeData

func (entry *ItemSoulboundTradeData) ToDB() *DBItemSoulboundTradeData {
	if entry == nil {
		return nil
	}
	return &DBItemSoulboundTradeData{
		ItemGuid: entry.ItemGuid,
		AllowedPlayers: entry.AllowedPlayers,
	}
}

func (entry *DBItemSoulboundTradeData) ToWeb() *ItemSoulboundTradeData {
	if entry == nil {
		return nil
	}
	return &ItemSoulboundTradeData{
		ItemGuid: entry.ItemGuid,
		AllowedPlayers: entry.AllowedPlayers,
	}
}

func (data ItemSoulboundTradeDataSlice) ToDB() DBItemSoulboundTradeDataSlice {
	result := make(DBItemSoulboundTradeDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemSoulboundTradeDataSlice) ToWeb() ItemSoulboundTradeDataSlice {
	result := make(ItemSoulboundTradeDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
