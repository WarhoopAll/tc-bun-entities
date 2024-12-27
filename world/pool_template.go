package model

import "github.com/uptrace/bun"

type PoolTemplate struct {
	Entry int `json:"entry,omitempty"`
	MaxLimit int `json:"max_limit,omitempty"`
	Description string `json:"description,omitempty"`
}

type PoolTemplateSlice []PoolTemplate

type DBPoolTemplate struct {
	bun.BaseModel `bun:"table:pool_template,alias:pool_template"`
	Entry int `bun:"entry"`
	MaxLimit int `bun:"max_limit"`
	Description string `bun:"description"`
}

type DBPoolTemplateSlice []DBPoolTemplate

func (entry *PoolTemplate) ToDB() *DBPoolTemplate {
	if entry == nil {
		return nil
	}
	return &DBPoolTemplate{
		Entry: entry.Entry,
		MaxLimit: entry.MaxLimit,
		Description: entry.Description,
	}
}

func (entry *DBPoolTemplate) ToWeb() *PoolTemplate {
	if entry == nil {
		return nil
	}
	return &PoolTemplate{
		Entry: entry.Entry,
		MaxLimit: entry.MaxLimit,
		Description: entry.Description,
	}
}

func (data PoolTemplateSlice) ToDB() DBPoolTemplateSlice {
	result := make(DBPoolTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPoolTemplateSlice) ToWeb() PoolTemplateSlice {
	result := make(PoolTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
