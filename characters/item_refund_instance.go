package model

import "github.com/uptrace/bun"

type ItemRefundInstance struct {
	ItemGuid int `json:"item_guid,omitempty"`
	PlayerGuid int `json:"player_guid,omitempty"`
	PaidMoney int `json:"paidmoney,omitempty"`
	PaidExtendedCost int16 `json:"paidextendedcost,omitempty"`
}

type ItemRefundInstanceSlice []ItemRefundInstance

type DBItemRefundInstance struct {
	bun.BaseModel `bun:"table:item_refund_instance,alias:item_refund_instance"`
	ItemGuid int `bun:"item_guid"`
	PlayerGuid int `bun:"player_guid"`
	PaidMoney int `bun:"paidMoney"`
	PaidExtendedCost int16 `bun:"paidExtendedCost"`
}

type DBItemRefundInstanceSlice []DBItemRefundInstance

func (entry *ItemRefundInstance) ToDB() *DBItemRefundInstance {
	if entry == nil {
		return nil
	}
	return &DBItemRefundInstance{
		ItemGuid: entry.ItemGuid,
		PlayerGuid: entry.PlayerGuid,
		PaidMoney: entry.PaidMoney,
		PaidExtendedCost: entry.PaidExtendedCost,
	}
}

func (entry *DBItemRefundInstance) ToWeb() *ItemRefundInstance {
	if entry == nil {
		return nil
	}
	return &ItemRefundInstance{
		ItemGuid: entry.ItemGuid,
		PlayerGuid: entry.PlayerGuid,
		PaidMoney: entry.PaidMoney,
		PaidExtendedCost: entry.PaidExtendedCost,
	}
}

func (data ItemRefundInstanceSlice) ToDB() DBItemRefundInstanceSlice {
	result := make(DBItemRefundInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemRefundInstanceSlice) ToWeb() ItemRefundInstanceSlice {
	result := make(ItemRefundInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
