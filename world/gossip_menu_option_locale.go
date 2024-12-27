package model

import "github.com/uptrace/bun"

type GossipMenuOptionLocale struct {
	MenuID int16 `json:"menuid,omitempty"`
	OptionID int16 `json:"optionid,omitempty"`
	Locale string `json:"locale,omitempty"`
	OptionText string `json:"optiontext,omitempty"`
	BoxText string `json:"boxtext,omitempty"`
}

type GossipMenuOptionLocaleSlice []GossipMenuOptionLocale

type DBGossipMenuOptionLocale struct {
	bun.BaseModel `bun:"table:gossip_menu_option_locale,alias:gossip_menu_option_locale"`
	MenuID int16 `bun:"MenuID"`
	OptionID int16 `bun:"OptionID"`
	Locale string `bun:"Locale"`
	OptionText string `bun:"OptionText"`
	BoxText string `bun:"BoxText"`
}

type DBGossipMenuOptionLocaleSlice []DBGossipMenuOptionLocale

func (entry *GossipMenuOptionLocale) ToDB() *DBGossipMenuOptionLocale {
	if entry == nil {
		return nil
	}
	return &DBGossipMenuOptionLocale{
		MenuID: entry.MenuID,
		OptionID: entry.OptionID,
		Locale: entry.Locale,
		OptionText: entry.OptionText,
		BoxText: entry.BoxText,
	}
}

func (entry *DBGossipMenuOptionLocale) ToWeb() *GossipMenuOptionLocale {
	if entry == nil {
		return nil
	}
	return &GossipMenuOptionLocale{
		MenuID: entry.MenuID,
		OptionID: entry.OptionID,
		Locale: entry.Locale,
		OptionText: entry.OptionText,
		BoxText: entry.BoxText,
	}
}

func (data GossipMenuOptionLocaleSlice) ToDB() DBGossipMenuOptionLocaleSlice {
	result := make(DBGossipMenuOptionLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGossipMenuOptionLocaleSlice) ToWeb() GossipMenuOptionLocaleSlice {
	result := make(GossipMenuOptionLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
