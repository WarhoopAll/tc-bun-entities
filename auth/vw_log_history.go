package model

import "github.com/uptrace/bun"

type VwLogHistory struct {
	First Logged time.Time `json:"first logged,omitempty"`
	Last Logged time.Time `json:"last logged,omitempty"`
	Occurrences int `json:"occurrences,omitempty"`
	Realm string `json:"realm,omitempty"`
	Type string `json:"type,omitempty"`
	Level int8 `json:"level,omitempty"`
	String string `json:"string,omitempty"`
}

type VwLogHistorySlice []VwLogHistory

type DBVwLogHistory struct {
	bun.BaseModel `bun:"table:vw_log_history,alias:vw_log_history"`
	First Logged time.Time `bun:"First Logged"`
	Last Logged time.Time `bun:"Last Logged"`
	Occurrences int `bun:"Occurrences"`
	Realm string `bun:"Realm"`
	Type string `bun:"type"`
	Level int8 `bun:"level"`
	String string `bun:"string"`
}

type DBVwLogHistorySlice []DBVwLogHistory

func (entry *VwLogHistory) ToDB() *DBVwLogHistory {
	if entry == nil {
		return nil
	}
	return &DBVwLogHistory{
		First Logged: entry.First Logged,
		Last Logged: entry.Last Logged,
		Occurrences: entry.Occurrences,
		Realm: entry.Realm,
		Type: entry.Type,
		Level: entry.Level,
		String: entry.String,
	}
}

func (entry *DBVwLogHistory) ToWeb() *VwLogHistory {
	if entry == nil {
		return nil
	}
	return &VwLogHistory{
		First Logged: entry.First Logged,
		Last Logged: entry.Last Logged,
		Occurrences: entry.Occurrences,
		Realm: entry.Realm,
		Type: entry.Type,
		Level: entry.Level,
		String: entry.String,
	}
}

func (data VwLogHistorySlice) ToDB() DBVwLogHistorySlice {
	result := make(DBVwLogHistorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVwLogHistorySlice) ToWeb() VwLogHistorySlice {
	result := make(VwLogHistorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
