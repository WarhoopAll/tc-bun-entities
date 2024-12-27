package model

import "github.com/uptrace/bun"

type GossipMenuOption struct {
	MenuID int16 `json:"menuid,omitempty"`
	OptionID int16 `json:"optionid,omitempty"`
	OptionIcon int `json:"optionicon,omitempty"`
	OptionText string `json:"optiontext,omitempty"`
	OptionBroadcastTextID int `json:"optionbroadcasttextid,omitempty"`
	OptionType int8 `json:"optiontype,omitempty"`
	OptionNpcFlag int `json:"optionnpcflag,omitempty"`
	ActionMenuID int `json:"actionmenuid,omitempty"`
	ActionPoiID int `json:"actionpoiid,omitempty"`
	BoxCoded int8 `json:"boxcoded,omitempty"`
	BoxMoney int `json:"boxmoney,omitempty"`
	BoxText string `json:"boxtext,omitempty"`
	BoxBroadcastTextID int `json:"boxbroadcasttextid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type GossipMenuOptionSlice []GossipMenuOption

type DBGossipMenuOption struct {
	bun.BaseModel `bun:"table:gossip_menu_option,alias:gossip_menu_option"`
	MenuID int16 `bun:"MenuID"`
	OptionID int16 `bun:"OptionID"`
	OptionIcon int `bun:"OptionIcon"`
	OptionText string `bun:"OptionText"`
	OptionBroadcastTextID int `bun:"OptionBroadcastTextID"`
	OptionType int8 `bun:"OptionType"`
	OptionNpcFlag int `bun:"OptionNpcFlag"`
	ActionMenuID int `bun:"ActionMenuID"`
	ActionPoiID int `bun:"ActionPoiID"`
	BoxCoded int8 `bun:"BoxCoded"`
	BoxMoney int `bun:"BoxMoney"`
	BoxText string `bun:"BoxText"`
	BoxBroadcastTextID int `bun:"BoxBroadcastTextID"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBGossipMenuOptionSlice []DBGossipMenuOption

func (entry *GossipMenuOption) ToDB() *DBGossipMenuOption {
	if entry == nil {
		return nil
	}
	return &DBGossipMenuOption{
		MenuID: entry.MenuID,
		OptionID: entry.OptionID,
		OptionIcon: entry.OptionIcon,
		OptionText: entry.OptionText,
		OptionBroadcastTextID: entry.OptionBroadcastTextID,
		OptionType: entry.OptionType,
		OptionNpcFlag: entry.OptionNpcFlag,
		ActionMenuID: entry.ActionMenuID,
		ActionPoiID: entry.ActionPoiID,
		BoxCoded: entry.BoxCoded,
		BoxMoney: entry.BoxMoney,
		BoxText: entry.BoxText,
		BoxBroadcastTextID: entry.BoxBroadcastTextID,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBGossipMenuOption) ToWeb() *GossipMenuOption {
	if entry == nil {
		return nil
	}
	return &GossipMenuOption{
		MenuID: entry.MenuID,
		OptionID: entry.OptionID,
		OptionIcon: entry.OptionIcon,
		OptionText: entry.OptionText,
		OptionBroadcastTextID: entry.OptionBroadcastTextID,
		OptionType: entry.OptionType,
		OptionNpcFlag: entry.OptionNpcFlag,
		ActionMenuID: entry.ActionMenuID,
		ActionPoiID: entry.ActionPoiID,
		BoxCoded: entry.BoxCoded,
		BoxMoney: entry.BoxMoney,
		BoxText: entry.BoxText,
		BoxBroadcastTextID: entry.BoxBroadcastTextID,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data GossipMenuOptionSlice) ToDB() DBGossipMenuOptionSlice {
	result := make(DBGossipMenuOptionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGossipMenuOptionSlice) ToWeb() GossipMenuOptionSlice {
	result := make(GossipMenuOptionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
