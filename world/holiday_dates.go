package model

import "github.com/uptrace/bun"

type HolidayDates struct {
	Id int `json:"id,omitempty"`
	DateId int8 `json:"date_id,omitempty"`
	DateValue int `json:"date_value,omitempty"`
	HolidayDuration int `json:"holiday_duration,omitempty"`
}

type HolidayDatesSlice []HolidayDates

type DBHolidayDates struct {
	bun.BaseModel `bun:"table:holiday_dates,alias:holiday_dates"`
	Id int `bun:"id"`
	DateId int8 `bun:"date_id"`
	DateValue int `bun:"date_value"`
	HolidayDuration int `bun:"holiday_duration"`
}

type DBHolidayDatesSlice []DBHolidayDates

func (entry *HolidayDates) ToDB() *DBHolidayDates {
	if entry == nil {
		return nil
	}
	return &DBHolidayDates{
		Id: entry.Id,
		DateId: entry.DateId,
		DateValue: entry.DateValue,
		HolidayDuration: entry.HolidayDuration,
	}
}

func (entry *DBHolidayDates) ToWeb() *HolidayDates {
	if entry == nil {
		return nil
	}
	return &HolidayDates{
		Id: entry.Id,
		DateId: entry.DateId,
		DateValue: entry.DateValue,
		HolidayDuration: entry.HolidayDuration,
	}
}

func (data HolidayDatesSlice) ToDB() DBHolidayDatesSlice {
	result := make(DBHolidayDatesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBHolidayDatesSlice) ToWeb() HolidayDatesSlice {
	result := make(HolidayDatesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
