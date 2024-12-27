package model

import "github.com/uptrace/bun"

type Auctionhouse struct {
	Id int `json:"id,omitempty"`
	Houseid int8 `json:"houseid,omitempty"`
	Itemguid int `json:"itemguid,omitempty"`
	Itemowner int `json:"itemowner,omitempty"`
	Buyoutprice int `json:"buyoutprice,omitempty"`
	Time int `json:"time,omitempty"`
	Buyguid int `json:"buyguid,omitempty"`
	Lastbid int `json:"lastbid,omitempty"`
	Startbid int `json:"startbid,omitempty"`
	Deposit int `json:"deposit,omitempty"`
	Flags int8 `json:"flags,omitempty"`
}

type AuctionhouseSlice []Auctionhouse

type DBAuctionhouse struct {
	bun.BaseModel `bun:"table:auctionhouse,alias:auctionhouse"`
	Id int `bun:"id"`
	Houseid int8 `bun:"houseid"`
	Itemguid int `bun:"itemguid"`
	Itemowner int `bun:"itemowner"`
	Buyoutprice int `bun:"buyoutprice"`
	Time int `bun:"time"`
	Buyguid int `bun:"buyguid"`
	Lastbid int `bun:"lastbid"`
	Startbid int `bun:"startbid"`
	Deposit int `bun:"deposit"`
	Flags int8 `bun:"Flags"`
}

type DBAuctionhouseSlice []DBAuctionhouse

func (entry *Auctionhouse) ToDB() *DBAuctionhouse {
	if entry == nil {
		return nil
	}
	return &DBAuctionhouse{
		Id: entry.Id,
		Houseid: entry.Houseid,
		Itemguid: entry.Itemguid,
		Itemowner: entry.Itemowner,
		Buyoutprice: entry.Buyoutprice,
		Time: entry.Time,
		Buyguid: entry.Buyguid,
		Lastbid: entry.Lastbid,
		Startbid: entry.Startbid,
		Deposit: entry.Deposit,
		Flags: entry.Flags,
	}
}

func (entry *DBAuctionhouse) ToWeb() *Auctionhouse {
	if entry == nil {
		return nil
	}
	return &Auctionhouse{
		Id: entry.Id,
		Houseid: entry.Houseid,
		Itemguid: entry.Itemguid,
		Itemowner: entry.Itemowner,
		Buyoutprice: entry.Buyoutprice,
		Time: entry.Time,
		Buyguid: entry.Buyguid,
		Lastbid: entry.Lastbid,
		Startbid: entry.Startbid,
		Deposit: entry.Deposit,
		Flags: entry.Flags,
	}
}

func (data AuctionhouseSlice) ToDB() DBAuctionhouseSlice {
	result := make(DBAuctionhouseSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAuctionhouseSlice) ToWeb() AuctionhouseSlice {
	result := make(AuctionhouseSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
