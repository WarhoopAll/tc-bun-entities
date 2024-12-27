package model

import "github.com/uptrace/bun"

type VwDisablesWithLabels struct {
	SourceType string `json:"sourcetype,omitempty"`
	Entry int `json:"entry,omitempty"`
	Flags int16 `json:"flags,omitempty"`
	Params0 string `json:"params_0,omitempty"`
	Params1 string `json:"params_1,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type VwDisablesWithLabelsSlice []VwDisablesWithLabels

type DBVwDisablesWithLabels struct {
	bun.BaseModel `bun:"table:vw_disables_with_labels,alias:vw_disables_with_labels"`
	SourceType string `bun:"sourceType"`
	Entry int `bun:"entry"`
	Flags int16 `bun:"flags"`
	Params0 string `bun:"params_0"`
	Params1 string `bun:"params_1"`
	Comment string `bun:"comment"`
}

type DBVwDisablesWithLabelsSlice []DBVwDisablesWithLabels

func (entry *VwDisablesWithLabels) ToDB() *DBVwDisablesWithLabels {
	if entry == nil {
		return nil
	}
	return &DBVwDisablesWithLabels{
		SourceType: entry.SourceType,
		Entry: entry.Entry,
		Flags: entry.Flags,
		Params0: entry.Params0,
		Params1: entry.Params1,
		Comment: entry.Comment,
	}
}

func (entry *DBVwDisablesWithLabels) ToWeb() *VwDisablesWithLabels {
	if entry == nil {
		return nil
	}
	return &VwDisablesWithLabels{
		SourceType: entry.SourceType,
		Entry: entry.Entry,
		Flags: entry.Flags,
		Params0: entry.Params0,
		Params1: entry.Params1,
		Comment: entry.Comment,
	}
}

func (data VwDisablesWithLabelsSlice) ToDB() DBVwDisablesWithLabelsSlice {
	result := make(DBVwDisablesWithLabelsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVwDisablesWithLabelsSlice) ToWeb() VwDisablesWithLabelsSlice {
	result := make(VwDisablesWithLabelsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
