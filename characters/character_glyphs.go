package model

import "github.com/uptrace/bun"

type CharacterGlyphs struct {
	Guid int `json:"guid,omitempty"`
	TalentGroup int8 `json:"talentgroup,omitempty"`
	Glyph1 int16 `json:"glyph1,omitempty"`
	Glyph2 int16 `json:"glyph2,omitempty"`
	Glyph3 int16 `json:"glyph3,omitempty"`
	Glyph4 int16 `json:"glyph4,omitempty"`
	Glyph5 int16 `json:"glyph5,omitempty"`
	Glyph6 int16 `json:"glyph6,omitempty"`
}

type CharacterGlyphsSlice []CharacterGlyphs

type DBCharacterGlyphs struct {
	bun.BaseModel `bun:"table:character_glyphs,alias:character_glyphs"`
	Guid int `bun:"guid"`
	TalentGroup int8 `bun:"talentGroup"`
	Glyph1 int16 `bun:"glyph1"`
	Glyph2 int16 `bun:"glyph2"`
	Glyph3 int16 `bun:"glyph3"`
	Glyph4 int16 `bun:"glyph4"`
	Glyph5 int16 `bun:"glyph5"`
	Glyph6 int16 `bun:"glyph6"`
}

type DBCharacterGlyphsSlice []DBCharacterGlyphs

func (entry *CharacterGlyphs) ToDB() *DBCharacterGlyphs {
	if entry == nil {
		return nil
	}
	return &DBCharacterGlyphs{
		Guid: entry.Guid,
		TalentGroup: entry.TalentGroup,
		Glyph1: entry.Glyph1,
		Glyph2: entry.Glyph2,
		Glyph3: entry.Glyph3,
		Glyph4: entry.Glyph4,
		Glyph5: entry.Glyph5,
		Glyph6: entry.Glyph6,
	}
}

func (entry *DBCharacterGlyphs) ToWeb() *CharacterGlyphs {
	if entry == nil {
		return nil
	}
	return &CharacterGlyphs{
		Guid: entry.Guid,
		TalentGroup: entry.TalentGroup,
		Glyph1: entry.Glyph1,
		Glyph2: entry.Glyph2,
		Glyph3: entry.Glyph3,
		Glyph4: entry.Glyph4,
		Glyph5: entry.Glyph5,
		Glyph6: entry.Glyph6,
	}
}

func (data CharacterGlyphsSlice) ToDB() DBCharacterGlyphsSlice {
	result := make(DBCharacterGlyphsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterGlyphsSlice) ToWeb() CharacterGlyphsSlice {
	result := make(CharacterGlyphsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
