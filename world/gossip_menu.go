package model

import "github.com/uptrace/bun"

type GossipMenu struct {
	MenuID int16 `json:"menuid,omitempty"`
	TextID int `json:"textid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type GossipMenuSlice []GossipMenu

type DBGossipMenu struct {
	bun.BaseModel `bun:"table:gossip_menu,alias:gossip_menu"`
	MenuID int16 `bun:"MenuID"`
	TextID int `bun:"TextID"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBGossipMenuSlice []DBGossipMenu

func (entry *GossipMenu) ToDB() *DBGossipMenu {
	if entry == nil {
		return nil
	}
	return &DBGossipMenu{
		MenuID: entry.MenuID,
		TextID: entry.TextID,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBGossipMenu) ToWeb() *GossipMenu {
	if entry == nil {
		return nil
	}
	return &GossipMenu{
		MenuID: entry.MenuID,
		TextID: entry.TextID,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data GossipMenuSlice) ToDB() DBGossipMenuSlice {
	result := make(DBGossipMenuSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGossipMenuSlice) ToWeb() GossipMenuSlice {
	result := make(GossipMenuSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
