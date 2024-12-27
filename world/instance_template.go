package model

import "github.com/uptrace/bun"

type InstanceTemplate struct {
	Map int16 `json:"map,omitempty"`
	Parent int16 `json:"parent,omitempty"`
	Script string `json:"script,omitempty"`
	AllowMount int8 `json:"allowmount,omitempty"`
}

type InstanceTemplateSlice []InstanceTemplate

type DBInstanceTemplate struct {
	bun.BaseModel `bun:"table:instance_template,alias:instance_template"`
	Map int16 `bun:"map"`
	Parent int16 `bun:"parent"`
	Script string `bun:"script"`
	AllowMount int8 `bun:"allowMount"`
}

type DBInstanceTemplateSlice []DBInstanceTemplate

func (entry *InstanceTemplate) ToDB() *DBInstanceTemplate {
	if entry == nil {
		return nil
	}
	return &DBInstanceTemplate{
		Map: entry.Map,
		Parent: entry.Parent,
		Script: entry.Script,
		AllowMount: entry.AllowMount,
	}
}

func (entry *DBInstanceTemplate) ToWeb() *InstanceTemplate {
	if entry == nil {
		return nil
	}
	return &InstanceTemplate{
		Map: entry.Map,
		Parent: entry.Parent,
		Script: entry.Script,
		AllowMount: entry.AllowMount,
	}
}

func (data InstanceTemplateSlice) ToDB() DBInstanceTemplateSlice {
	result := make(DBInstanceTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBInstanceTemplateSlice) ToWeb() InstanceTemplateSlice {
	result := make(InstanceTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
