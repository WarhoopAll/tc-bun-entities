package model

import "github.com/uptrace/bun"

type GameWeather struct {
	Zone int `json:"zone,omitempty"`
	SpringRainChance int8 `json:"spring_rain_chance,omitempty"`
	SpringSnowChance int8 `json:"spring_snow_chance,omitempty"`
	SpringStormChance int8 `json:"spring_storm_chance,omitempty"`
	SummerRainChance int8 `json:"summer_rain_chance,omitempty"`
	SummerSnowChance int8 `json:"summer_snow_chance,omitempty"`
	SummerStormChance int8 `json:"summer_storm_chance,omitempty"`
	FallRainChance int8 `json:"fall_rain_chance,omitempty"`
	FallSnowChance int8 `json:"fall_snow_chance,omitempty"`
	FallStormChance int8 `json:"fall_storm_chance,omitempty"`
	WinterRainChance int8 `json:"winter_rain_chance,omitempty"`
	WinterSnowChance int8 `json:"winter_snow_chance,omitempty"`
	WinterStormChance int8 `json:"winter_storm_chance,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
}

type GameWeatherSlice []GameWeather

type DBGameWeather struct {
	bun.BaseModel `bun:"table:game_weather,alias:game_weather"`
	Zone int `bun:"zone"`
	SpringRainChance int8 `bun:"spring_rain_chance"`
	SpringSnowChance int8 `bun:"spring_snow_chance"`
	SpringStormChance int8 `bun:"spring_storm_chance"`
	SummerRainChance int8 `bun:"summer_rain_chance"`
	SummerSnowChance int8 `bun:"summer_snow_chance"`
	SummerStormChance int8 `bun:"summer_storm_chance"`
	FallRainChance int8 `bun:"fall_rain_chance"`
	FallSnowChance int8 `bun:"fall_snow_chance"`
	FallStormChance int8 `bun:"fall_storm_chance"`
	WinterRainChance int8 `bun:"winter_rain_chance"`
	WinterSnowChance int8 `bun:"winter_snow_chance"`
	WinterStormChance int8 `bun:"winter_storm_chance"`
	ScriptName string `bun:"ScriptName"`
}

type DBGameWeatherSlice []DBGameWeather

func (entry *GameWeather) ToDB() *DBGameWeather {
	if entry == nil {
		return nil
	}
	return &DBGameWeather{
		Zone: entry.Zone,
		SpringRainChance: entry.SpringRainChance,
		SpringSnowChance: entry.SpringSnowChance,
		SpringStormChance: entry.SpringStormChance,
		SummerRainChance: entry.SummerRainChance,
		SummerSnowChance: entry.SummerSnowChance,
		SummerStormChance: entry.SummerStormChance,
		FallRainChance: entry.FallRainChance,
		FallSnowChance: entry.FallSnowChance,
		FallStormChance: entry.FallStormChance,
		WinterRainChance: entry.WinterRainChance,
		WinterSnowChance: entry.WinterSnowChance,
		WinterStormChance: entry.WinterStormChance,
		ScriptName: entry.ScriptName,
	}
}

func (entry *DBGameWeather) ToWeb() *GameWeather {
	if entry == nil {
		return nil
	}
	return &GameWeather{
		Zone: entry.Zone,
		SpringRainChance: entry.SpringRainChance,
		SpringSnowChance: entry.SpringSnowChance,
		SpringStormChance: entry.SpringStormChance,
		SummerRainChance: entry.SummerRainChance,
		SummerSnowChance: entry.SummerSnowChance,
		SummerStormChance: entry.SummerStormChance,
		FallRainChance: entry.FallRainChance,
		FallSnowChance: entry.FallSnowChance,
		FallStormChance: entry.FallStormChance,
		WinterRainChance: entry.WinterRainChance,
		WinterSnowChance: entry.WinterSnowChance,
		WinterStormChance: entry.WinterStormChance,
		ScriptName: entry.ScriptName,
	}
}

func (data GameWeatherSlice) ToDB() DBGameWeatherSlice {
	result := make(DBGameWeatherSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameWeatherSlice) ToWeb() GameWeatherSlice {
	result := make(GameWeatherSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
