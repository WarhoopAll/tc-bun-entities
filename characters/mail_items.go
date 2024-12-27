package model

import "github.com/uptrace/bun"

type MailItems struct {
	MailId int `json:"mail_id,omitempty"`
	ItemGuid int `json:"item_guid,omitempty"`
	Receiver int `json:"receiver,omitempty"`
}

type MailItemsSlice []MailItems

type DBMailItems struct {
	bun.BaseModel `bun:"table:mail_items,alias:mail_items"`
	MailId int `bun:"mail_id"`
	ItemGuid int `bun:"item_guid"`
	Receiver int `bun:"receiver"`
}

type DBMailItemsSlice []DBMailItems

func (entry *MailItems) ToDB() *DBMailItems {
	if entry == nil {
		return nil
	}
	return &DBMailItems{
		MailId: entry.MailId,
		ItemGuid: entry.ItemGuid,
		Receiver: entry.Receiver,
	}
}

func (entry *DBMailItems) ToWeb() *MailItems {
	if entry == nil {
		return nil
	}
	return &MailItems{
		MailId: entry.MailId,
		ItemGuid: entry.ItemGuid,
		Receiver: entry.Receiver,
	}
}

func (data MailItemsSlice) ToDB() DBMailItemsSlice {
	result := make(DBMailItemsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBMailItemsSlice) ToWeb() MailItemsSlice {
	result := make(MailItemsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
