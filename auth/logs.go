package model

import "github.com/uptrace/bun"

type Logs struct {
	Time int `json:"time,omitempty"`
	Realm int `json:"realm,omitempty"`
	Type string `json:"type,omitempty"`
	Level int8 `json:"level,omitempty"`
	String string `json:"string,omitempty"`
}

type LogsSlice []Logs

type DBLogs struct {
	bun.BaseModel `bun:"table:logs,alias:logs"`
	Time int `bun:"time"`
	Realm int `bun:"realm"`
	Type string `bun:"type"`
	Level int8 `bun:"level"`
	String string `bun:"string"`
}

type DBLogsSlice []DBLogs

func (entry *Logs) ToDB() *DBLogs {
	if entry == nil {
		return nil
	}
	return &DBLogs{
		Time: entry.Time,
		Realm: entry.Realm,
		Type: entry.Type,
		Level: entry.Level,
		String: entry.String,
	}
}

func (entry *DBLogs) ToWeb() *Logs {
	if entry == nil {
		return nil
	}
	return &Logs{
		Time: entry.Time,
		Realm: entry.Realm,
		Type: entry.Type,
		Level: entry.Level,
		String: entry.String,
	}
}

func (data LogsSlice) ToDB() DBLogsSlice {
	result := make(DBLogsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBLogsSlice) ToWeb() LogsSlice {
	result := make(LogsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
