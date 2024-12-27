package model

import "github.com/uptrace/bun"

type Bugreport struct {
	Id int `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

type BugreportSlice []Bugreport

type DBBugreport struct {
	bun.BaseModel `bun:"table:bugreport,alias:bugreport"`
	Id int `bun:"id"`
	Type string `bun:"type"`
	Content string `bun:"content"`
}

type DBBugreportSlice []DBBugreport

func (entry *Bugreport) ToDB() *DBBugreport {
	if entry == nil {
		return nil
	}
	return &DBBugreport{
		Id: entry.Id,
		Type: entry.Type,
		Content: entry.Content,
	}
}

func (entry *DBBugreport) ToWeb() *Bugreport {
	if entry == nil {
		return nil
	}
	return &Bugreport{
		Id: entry.Id,
		Type: entry.Type,
		Content: entry.Content,
	}
}

func (data BugreportSlice) ToDB() DBBugreportSlice {
	result := make(DBBugreportSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBugreportSlice) ToWeb() BugreportSlice {
	result := make(BugreportSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
