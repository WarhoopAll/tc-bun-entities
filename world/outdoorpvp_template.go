package model

import "github.com/uptrace/bun"

type OutdoorpvpTemplate struct {
	TypeId int8 `json:"typeid,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type OutdoorpvpTemplateSlice []OutdoorpvpTemplate

type DBOutdoorpvpTemplate struct {
	bun.BaseModel `bun:"table:outdoorpvp_template,alias:outdoorpvp_template"`
	TypeId int8 `bun:"TypeId"`
	ScriptName string `bun:"ScriptName"`
	Comment string `bun:"comment"`
}

type DBOutdoorpvpTemplateSlice []DBOutdoorpvpTemplate

func (entry *OutdoorpvpTemplate) ToDB() *DBOutdoorpvpTemplate {
	if entry == nil {
		return nil
	}
	return &DBOutdoorpvpTemplate{
		TypeId: entry.TypeId,
		ScriptName: entry.ScriptName,
		Comment: entry.Comment,
	}
}

func (entry *DBOutdoorpvpTemplate) ToWeb() *OutdoorpvpTemplate {
	if entry == nil {
		return nil
	}
	return &OutdoorpvpTemplate{
		TypeId: entry.TypeId,
		ScriptName: entry.ScriptName,
		Comment: entry.Comment,
	}
}

func (data OutdoorpvpTemplateSlice) ToDB() DBOutdoorpvpTemplateSlice {
	result := make(DBOutdoorpvpTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBOutdoorpvpTemplateSlice) ToWeb() OutdoorpvpTemplateSlice {
	result := make(OutdoorpvpTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
