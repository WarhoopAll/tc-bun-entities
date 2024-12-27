package model

import "github.com/uptrace/bun"

type AccountBanned struct {
	Id int `json:"id,omitempty"`
	Bandate int `json:"bandate,omitempty"`
	Unbandate int `json:"unbandate,omitempty"`
	Bannedby string `json:"bannedby,omitempty"`
	Banreason string `json:"banreason,omitempty"`
	Active int8 `json:"active,omitempty"`
}

type AccountBannedSlice []AccountBanned

type DBAccountBanned struct {
	bun.BaseModel `bun:"table:account_banned,alias:account_banned"`
	Id int `bun:"id"`
	Bandate int `bun:"bandate"`
	Unbandate int `bun:"unbandate"`
	Bannedby string `bun:"bannedby"`
	Banreason string `bun:"banreason"`
	Active int8 `bun:"active"`
}

type DBAccountBannedSlice []DBAccountBanned

func (entry *AccountBanned) ToDB() *DBAccountBanned {
	if entry == nil {
		return nil
	}
	return &DBAccountBanned{
		Id: entry.Id,
		Bandate: entry.Bandate,
		Unbandate: entry.Unbandate,
		Bannedby: entry.Bannedby,
		Banreason: entry.Banreason,
		Active: entry.Active,
	}
}

func (entry *DBAccountBanned) ToWeb() *AccountBanned {
	if entry == nil {
		return nil
	}
	return &AccountBanned{
		Id: entry.Id,
		Bandate: entry.Bandate,
		Unbandate: entry.Unbandate,
		Bannedby: entry.Bannedby,
		Banreason: entry.Banreason,
		Active: entry.Active,
	}
}

func (data AccountBannedSlice) ToDB() DBAccountBannedSlice {
	result := make(DBAccountBannedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAccountBannedSlice) ToWeb() AccountBannedSlice {
	result := make(AccountBannedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
