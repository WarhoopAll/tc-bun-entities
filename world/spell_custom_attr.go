package model

import "github.com/uptrace/bun"

type SpellCustomAttr struct {
	Entry int `json:"entry,omitempty"`
	Attributes int `json:"attributes,omitempty"`
}

type SpellCustomAttrSlice []SpellCustomAttr

type DBSpellCustomAttr struct {
	bun.BaseModel `bun:"table:spell_custom_attr,alias:spell_custom_attr"`
	Entry int `bun:"entry"`
	Attributes int `bun:"attributes"`
}

type DBSpellCustomAttrSlice []DBSpellCustomAttr

func (entry *SpellCustomAttr) ToDB() *DBSpellCustomAttr {
	if entry == nil {
		return nil
	}
	return &DBSpellCustomAttr{
		Entry: entry.Entry,
		Attributes: entry.Attributes,
	}
}

func (entry *DBSpellCustomAttr) ToWeb() *SpellCustomAttr {
	if entry == nil {
		return nil
	}
	return &SpellCustomAttr{
		Entry: entry.Entry,
		Attributes: entry.Attributes,
	}
}

func (data SpellCustomAttrSlice) ToDB() DBSpellCustomAttrSlice {
	result := make(DBSpellCustomAttrSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellCustomAttrSlice) ToWeb() SpellCustomAttrSlice {
	result := make(SpellCustomAttrSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
