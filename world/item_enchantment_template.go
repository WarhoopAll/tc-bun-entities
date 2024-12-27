package model

import "github.com/uptrace/bun"

type ItemEnchantmentTemplate struct {
	Entry int `json:"entry,omitempty"`
	Ench int `json:"ench,omitempty"`
	Chance float64 `json:"chance,omitempty"`
}

type ItemEnchantmentTemplateSlice []ItemEnchantmentTemplate

type DBItemEnchantmentTemplate struct {
	bun.BaseModel `bun:"table:item_enchantment_template,alias:item_enchantment_template"`
	Entry int `bun:"entry"`
	Ench int `bun:"ench"`
	Chance float64 `bun:"chance"`
}

type DBItemEnchantmentTemplateSlice []DBItemEnchantmentTemplate

func (entry *ItemEnchantmentTemplate) ToDB() *DBItemEnchantmentTemplate {
	if entry == nil {
		return nil
	}
	return &DBItemEnchantmentTemplate{
		Entry: entry.Entry,
		Ench: entry.Ench,
		Chance: entry.Chance,
	}
}

func (entry *DBItemEnchantmentTemplate) ToWeb() *ItemEnchantmentTemplate {
	if entry == nil {
		return nil
	}
	return &ItemEnchantmentTemplate{
		Entry: entry.Entry,
		Ench: entry.Ench,
		Chance: entry.Chance,
	}
}

func (data ItemEnchantmentTemplateSlice) ToDB() DBItemEnchantmentTemplateSlice {
	result := make(DBItemEnchantmentTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemEnchantmentTemplateSlice) ToWeb() ItemEnchantmentTemplateSlice {
	result := make(ItemEnchantmentTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
