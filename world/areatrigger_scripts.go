package model

import "github.com/uptrace/bun"

type AreatriggerScripts struct {
	Entry int `json:"entry,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
}

type AreatriggerScriptsSlice []AreatriggerScripts

type DBAreatriggerScripts struct {
	bun.BaseModel `bun:"table:areatrigger_scripts,alias:areatrigger_scripts"`
	Entry int `bun:"entry"`
	ScriptName string `bun:"ScriptName"`
}

type DBAreatriggerScriptsSlice []DBAreatriggerScripts

func (entry *AreatriggerScripts) ToDB() *DBAreatriggerScripts {
	if entry == nil {
		return nil
	}
	return &DBAreatriggerScripts{
		Entry: entry.Entry,
		ScriptName: entry.ScriptName,
	}
}

func (entry *DBAreatriggerScripts) ToWeb() *AreatriggerScripts {
	if entry == nil {
		return nil
	}
	return &AreatriggerScripts{
		Entry: entry.Entry,
		ScriptName: entry.ScriptName,
	}
}

func (data AreatriggerScriptsSlice) ToDB() DBAreatriggerScriptsSlice {
	result := make(DBAreatriggerScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAreatriggerScriptsSlice) ToWeb() AreatriggerScriptsSlice {
	result := make(AreatriggerScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
