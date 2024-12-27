package model

import "github.com/uptrace/bun"

type SpellScriptNames struct {
	SpellId int `json:"spell_id,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
}

type SpellScriptNamesSlice []SpellScriptNames

type DBSpellScriptNames struct {
	bun.BaseModel `bun:"table:spell_script_names,alias:spell_script_names"`
	SpellId int `bun:"spell_id"`
	ScriptName string `bun:"ScriptName"`
}

type DBSpellScriptNamesSlice []DBSpellScriptNames

func (entry *SpellScriptNames) ToDB() *DBSpellScriptNames {
	if entry == nil {
		return nil
	}
	return &DBSpellScriptNames{
		SpellId: entry.SpellId,
		ScriptName: entry.ScriptName,
	}
}

func (entry *DBSpellScriptNames) ToWeb() *SpellScriptNames {
	if entry == nil {
		return nil
	}
	return &SpellScriptNames{
		SpellId: entry.SpellId,
		ScriptName: entry.ScriptName,
	}
}

func (data SpellScriptNamesSlice) ToDB() DBSpellScriptNamesSlice {
	result := make(DBSpellScriptNamesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellScriptNamesSlice) ToWeb() SpellScriptNamesSlice {
	result := make(SpellScriptNamesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
