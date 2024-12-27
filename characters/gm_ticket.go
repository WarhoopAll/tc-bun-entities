package model

import "github.com/uptrace/bun"

type GmTicket struct {
	Id int `json:"id,omitempty"`
	Type int8 `json:"type,omitempty"`
	PlayerGuid int `json:"playerguid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreateTime int `json:"createtime,omitempty"`
	MapId int16 `json:"mapid,omitempty"`
	PosX float64 `json:"posx,omitempty"`
	PosY float64 `json:"posy,omitempty"`
	PosZ float64 `json:"posz,omitempty"`
	LastModifiedTime int `json:"lastmodifiedtime,omitempty"`
	ClosedBy int `json:"closedby,omitempty"`
	AssignedTo int `json:"assignedto,omitempty"`
	Comment string `json:"comment,omitempty"`
	Response string `json:"response,omitempty"`
	Completed int8 `json:"completed,omitempty"`
	Escalated int8 `json:"escalated,omitempty"`
	Viewed int8 `json:"viewed,omitempty"`
	NeedMoreHelp int8 `json:"needmorehelp,omitempty"`
	ResolvedBy int `json:"resolvedby,omitempty"`
}

type GmTicketSlice []GmTicket

type DBGmTicket struct {
	bun.BaseModel `bun:"table:gm_ticket,alias:gm_ticket"`
	Id int `bun:"id"`
	Type int8 `bun:"type"`
	PlayerGuid int `bun:"playerGuid"`
	Name string `bun:"name"`
	Description string `bun:"description"`
	CreateTime int `bun:"createTime"`
	MapId int16 `bun:"mapId"`
	PosX float64 `bun:"posX"`
	PosY float64 `bun:"posY"`
	PosZ float64 `bun:"posZ"`
	LastModifiedTime int `bun:"lastModifiedTime"`
	ClosedBy int `bun:"closedBy"`
	AssignedTo int `bun:"assignedTo"`
	Comment string `bun:"comment"`
	Response string `bun:"response"`
	Completed int8 `bun:"completed"`
	Escalated int8 `bun:"escalated"`
	Viewed int8 `bun:"viewed"`
	NeedMoreHelp int8 `bun:"needMoreHelp"`
	ResolvedBy int `bun:"resolvedBy"`
}

type DBGmTicketSlice []DBGmTicket

func (entry *GmTicket) ToDB() *DBGmTicket {
	if entry == nil {
		return nil
	}
	return &DBGmTicket{
		Id: entry.Id,
		Type: entry.Type,
		PlayerGuid: entry.PlayerGuid,
		Name: entry.Name,
		Description: entry.Description,
		CreateTime: entry.CreateTime,
		MapId: entry.MapId,
		PosX: entry.PosX,
		PosY: entry.PosY,
		PosZ: entry.PosZ,
		LastModifiedTime: entry.LastModifiedTime,
		ClosedBy: entry.ClosedBy,
		AssignedTo: entry.AssignedTo,
		Comment: entry.Comment,
		Response: entry.Response,
		Completed: entry.Completed,
		Escalated: entry.Escalated,
		Viewed: entry.Viewed,
		NeedMoreHelp: entry.NeedMoreHelp,
		ResolvedBy: entry.ResolvedBy,
	}
}

func (entry *DBGmTicket) ToWeb() *GmTicket {
	if entry == nil {
		return nil
	}
	return &GmTicket{
		Id: entry.Id,
		Type: entry.Type,
		PlayerGuid: entry.PlayerGuid,
		Name: entry.Name,
		Description: entry.Description,
		CreateTime: entry.CreateTime,
		MapId: entry.MapId,
		PosX: entry.PosX,
		PosY: entry.PosY,
		PosZ: entry.PosZ,
		LastModifiedTime: entry.LastModifiedTime,
		ClosedBy: entry.ClosedBy,
		AssignedTo: entry.AssignedTo,
		Comment: entry.Comment,
		Response: entry.Response,
		Completed: entry.Completed,
		Escalated: entry.Escalated,
		Viewed: entry.Viewed,
		NeedMoreHelp: entry.NeedMoreHelp,
		ResolvedBy: entry.ResolvedBy,
	}
}

func (data GmTicketSlice) ToDB() DBGmTicketSlice {
	result := make(DBGmTicketSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGmTicketSlice) ToWeb() GmTicketSlice {
	result := make(GmTicketSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
