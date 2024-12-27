package model

import "github.com/uptrace/bun"

type LogsIpActions struct {
	Id int `json:"id,omitempty"`
	AccountId int `json:"account_id,omitempty"`
	CharacterGuid int `json:"character_guid,omitempty"`
	RealmId int `json:"realm_id,omitempty"`
	Type int8 `json:"type,omitempty"`
	Ip string `json:"ip,omitempty"`
	Systemnote string `json:"systemnote,omitempty"`
	Unixtime int `json:"unixtime,omitempty"`
	Time time.Time `json:"time,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type LogsIpActionsSlice []LogsIpActions

type DBLogsIpActions struct {
	bun.BaseModel `bun:"table:logs_ip_actions,alias:logs_ip_actions"`
	Id int `bun:"id"`
	AccountId int `bun:"account_id"`
	CharacterGuid int `bun:"character_guid"`
	RealmId int `bun:"realm_id"`
	Type int8 `bun:"type"`
	Ip string `bun:"ip"`
	Systemnote string `bun:"systemnote"`
	Unixtime int `bun:"unixtime"`
	Time time.Time `bun:"time"`
	Comment string `bun:"comment"`
}

type DBLogsIpActionsSlice []DBLogsIpActions

func (entry *LogsIpActions) ToDB() *DBLogsIpActions {
	if entry == nil {
		return nil
	}
	return &DBLogsIpActions{
		Id: entry.Id,
		AccountId: entry.AccountId,
		CharacterGuid: entry.CharacterGuid,
		RealmId: entry.RealmId,
		Type: entry.Type,
		Ip: entry.Ip,
		Systemnote: entry.Systemnote,
		Unixtime: entry.Unixtime,
		Time: entry.Time,
		Comment: entry.Comment,
	}
}

func (entry *DBLogsIpActions) ToWeb() *LogsIpActions {
	if entry == nil {
		return nil
	}
	return &LogsIpActions{
		Id: entry.Id,
		AccountId: entry.AccountId,
		CharacterGuid: entry.CharacterGuid,
		RealmId: entry.RealmId,
		Type: entry.Type,
		Ip: entry.Ip,
		Systemnote: entry.Systemnote,
		Unixtime: entry.Unixtime,
		Time: entry.Time,
		Comment: entry.Comment,
	}
}

func (data LogsIpActionsSlice) ToDB() DBLogsIpActionsSlice {
	result := make(DBLogsIpActionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBLogsIpActionsSlice) ToWeb() LogsIpActionsSlice {
	result := make(LogsIpActionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
