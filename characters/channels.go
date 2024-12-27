package model

import "github.com/uptrace/bun"

type Channels struct {
	Name string `json:"name,omitempty"`
	Team int `json:"team,omitempty"`
	Announce int8 `json:"announce,omitempty"`
	Ownership int8 `json:"ownership,omitempty"`
	Password string `json:"password,omitempty"`
	BannedList string `json:"bannedlist,omitempty"`
	LastUsed int `json:"lastused,omitempty"`
}

type ChannelsSlice []Channels

type DBChannels struct {
	bun.BaseModel `bun:"table:channels,alias:channels"`
	Name string `bun:"name"`
	Team int `bun:"team"`
	Announce int8 `bun:"announce"`
	Ownership int8 `bun:"ownership"`
	Password string `bun:"password"`
	BannedList string `bun:"bannedList"`
	LastUsed int `bun:"lastUsed"`
}

type DBChannelsSlice []DBChannels

func (entry *Channels) ToDB() *DBChannels {
	if entry == nil {
		return nil
	}
	return &DBChannels{
		Name: entry.Name,
		Team: entry.Team,
		Announce: entry.Announce,
		Ownership: entry.Ownership,
		Password: entry.Password,
		BannedList: entry.BannedList,
		LastUsed: entry.LastUsed,
	}
}

func (entry *DBChannels) ToWeb() *Channels {
	if entry == nil {
		return nil
	}
	return &Channels{
		Name: entry.Name,
		Team: entry.Team,
		Announce: entry.Announce,
		Ownership: entry.Ownership,
		Password: entry.Password,
		BannedList: entry.BannedList,
		LastUsed: entry.LastUsed,
	}
}

func (data ChannelsSlice) ToDB() DBChannelsSlice {
	result := make(DBChannelsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBChannelsSlice) ToWeb() ChannelsSlice {
	result := make(ChannelsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
