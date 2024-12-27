package model

import "github.com/uptrace/bun"

type GameobjectTemplate struct {
	Entry int `json:"entry,omitempty"`
	Type int8 `json:"type,omitempty"`
	DisplayId int `json:"displayid,omitempty"`
	Name string `json:"name,omitempty"`
	IconName string `json:"iconname,omitempty"`
	CastBarCaption string `json:"castbarcaption,omitempty"`
	Unk1 string `json:"unk1,omitempty"`
	Size float64 `json:"size,omitempty"`
	Data0 int `json:"data0,omitempty"`
	Data1 int `json:"data1,omitempty"`
	Data2 int `json:"data2,omitempty"`
	Data3 int `json:"data3,omitempty"`
	Data4 int `json:"data4,omitempty"`
	Data5 int `json:"data5,omitempty"`
	Data6 int `json:"data6,omitempty"`
	Data7 int `json:"data7,omitempty"`
	Data8 int `json:"data8,omitempty"`
	Data9 int `json:"data9,omitempty"`
	Data10 int `json:"data10,omitempty"`
	Data11 int `json:"data11,omitempty"`
	Data12 int `json:"data12,omitempty"`
	Data13 int `json:"data13,omitempty"`
	Data14 int `json:"data14,omitempty"`
	Data15 int `json:"data15,omitempty"`
	Data16 int `json:"data16,omitempty"`
	Data17 int `json:"data17,omitempty"`
	Data18 int `json:"data18,omitempty"`
	Data19 int `json:"data19,omitempty"`
	Data20 int `json:"data20,omitempty"`
	Data21 int `json:"data21,omitempty"`
	Data22 int `json:"data22,omitempty"`
	Data23 int `json:"data23,omitempty"`
	AIName string `json:"ainame,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	StringId string `json:"stringid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type GameobjectTemplateSlice []GameobjectTemplate

type DBGameobjectTemplate struct {
	bun.BaseModel `bun:"table:gameobject_template,alias:gameobject_template"`
	Entry int `bun:"entry"`
	Type int8 `bun:"type"`
	DisplayId int `bun:"displayId"`
	Name string `bun:"name"`
	IconName string `bun:"IconName"`
	CastBarCaption string `bun:"castBarCaption"`
	Unk1 string `bun:"unk1"`
	Size float64 `bun:"size"`
	Data0 int `bun:"Data0"`
	Data1 int `bun:"Data1"`
	Data2 int `bun:"Data2"`
	Data3 int `bun:"Data3"`
	Data4 int `bun:"Data4"`
	Data5 int `bun:"Data5"`
	Data6 int `bun:"Data6"`
	Data7 int `bun:"Data7"`
	Data8 int `bun:"Data8"`
	Data9 int `bun:"Data9"`
	Data10 int `bun:"Data10"`
	Data11 int `bun:"Data11"`
	Data12 int `bun:"Data12"`
	Data13 int `bun:"Data13"`
	Data14 int `bun:"Data14"`
	Data15 int `bun:"Data15"`
	Data16 int `bun:"Data16"`
	Data17 int `bun:"Data17"`
	Data18 int `bun:"Data18"`
	Data19 int `bun:"Data19"`
	Data20 int `bun:"Data20"`
	Data21 int `bun:"Data21"`
	Data22 int `bun:"Data22"`
	Data23 int `bun:"Data23"`
	AIName string `bun:"AIName"`
	ScriptName string `bun:"ScriptName"`
	StringId string `bun:"StringId"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBGameobjectTemplateSlice []DBGameobjectTemplate

func (entry *GameobjectTemplate) ToDB() *DBGameobjectTemplate {
	if entry == nil {
		return nil
	}
	return &DBGameobjectTemplate{
		Entry: entry.Entry,
		Type: entry.Type,
		DisplayId: entry.DisplayId,
		Name: entry.Name,
		IconName: entry.IconName,
		CastBarCaption: entry.CastBarCaption,
		Unk1: entry.Unk1,
		Size: entry.Size,
		Data0: entry.Data0,
		Data1: entry.Data1,
		Data2: entry.Data2,
		Data3: entry.Data3,
		Data4: entry.Data4,
		Data5: entry.Data5,
		Data6: entry.Data6,
		Data7: entry.Data7,
		Data8: entry.Data8,
		Data9: entry.Data9,
		Data10: entry.Data10,
		Data11: entry.Data11,
		Data12: entry.Data12,
		Data13: entry.Data13,
		Data14: entry.Data14,
		Data15: entry.Data15,
		Data16: entry.Data16,
		Data17: entry.Data17,
		Data18: entry.Data18,
		Data19: entry.Data19,
		Data20: entry.Data20,
		Data21: entry.Data21,
		Data22: entry.Data22,
		Data23: entry.Data23,
		AIName: entry.AIName,
		ScriptName: entry.ScriptName,
		StringId: entry.StringId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBGameobjectTemplate) ToWeb() *GameobjectTemplate {
	if entry == nil {
		return nil
	}
	return &GameobjectTemplate{
		Entry: entry.Entry,
		Type: entry.Type,
		DisplayId: entry.DisplayId,
		Name: entry.Name,
		IconName: entry.IconName,
		CastBarCaption: entry.CastBarCaption,
		Unk1: entry.Unk1,
		Size: entry.Size,
		Data0: entry.Data0,
		Data1: entry.Data1,
		Data2: entry.Data2,
		Data3: entry.Data3,
		Data4: entry.Data4,
		Data5: entry.Data5,
		Data6: entry.Data6,
		Data7: entry.Data7,
		Data8: entry.Data8,
		Data9: entry.Data9,
		Data10: entry.Data10,
		Data11: entry.Data11,
		Data12: entry.Data12,
		Data13: entry.Data13,
		Data14: entry.Data14,
		Data15: entry.Data15,
		Data16: entry.Data16,
		Data17: entry.Data17,
		Data18: entry.Data18,
		Data19: entry.Data19,
		Data20: entry.Data20,
		Data21: entry.Data21,
		Data22: entry.Data22,
		Data23: entry.Data23,
		AIName: entry.AIName,
		ScriptName: entry.ScriptName,
		StringId: entry.StringId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data GameobjectTemplateSlice) ToDB() DBGameobjectTemplateSlice {
	result := make(DBGameobjectTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectTemplateSlice) ToWeb() GameobjectTemplateSlice {
	result := make(GameobjectTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
