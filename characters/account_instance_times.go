package model

import "github.com/uptrace/bun"

type AccountInstanceTimes struct {
	AccountId int `json:"accountid,omitempty"`
	InstanceId int `json:"instanceid,omitempty"`
	ReleaseTime int `json:"releasetime,omitempty"`
}

type AccountInstanceTimesSlice []AccountInstanceTimes

type DBAccountInstanceTimes struct {
	bun.BaseModel `bun:"table:account_instance_times,alias:account_instance_times"`
	AccountId int `bun:"accountId"`
	InstanceId int `bun:"instanceId"`
	ReleaseTime int `bun:"releaseTime"`
}

type DBAccountInstanceTimesSlice []DBAccountInstanceTimes

func (entry *AccountInstanceTimes) ToDB() *DBAccountInstanceTimes {
	if entry == nil {
		return nil
	}
	return &DBAccountInstanceTimes{
		AccountId: entry.AccountId,
		InstanceId: entry.InstanceId,
		ReleaseTime: entry.ReleaseTime,
	}
}

func (entry *DBAccountInstanceTimes) ToWeb() *AccountInstanceTimes {
	if entry == nil {
		return nil
	}
	return &AccountInstanceTimes{
		AccountId: entry.AccountId,
		InstanceId: entry.InstanceId,
		ReleaseTime: entry.ReleaseTime,
	}
}

func (data AccountInstanceTimesSlice) ToDB() DBAccountInstanceTimesSlice {
	result := make(DBAccountInstanceTimesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAccountInstanceTimesSlice) ToWeb() AccountInstanceTimesSlice {
	result := make(AccountInstanceTimesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
