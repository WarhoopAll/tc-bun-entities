package model

import "github.com/uptrace/bun"

type PointsOfInterest struct {
	ID int `json:"id,omitempty"`
	PositionX float64 `json:"positionx,omitempty"`
	PositionY float64 `json:"positiony,omitempty"`
	Icon int `json:"icon,omitempty"`
	Flags int `json:"flags,omitempty"`
	Importance int `json:"importance,omitempty"`
	Name string `json:"name,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type PointsOfInterestSlice []PointsOfInterest

type DBPointsOfInterest struct {
	bun.BaseModel `bun:"table:points_of_interest,alias:points_of_interest"`
	ID int `bun:"ID"`
	PositionX float64 `bun:"PositionX"`
	PositionY float64 `bun:"PositionY"`
	Icon int `bun:"Icon"`
	Flags int `bun:"Flags"`
	Importance int `bun:"Importance"`
	Name string `bun:"Name"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBPointsOfInterestSlice []DBPointsOfInterest

func (entry *PointsOfInterest) ToDB() *DBPointsOfInterest {
	if entry == nil {
		return nil
	}
	return &DBPointsOfInterest{
		ID: entry.ID,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		Icon: entry.Icon,
		Flags: entry.Flags,
		Importance: entry.Importance,
		Name: entry.Name,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBPointsOfInterest) ToWeb() *PointsOfInterest {
	if entry == nil {
		return nil
	}
	return &PointsOfInterest{
		ID: entry.ID,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		Icon: entry.Icon,
		Flags: entry.Flags,
		Importance: entry.Importance,
		Name: entry.Name,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data PointsOfInterestSlice) ToDB() DBPointsOfInterestSlice {
	result := make(DBPointsOfInterestSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPointsOfInterestSlice) ToWeb() PointsOfInterestSlice {
	result := make(PointsOfInterestSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
