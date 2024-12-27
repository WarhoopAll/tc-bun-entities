package model

import "github.com/uptrace/bun"

type BattlefieldTemplate struct {
	TypeId int8 `json:"typeid,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type BattlefieldTemplateSlice []BattlefieldTemplate

type DBBattlefieldTemplate struct {
	bun.BaseModel `bun:"table:battlefield_template,alias:battlefield_template"`
	TypeId int8 `bun:"TypeId"`
	ScriptName string `bun:"ScriptName"`
	Comment string `bun:"comment"`
}

type DBBattlefieldTemplateSlice []DBBattlefieldTemplate

func (entry *BattlefieldTemplate) ToDB() *DBBattlefieldTemplate {
	if entry == nil {
		return nil
	}
	return &DBBattlefieldTemplate{
		TypeId: entry.TypeId,
		ScriptName: entry.ScriptName,
		Comment: entry.Comment,
	}
}

func (entry *DBBattlefieldTemplate) ToWeb() *BattlefieldTemplate {
	if entry == nil {
		return nil
	}
	return &BattlefieldTemplate{
		TypeId: entry.TypeId,
		ScriptName: entry.ScriptName,
		Comment: entry.Comment,
	}
}

func (data BattlefieldTemplateSlice) ToDB() DBBattlefieldTemplateSlice {
	result := make(DBBattlefieldTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBattlefieldTemplateSlice) ToWeb() BattlefieldTemplateSlice {
	result := make(BattlefieldTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
