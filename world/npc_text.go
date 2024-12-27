package model

import "github.com/uptrace/bun"

type NpcText struct {
	ID int `json:"id,omitempty"`
	Text00 string `json:"text0_0,omitempty"`
	Text01 string `json:"text0_1,omitempty"`
	BroadcastTextID0 int `json:"broadcasttextid0,omitempty"`
	Lang0 int8 `json:"lang0,omitempty"`
	Probability0 float64 `json:"probability0,omitempty"`
	EmoteDelay00 int16 `json:"emotedelay0_0,omitempty"`
	Emote00 int16 `json:"emote0_0,omitempty"`
	EmoteDelay01 int16 `json:"emotedelay0_1,omitempty"`
	Emote01 int16 `json:"emote0_1,omitempty"`
	EmoteDelay02 int16 `json:"emotedelay0_2,omitempty"`
	Emote02 int16 `json:"emote0_2,omitempty"`
	Text10 string `json:"text1_0,omitempty"`
	Text11 string `json:"text1_1,omitempty"`
	BroadcastTextID1 int `json:"broadcasttextid1,omitempty"`
	Lang1 int8 `json:"lang1,omitempty"`
	Probability1 float64 `json:"probability1,omitempty"`
	EmoteDelay10 int16 `json:"emotedelay1_0,omitempty"`
	Emote10 int16 `json:"emote1_0,omitempty"`
	EmoteDelay11 int16 `json:"emotedelay1_1,omitempty"`
	Emote11 int16 `json:"emote1_1,omitempty"`
	EmoteDelay12 int16 `json:"emotedelay1_2,omitempty"`
	Emote12 int16 `json:"emote1_2,omitempty"`
	Text20 string `json:"text2_0,omitempty"`
	Text21 string `json:"text2_1,omitempty"`
	BroadcastTextID2 int `json:"broadcasttextid2,omitempty"`
	Lang2 int8 `json:"lang2,omitempty"`
	Probability2 float64 `json:"probability2,omitempty"`
	EmoteDelay20 int16 `json:"emotedelay2_0,omitempty"`
	Emote20 int16 `json:"emote2_0,omitempty"`
	EmoteDelay21 int16 `json:"emotedelay2_1,omitempty"`
	Emote21 int16 `json:"emote2_1,omitempty"`
	EmoteDelay22 int16 `json:"emotedelay2_2,omitempty"`
	Emote22 int16 `json:"emote2_2,omitempty"`
	Text30 string `json:"text3_0,omitempty"`
	Text31 string `json:"text3_1,omitempty"`
	BroadcastTextID3 int `json:"broadcasttextid3,omitempty"`
	Lang3 int8 `json:"lang3,omitempty"`
	Probability3 float64 `json:"probability3,omitempty"`
	EmoteDelay30 int16 `json:"emotedelay3_0,omitempty"`
	Emote30 int16 `json:"emote3_0,omitempty"`
	EmoteDelay31 int16 `json:"emotedelay3_1,omitempty"`
	Emote31 int16 `json:"emote3_1,omitempty"`
	EmoteDelay32 int16 `json:"emotedelay3_2,omitempty"`
	Emote32 int16 `json:"emote3_2,omitempty"`
	Text40 string `json:"text4_0,omitempty"`
	Text41 string `json:"text4_1,omitempty"`
	BroadcastTextID4 int `json:"broadcasttextid4,omitempty"`
	Lang4 int8 `json:"lang4,omitempty"`
	Probability4 float64 `json:"probability4,omitempty"`
	EmoteDelay40 int16 `json:"emotedelay4_0,omitempty"`
	Emote40 int16 `json:"emote4_0,omitempty"`
	EmoteDelay41 int16 `json:"emotedelay4_1,omitempty"`
	Emote41 int16 `json:"emote4_1,omitempty"`
	EmoteDelay42 int16 `json:"emotedelay4_2,omitempty"`
	Emote42 int16 `json:"emote4_2,omitempty"`
	Text50 string `json:"text5_0,omitempty"`
	Text51 string `json:"text5_1,omitempty"`
	BroadcastTextID5 int `json:"broadcasttextid5,omitempty"`
	Lang5 int8 `json:"lang5,omitempty"`
	Probability5 float64 `json:"probability5,omitempty"`
	EmoteDelay50 int16 `json:"emotedelay5_0,omitempty"`
	Emote50 int16 `json:"emote5_0,omitempty"`
	EmoteDelay51 int16 `json:"emotedelay5_1,omitempty"`
	Emote51 int16 `json:"emote5_1,omitempty"`
	EmoteDelay52 int16 `json:"emotedelay5_2,omitempty"`
	Emote52 int16 `json:"emote5_2,omitempty"`
	Text60 string `json:"text6_0,omitempty"`
	Text61 string `json:"text6_1,omitempty"`
	BroadcastTextID6 int `json:"broadcasttextid6,omitempty"`
	Lang6 int8 `json:"lang6,omitempty"`
	Probability6 float64 `json:"probability6,omitempty"`
	EmoteDelay60 int16 `json:"emotedelay6_0,omitempty"`
	Emote60 int16 `json:"emote6_0,omitempty"`
	EmoteDelay61 int16 `json:"emotedelay6_1,omitempty"`
	Emote61 int16 `json:"emote6_1,omitempty"`
	EmoteDelay62 int16 `json:"emotedelay6_2,omitempty"`
	Emote62 int16 `json:"emote6_2,omitempty"`
	Text70 string `json:"text7_0,omitempty"`
	Text71 string `json:"text7_1,omitempty"`
	BroadcastTextID7 int `json:"broadcasttextid7,omitempty"`
	Lang7 int8 `json:"lang7,omitempty"`
	Probability7 float64 `json:"probability7,omitempty"`
	EmoteDelay70 int16 `json:"emotedelay7_0,omitempty"`
	Emote70 int16 `json:"emote7_0,omitempty"`
	EmoteDelay71 int16 `json:"emotedelay7_1,omitempty"`
	Emote71 int16 `json:"emote7_1,omitempty"`
	EmoteDelay72 int16 `json:"emotedelay7_2,omitempty"`
	Emote72 int16 `json:"emote7_2,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type NpcTextSlice []NpcText

type DBNpcText struct {
	bun.BaseModel `bun:"table:npc_text,alias:npc_text"`
	ID int `bun:"ID"`
	Text00 string `bun:"text0_0"`
	Text01 string `bun:"text0_1"`
	BroadcastTextID0 int `bun:"BroadcastTextID0"`
	Lang0 int8 `bun:"lang0"`
	Probability0 float64 `bun:"Probability0"`
	EmoteDelay00 int16 `bun:"EmoteDelay0_0"`
	Emote00 int16 `bun:"Emote0_0"`
	EmoteDelay01 int16 `bun:"EmoteDelay0_1"`
	Emote01 int16 `bun:"Emote0_1"`
	EmoteDelay02 int16 `bun:"EmoteDelay0_2"`
	Emote02 int16 `bun:"Emote0_2"`
	Text10 string `bun:"text1_0"`
	Text11 string `bun:"text1_1"`
	BroadcastTextID1 int `bun:"BroadcastTextID1"`
	Lang1 int8 `bun:"lang1"`
	Probability1 float64 `bun:"Probability1"`
	EmoteDelay10 int16 `bun:"EmoteDelay1_0"`
	Emote10 int16 `bun:"Emote1_0"`
	EmoteDelay11 int16 `bun:"EmoteDelay1_1"`
	Emote11 int16 `bun:"Emote1_1"`
	EmoteDelay12 int16 `bun:"EmoteDelay1_2"`
	Emote12 int16 `bun:"Emote1_2"`
	Text20 string `bun:"text2_0"`
	Text21 string `bun:"text2_1"`
	BroadcastTextID2 int `bun:"BroadcastTextID2"`
	Lang2 int8 `bun:"lang2"`
	Probability2 float64 `bun:"Probability2"`
	EmoteDelay20 int16 `bun:"EmoteDelay2_0"`
	Emote20 int16 `bun:"Emote2_0"`
	EmoteDelay21 int16 `bun:"EmoteDelay2_1"`
	Emote21 int16 `bun:"Emote2_1"`
	EmoteDelay22 int16 `bun:"EmoteDelay2_2"`
	Emote22 int16 `bun:"Emote2_2"`
	Text30 string `bun:"text3_0"`
	Text31 string `bun:"text3_1"`
	BroadcastTextID3 int `bun:"BroadcastTextID3"`
	Lang3 int8 `bun:"lang3"`
	Probability3 float64 `bun:"Probability3"`
	EmoteDelay30 int16 `bun:"EmoteDelay3_0"`
	Emote30 int16 `bun:"Emote3_0"`
	EmoteDelay31 int16 `bun:"EmoteDelay3_1"`
	Emote31 int16 `bun:"Emote3_1"`
	EmoteDelay32 int16 `bun:"EmoteDelay3_2"`
	Emote32 int16 `bun:"Emote3_2"`
	Text40 string `bun:"text4_0"`
	Text41 string `bun:"text4_1"`
	BroadcastTextID4 int `bun:"BroadcastTextID4"`
	Lang4 int8 `bun:"lang4"`
	Probability4 float64 `bun:"Probability4"`
	EmoteDelay40 int16 `bun:"EmoteDelay4_0"`
	Emote40 int16 `bun:"Emote4_0"`
	EmoteDelay41 int16 `bun:"EmoteDelay4_1"`
	Emote41 int16 `bun:"Emote4_1"`
	EmoteDelay42 int16 `bun:"EmoteDelay4_2"`
	Emote42 int16 `bun:"Emote4_2"`
	Text50 string `bun:"text5_0"`
	Text51 string `bun:"text5_1"`
	BroadcastTextID5 int `bun:"BroadcastTextID5"`
	Lang5 int8 `bun:"lang5"`
	Probability5 float64 `bun:"Probability5"`
	EmoteDelay50 int16 `bun:"EmoteDelay5_0"`
	Emote50 int16 `bun:"Emote5_0"`
	EmoteDelay51 int16 `bun:"EmoteDelay5_1"`
	Emote51 int16 `bun:"Emote5_1"`
	EmoteDelay52 int16 `bun:"EmoteDelay5_2"`
	Emote52 int16 `bun:"Emote5_2"`
	Text60 string `bun:"text6_0"`
	Text61 string `bun:"text6_1"`
	BroadcastTextID6 int `bun:"BroadcastTextID6"`
	Lang6 int8 `bun:"lang6"`
	Probability6 float64 `bun:"Probability6"`
	EmoteDelay60 int16 `bun:"EmoteDelay6_0"`
	Emote60 int16 `bun:"Emote6_0"`
	EmoteDelay61 int16 `bun:"EmoteDelay6_1"`
	Emote61 int16 `bun:"Emote6_1"`
	EmoteDelay62 int16 `bun:"EmoteDelay6_2"`
	Emote62 int16 `bun:"Emote6_2"`
	Text70 string `bun:"text7_0"`
	Text71 string `bun:"text7_1"`
	BroadcastTextID7 int `bun:"BroadcastTextID7"`
	Lang7 int8 `bun:"lang7"`
	Probability7 float64 `bun:"Probability7"`
	EmoteDelay70 int16 `bun:"EmoteDelay7_0"`
	Emote70 int16 `bun:"Emote7_0"`
	EmoteDelay71 int16 `bun:"EmoteDelay7_1"`
	Emote71 int16 `bun:"Emote7_1"`
	EmoteDelay72 int16 `bun:"EmoteDelay7_2"`
	Emote72 int16 `bun:"Emote7_2"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBNpcTextSlice []DBNpcText

func (entry *NpcText) ToDB() *DBNpcText {
	if entry == nil {
		return nil
	}
	return &DBNpcText{
		ID: entry.ID,
		Text00: entry.Text00,
		Text01: entry.Text01,
		BroadcastTextID0: entry.BroadcastTextID0,
		Lang0: entry.Lang0,
		Probability0: entry.Probability0,
		EmoteDelay00: entry.EmoteDelay00,
		Emote00: entry.Emote00,
		EmoteDelay01: entry.EmoteDelay01,
		Emote01: entry.Emote01,
		EmoteDelay02: entry.EmoteDelay02,
		Emote02: entry.Emote02,
		Text10: entry.Text10,
		Text11: entry.Text11,
		BroadcastTextID1: entry.BroadcastTextID1,
		Lang1: entry.Lang1,
		Probability1: entry.Probability1,
		EmoteDelay10: entry.EmoteDelay10,
		Emote10: entry.Emote10,
		EmoteDelay11: entry.EmoteDelay11,
		Emote11: entry.Emote11,
		EmoteDelay12: entry.EmoteDelay12,
		Emote12: entry.Emote12,
		Text20: entry.Text20,
		Text21: entry.Text21,
		BroadcastTextID2: entry.BroadcastTextID2,
		Lang2: entry.Lang2,
		Probability2: entry.Probability2,
		EmoteDelay20: entry.EmoteDelay20,
		Emote20: entry.Emote20,
		EmoteDelay21: entry.EmoteDelay21,
		Emote21: entry.Emote21,
		EmoteDelay22: entry.EmoteDelay22,
		Emote22: entry.Emote22,
		Text30: entry.Text30,
		Text31: entry.Text31,
		BroadcastTextID3: entry.BroadcastTextID3,
		Lang3: entry.Lang3,
		Probability3: entry.Probability3,
		EmoteDelay30: entry.EmoteDelay30,
		Emote30: entry.Emote30,
		EmoteDelay31: entry.EmoteDelay31,
		Emote31: entry.Emote31,
		EmoteDelay32: entry.EmoteDelay32,
		Emote32: entry.Emote32,
		Text40: entry.Text40,
		Text41: entry.Text41,
		BroadcastTextID4: entry.BroadcastTextID4,
		Lang4: entry.Lang4,
		Probability4: entry.Probability4,
		EmoteDelay40: entry.EmoteDelay40,
		Emote40: entry.Emote40,
		EmoteDelay41: entry.EmoteDelay41,
		Emote41: entry.Emote41,
		EmoteDelay42: entry.EmoteDelay42,
		Emote42: entry.Emote42,
		Text50: entry.Text50,
		Text51: entry.Text51,
		BroadcastTextID5: entry.BroadcastTextID5,
		Lang5: entry.Lang5,
		Probability5: entry.Probability5,
		EmoteDelay50: entry.EmoteDelay50,
		Emote50: entry.Emote50,
		EmoteDelay51: entry.EmoteDelay51,
		Emote51: entry.Emote51,
		EmoteDelay52: entry.EmoteDelay52,
		Emote52: entry.Emote52,
		Text60: entry.Text60,
		Text61: entry.Text61,
		BroadcastTextID6: entry.BroadcastTextID6,
		Lang6: entry.Lang6,
		Probability6: entry.Probability6,
		EmoteDelay60: entry.EmoteDelay60,
		Emote60: entry.Emote60,
		EmoteDelay61: entry.EmoteDelay61,
		Emote61: entry.Emote61,
		EmoteDelay62: entry.EmoteDelay62,
		Emote62: entry.Emote62,
		Text70: entry.Text70,
		Text71: entry.Text71,
		BroadcastTextID7: entry.BroadcastTextID7,
		Lang7: entry.Lang7,
		Probability7: entry.Probability7,
		EmoteDelay70: entry.EmoteDelay70,
		Emote70: entry.Emote70,
		EmoteDelay71: entry.EmoteDelay71,
		Emote71: entry.Emote71,
		EmoteDelay72: entry.EmoteDelay72,
		Emote72: entry.Emote72,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBNpcText) ToWeb() *NpcText {
	if entry == nil {
		return nil
	}
	return &NpcText{
		ID: entry.ID,
		Text00: entry.Text00,
		Text01: entry.Text01,
		BroadcastTextID0: entry.BroadcastTextID0,
		Lang0: entry.Lang0,
		Probability0: entry.Probability0,
		EmoteDelay00: entry.EmoteDelay00,
		Emote00: entry.Emote00,
		EmoteDelay01: entry.EmoteDelay01,
		Emote01: entry.Emote01,
		EmoteDelay02: entry.EmoteDelay02,
		Emote02: entry.Emote02,
		Text10: entry.Text10,
		Text11: entry.Text11,
		BroadcastTextID1: entry.BroadcastTextID1,
		Lang1: entry.Lang1,
		Probability1: entry.Probability1,
		EmoteDelay10: entry.EmoteDelay10,
		Emote10: entry.Emote10,
		EmoteDelay11: entry.EmoteDelay11,
		Emote11: entry.Emote11,
		EmoteDelay12: entry.EmoteDelay12,
		Emote12: entry.Emote12,
		Text20: entry.Text20,
		Text21: entry.Text21,
		BroadcastTextID2: entry.BroadcastTextID2,
		Lang2: entry.Lang2,
		Probability2: entry.Probability2,
		EmoteDelay20: entry.EmoteDelay20,
		Emote20: entry.Emote20,
		EmoteDelay21: entry.EmoteDelay21,
		Emote21: entry.Emote21,
		EmoteDelay22: entry.EmoteDelay22,
		Emote22: entry.Emote22,
		Text30: entry.Text30,
		Text31: entry.Text31,
		BroadcastTextID3: entry.BroadcastTextID3,
		Lang3: entry.Lang3,
		Probability3: entry.Probability3,
		EmoteDelay30: entry.EmoteDelay30,
		Emote30: entry.Emote30,
		EmoteDelay31: entry.EmoteDelay31,
		Emote31: entry.Emote31,
		EmoteDelay32: entry.EmoteDelay32,
		Emote32: entry.Emote32,
		Text40: entry.Text40,
		Text41: entry.Text41,
		BroadcastTextID4: entry.BroadcastTextID4,
		Lang4: entry.Lang4,
		Probability4: entry.Probability4,
		EmoteDelay40: entry.EmoteDelay40,
		Emote40: entry.Emote40,
		EmoteDelay41: entry.EmoteDelay41,
		Emote41: entry.Emote41,
		EmoteDelay42: entry.EmoteDelay42,
		Emote42: entry.Emote42,
		Text50: entry.Text50,
		Text51: entry.Text51,
		BroadcastTextID5: entry.BroadcastTextID5,
		Lang5: entry.Lang5,
		Probability5: entry.Probability5,
		EmoteDelay50: entry.EmoteDelay50,
		Emote50: entry.Emote50,
		EmoteDelay51: entry.EmoteDelay51,
		Emote51: entry.Emote51,
		EmoteDelay52: entry.EmoteDelay52,
		Emote52: entry.Emote52,
		Text60: entry.Text60,
		Text61: entry.Text61,
		BroadcastTextID6: entry.BroadcastTextID6,
		Lang6: entry.Lang6,
		Probability6: entry.Probability6,
		EmoteDelay60: entry.EmoteDelay60,
		Emote60: entry.Emote60,
		EmoteDelay61: entry.EmoteDelay61,
		Emote61: entry.Emote61,
		EmoteDelay62: entry.EmoteDelay62,
		Emote62: entry.Emote62,
		Text70: entry.Text70,
		Text71: entry.Text71,
		BroadcastTextID7: entry.BroadcastTextID7,
		Lang7: entry.Lang7,
		Probability7: entry.Probability7,
		EmoteDelay70: entry.EmoteDelay70,
		Emote70: entry.Emote70,
		EmoteDelay71: entry.EmoteDelay71,
		Emote71: entry.Emote71,
		EmoteDelay72: entry.EmoteDelay72,
		Emote72: entry.Emote72,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data NpcTextSlice) ToDB() DBNpcTextSlice {
	result := make(DBNpcTextSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBNpcTextSlice) ToWeb() NpcTextSlice {
	result := make(NpcTextSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
