package model

import "github.com/uptrace/bun"

type PointsOfInterestLocale struct {
	ID int `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	Name string `json:"name,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type PointsOfInterestLocaleSlice []PointsOfInterestLocale

type DBPointsOfInterestLocale struct {
	bun.BaseModel `bun:"table:points_of_interest_locale,alias:points_of_interest_locale"`
	ID int `bun:"ID"`
	Locale string `bun:"locale"`
	Name string `bun:"Name"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBPointsOfInterestLocaleSlice []DBPointsOfInterestLocale

func (entry *PointsOfInterestLocale) ToDB() *DBPointsOfInterestLocale {
	if entry == nil {
		return nil
	}
	return &DBPointsOfInterestLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Name: entry.Name,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBPointsOfInterestLocale) ToWeb() *PointsOfInterestLocale {
	if entry == nil {
		return nil
	}
	return &PointsOfInterestLocale{
		ID: entry.ID,
		Locale: entry.Locale,
		Name: entry.Name,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data PointsOfInterestLocaleSlice) ToDB() DBPointsOfInterestLocaleSlice {
	result := make(DBPointsOfInterestLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPointsOfInterestLocaleSlice) ToWeb() PointsOfInterestLocaleSlice {
	result := make(PointsOfInterestLocaleSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
