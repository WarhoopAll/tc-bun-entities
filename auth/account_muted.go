package model

import "github.com/uptrace/bun"

type AccountMuted struct {
	Guid int `json:"guid,omitempty"`
	Mutedate int `json:"mutedate,omitempty"`
	Mutetime int `json:"mutetime,omitempty"`
	Mutedby string `json:"mutedby,omitempty"`
	Mutereason string `json:"mutereason,omitempty"`
}

type AccountMutedSlice []AccountMuted

type DBAccountMuted struct {
	bun.BaseModel `bun:"table:account_muted,alias:account_muted"`
	Guid int `bun:"guid"`
	Mutedate int `bun:"mutedate"`
	Mutetime int `bun:"mutetime"`
	Mutedby string `bun:"mutedby"`
	Mutereason string `bun:"mutereason"`
}

type DBAccountMutedSlice []DBAccountMuted

func (entry *AccountMuted) ToDB() *DBAccountMuted {
	if entry == nil {
		return nil
	}
	return &DBAccountMuted{
		Guid: entry.Guid,
		Mutedate: entry.Mutedate,
		Mutetime: entry.Mutetime,
		Mutedby: entry.Mutedby,
		Mutereason: entry.Mutereason,
	}
}

func (entry *DBAccountMuted) ToWeb() *AccountMuted {
	if entry == nil {
		return nil
	}
	return &AccountMuted{
		Guid: entry.Guid,
		Mutedate: entry.Mutedate,
		Mutetime: entry.Mutetime,
		Mutedby: entry.Mutedby,
		Mutereason: entry.Mutereason,
	}
}

func (data AccountMutedSlice) ToDB() DBAccountMutedSlice {
	result := make(DBAccountMutedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAccountMutedSlice) ToWeb() AccountMutedSlice {
	result := make(AccountMutedSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
