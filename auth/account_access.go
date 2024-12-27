package model

import "github.com/uptrace/bun"

type AccountAccess struct {
	AccountID int `json:"accountid,omitempty"`
	SecurityLevel int8 `json:"securitylevel,omitempty"`
	RealmID int `json:"realmid,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type AccountAccessSlice []AccountAccess

type DBAccountAccess struct {
	bun.BaseModel `bun:"table:account_access,alias:account_access"`
	AccountID int `bun:"AccountID"`
	SecurityLevel int8 `bun:"SecurityLevel"`
	RealmID int `bun:"RealmID"`
	Comment string `bun:"Comment"`
}

type DBAccountAccessSlice []DBAccountAccess

func (entry *AccountAccess) ToDB() *DBAccountAccess {
	if entry == nil {
		return nil
	}
	return &DBAccountAccess{
		AccountID: entry.AccountID,
		SecurityLevel: entry.SecurityLevel,
		RealmID: entry.RealmID,
		Comment: entry.Comment,
	}
}

func (entry *DBAccountAccess) ToWeb() *AccountAccess {
	if entry == nil {
		return nil
	}
	return &AccountAccess{
		AccountID: entry.AccountID,
		SecurityLevel: entry.SecurityLevel,
		RealmID: entry.RealmID,
		Comment: entry.Comment,
	}
}

func (data AccountAccessSlice) ToDB() DBAccountAccessSlice {
	result := make(DBAccountAccessSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAccountAccessSlice) ToWeb() AccountAccessSlice {
	result := make(AccountAccessSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
