package model

import "github.com/uptrace/bun"

type AccountTutorial struct {
	AccountId int `json:"accountid,omitempty"`
	Tut0 int `json:"tut0,omitempty"`
	Tut1 int `json:"tut1,omitempty"`
	Tut2 int `json:"tut2,omitempty"`
	Tut3 int `json:"tut3,omitempty"`
	Tut4 int `json:"tut4,omitempty"`
	Tut5 int `json:"tut5,omitempty"`
	Tut6 int `json:"tut6,omitempty"`
	Tut7 int `json:"tut7,omitempty"`
}

type AccountTutorialSlice []AccountTutorial

type DBAccountTutorial struct {
	bun.BaseModel `bun:"table:account_tutorial,alias:account_tutorial"`
	AccountId int `bun:"accountId"`
	Tut0 int `bun:"tut0"`
	Tut1 int `bun:"tut1"`
	Tut2 int `bun:"tut2"`
	Tut3 int `bun:"tut3"`
	Tut4 int `bun:"tut4"`
	Tut5 int `bun:"tut5"`
	Tut6 int `bun:"tut6"`
	Tut7 int `bun:"tut7"`
}

type DBAccountTutorialSlice []DBAccountTutorial

func (entry *AccountTutorial) ToDB() *DBAccountTutorial {
	if entry == nil {
		return nil
	}
	return &DBAccountTutorial{
		AccountId: entry.AccountId,
		Tut0: entry.Tut0,
		Tut1: entry.Tut1,
		Tut2: entry.Tut2,
		Tut3: entry.Tut3,
		Tut4: entry.Tut4,
		Tut5: entry.Tut5,
		Tut6: entry.Tut6,
		Tut7: entry.Tut7,
	}
}

func (entry *DBAccountTutorial) ToWeb() *AccountTutorial {
	if entry == nil {
		return nil
	}
	return &AccountTutorial{
		AccountId: entry.AccountId,
		Tut0: entry.Tut0,
		Tut1: entry.Tut1,
		Tut2: entry.Tut2,
		Tut3: entry.Tut3,
		Tut4: entry.Tut4,
		Tut5: entry.Tut5,
		Tut6: entry.Tut6,
		Tut7: entry.Tut7,
	}
}

func (data AccountTutorialSlice) ToDB() DBAccountTutorialSlice {
	result := make(DBAccountTutorialSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAccountTutorialSlice) ToWeb() AccountTutorialSlice {
	result := make(AccountTutorialSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
