package model

import "github.com/uptrace/bun"

type Mail struct {
	Id int `json:"id,omitempty"`
	MessageType int8 `json:"messagetype,omitempty"`
	Stationery int8 `json:"stationery,omitempty"`
	MailTemplateId int16 `json:"mailtemplateid,omitempty"`
	Sender int `json:"sender,omitempty"`
	Receiver int `json:"receiver,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body string `json:"body,omitempty"`
	HasItems int8 `json:"has_items,omitempty"`
	ExpireTime int `json:"expire_time,omitempty"`
	DeliverTime int `json:"deliver_time,omitempty"`
	Money int `json:"money,omitempty"`
	Cod int `json:"cod,omitempty"`
	Checked int8 `json:"checked,omitempty"`
}

type MailSlice []Mail

type DBMail struct {
	bun.BaseModel `bun:"table:mail,alias:mail"`
	Id int `bun:"id"`
	MessageType int8 `bun:"messageType"`
	Stationery int8 `bun:"stationery"`
	MailTemplateId int16 `bun:"mailTemplateId"`
	Sender int `bun:"sender"`
	Receiver int `bun:"receiver"`
	Subject string `bun:"subject"`
	Body string `bun:"body"`
	HasItems int8 `bun:"has_items"`
	ExpireTime int `bun:"expire_time"`
	DeliverTime int `bun:"deliver_time"`
	Money int `bun:"money"`
	Cod int `bun:"cod"`
	Checked int8 `bun:"checked"`
}

type DBMailSlice []DBMail

func (entry *Mail) ToDB() *DBMail {
	if entry == nil {
		return nil
	}
	return &DBMail{
		Id: entry.Id,
		MessageType: entry.MessageType,
		Stationery: entry.Stationery,
		MailTemplateId: entry.MailTemplateId,
		Sender: entry.Sender,
		Receiver: entry.Receiver,
		Subject: entry.Subject,
		Body: entry.Body,
		HasItems: entry.HasItems,
		ExpireTime: entry.ExpireTime,
		DeliverTime: entry.DeliverTime,
		Money: entry.Money,
		Cod: entry.Cod,
		Checked: entry.Checked,
	}
}

func (entry *DBMail) ToWeb() *Mail {
	if entry == nil {
		return nil
	}
	return &Mail{
		Id: entry.Id,
		MessageType: entry.MessageType,
		Stationery: entry.Stationery,
		MailTemplateId: entry.MailTemplateId,
		Sender: entry.Sender,
		Receiver: entry.Receiver,
		Subject: entry.Subject,
		Body: entry.Body,
		HasItems: entry.HasItems,
		ExpireTime: entry.ExpireTime,
		DeliverTime: entry.DeliverTime,
		Money: entry.Money,
		Cod: entry.Cod,
		Checked: entry.Checked,
	}
}

func (data MailSlice) ToDB() DBMailSlice {
	result := make(DBMailSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBMailSlice) ToWeb() MailSlice {
	result := make(MailSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
