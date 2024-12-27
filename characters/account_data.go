package model

import "github.com/uptrace/bun"

type AccountData struct {
	AccountId int `json:"accountid,omitempty"`
	Type int8 `json:"type,omitempty"`
	Time int `json:"time,omitempty"`
	Data []byte `json:"data,omitempty"`
}

type AccountDataSlice []AccountData

type DBAccountData struct {
	bun.BaseModel `bun:"table:account_data,alias:account_data"`
	AccountId int `bun:"accountId"`
	Type int8 `bun:"type"`
	Time int `bun:"time"`
	Data []byte `bun:"data"`
}

type DBAccountDataSlice []DBAccountData

func (entry *AccountData) ToDB() *DBAccountData {
	if entry == nil {
		return nil
	}
	return &DBAccountData{
		AccountId: entry.AccountId,
		Type: entry.Type,
		Time: entry.Time,
		Data: entry.Data,
	}
}

func (entry *DBAccountData) ToWeb() *AccountData {
	if entry == nil {
		return nil
	}
	return &AccountData{
		AccountId: entry.AccountId,
		Type: entry.Type,
		Time: entry.Time,
		Data: entry.Data,
	}
}

func (data AccountDataSlice) ToDB() DBAccountDataSlice {
	result := make(DBAccountDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAccountDataSlice) ToWeb() AccountDataSlice {
	result := make(AccountDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
