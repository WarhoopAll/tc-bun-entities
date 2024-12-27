package model

import "github.com/uptrace/bun"

type IpBanned struct {
	Ip string `json:"ip,omitempty"`
	Bandate int `json:"bandate,omitempty"`
	Unbandate int `json:"unbandate,omitempty"`
	Bannedby string `json:"bannedby,omitempty"`
	Banreason string `json:"banreason,omitempty"`
}

type IpBannedSlice []IpBanned

type DBIpBanned struct {
	bun.BaseModel `bun:"table:ip_banned,alias:ip_banned"`
	Ip string `bun:"ip"`
	Bandate int `bun:"bandate"`
	Unbandate int `bun:"unbandate"`
	Bannedby string `bun:"bannedby"`
	Banreason string `bun:"banreason"`
}

type DBIpBannedSlice []DBIpBanned

func (entry *IpBanned) ToDB() *DBIpBanned {
	if entry == nil {
		return nil
	}
	return &DBIpBanned{
		Ip: entry.Ip,
		Bandate: entry.Bandate,
		Unbandate: entry.Unbandate,
		Bannedby: entry.Bannedby,
		Banreason: entry.Banreason,
	}
}

func (entry *DBIpBanned) ToWeb() *IpBanned {
	if entry == nil {
		return nil
	}
	return &IpBanned{
		Ip: entry.Ip,
		Bandate: entry.Bandate,
		Unbandate: entry.Unbandate,
		Bannedby: entry.Bannedby,
		Banreason: entry.Banreason,
	}
}

func (data IpBannedSlice) ToDB() DBIpBannedSlice {
	result := make(DBIpBannedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBIpBannedSlice) ToWeb() IpBannedSlice {
	result := make(IpBannedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
