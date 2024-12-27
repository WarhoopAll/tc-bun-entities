package model

import "github.com/uptrace/bun"

type WardenChecks struct {
	Id int16 `json:"id,omitempty"`
	Type int8 `json:"type,omitempty"`
	Str string `json:"str,omitempty"`
	Address int `json:"address,omitempty"`
	Length int8 `json:"length,omitempty"`
	Comment string `json:"comment,omitempty"`
	Data []byte `json:"data,omitempty"`
	Result []byte `json:"result,omitempty"`
}

type WardenChecksSlice []WardenChecks

type DBWardenChecks struct {
	bun.BaseModel `bun:"table:warden_checks,alias:warden_checks"`
	Id int16 `bun:"id"`
	Type int8 `bun:"type"`
	Str string `bun:"str"`
	Address int `bun:"address"`
	Length int8 `bun:"length"`
	Comment string `bun:"comment"`
	Data []byte `bun:"data"`
	Result []byte `bun:"result"`
}

type DBWardenChecksSlice []DBWardenChecks

func (entry *WardenChecks) ToDB() *DBWardenChecks {
	if entry == nil {
		return nil
	}
	return &DBWardenChecks{
		Id: entry.Id,
		Type: entry.Type,
		Str: entry.Str,
		Address: entry.Address,
		Length: entry.Length,
		Comment: entry.Comment,
		Data: entry.Data,
		Result: entry.Result,
	}
}

func (entry *DBWardenChecks) ToWeb() *WardenChecks {
	if entry == nil {
		return nil
	}
	return &WardenChecks{
		Id: entry.Id,
		Type: entry.Type,
		Str: entry.Str,
		Address: entry.Address,
		Length: entry.Length,
		Comment: entry.Comment,
		Data: entry.Data,
		Result: entry.Result,
	}
}

func (data WardenChecksSlice) ToDB() DBWardenChecksSlice {
	result := make(DBWardenChecksSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBWardenChecksSlice) ToWeb() WardenChecksSlice {
	result := make(WardenChecksSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
