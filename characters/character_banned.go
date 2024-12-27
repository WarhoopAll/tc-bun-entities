package model

import "github.com/uptrace/bun"

type CharacterBanned struct {
	Guid int `json:"guid,omitempty"`
	Bandate int `json:"bandate,omitempty"`
	Unbandate int `json:"unbandate,omitempty"`
	Bannedby string `json:"bannedby,omitempty"`
	Banreason string `json:"banreason,omitempty"`
	Active int8 `json:"active,omitempty"`
}

type CharacterBannedSlice []CharacterBanned

type DBCharacterBanned struct {
	bun.BaseModel `bun:"table:character_banned,alias:character_banned"`
	Guid int `bun:"guid"`
	Bandate int `bun:"bandate"`
	Unbandate int `bun:"unbandate"`
	Bannedby string `bun:"bannedby"`
	Banreason string `bun:"banreason"`
	Active int8 `bun:"active"`
}

type DBCharacterBannedSlice []DBCharacterBanned

func (entry *CharacterBanned) ToDB() *DBCharacterBanned {
	if entry == nil {
		return nil
	}
	return &DBCharacterBanned{
		Guid: entry.Guid,
		Bandate: entry.Bandate,
		Unbandate: entry.Unbandate,
		Bannedby: entry.Bannedby,
		Banreason: entry.Banreason,
		Active: entry.Active,
	}
}

func (entry *DBCharacterBanned) ToWeb() *CharacterBanned {
	if entry == nil {
		return nil
	}
	return &CharacterBanned{
		Guid: entry.Guid,
		Bandate: entry.Bandate,
		Unbandate: entry.Unbandate,
		Bannedby: entry.Bannedby,
		Banreason: entry.Banreason,
		Active: entry.Active,
	}
}

func (data CharacterBannedSlice) ToDB() DBCharacterBannedSlice {
	result := make(DBCharacterBannedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterBannedSlice) ToWeb() CharacterBannedSlice {
	result := make(CharacterBannedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
