package model

import "github.com/uptrace/bun"

type Auctionbidders struct {
	Id int `json:"id,omitempty"`
	Bidderguid int `json:"bidderguid,omitempty"`
}

type AuctionbiddersSlice []Auctionbidders

type DBAuctionbidders struct {
	bun.BaseModel `bun:"table:auctionbidders,alias:auctionbidders"`
	Id int `bun:"id"`
	Bidderguid int `bun:"bidderguid"`
}

type DBAuctionbiddersSlice []DBAuctionbidders

func (entry *Auctionbidders) ToDB() *DBAuctionbidders {
	if entry == nil {
		return nil
	}
	return &DBAuctionbidders{
		Id: entry.Id,
		Bidderguid: entry.Bidderguid,
	}
}

func (entry *DBAuctionbidders) ToWeb() *Auctionbidders {
	if entry == nil {
		return nil
	}
	return &Auctionbidders{
		Id: entry.Id,
		Bidderguid: entry.Bidderguid,
	}
}

func (data AuctionbiddersSlice) ToDB() DBAuctionbiddersSlice {
	result := make(DBAuctionbiddersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAuctionbiddersSlice) ToWeb() AuctionbiddersSlice {
	result := make(AuctionbiddersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
