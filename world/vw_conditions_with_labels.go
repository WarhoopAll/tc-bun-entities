package model

import "github.com/uptrace/bun"

type VwConditionsWithLabels struct {
	SourceTypeOrReferenceId string `json:"sourcetypeorreferenceid,omitempty"`
	SourceGroup int `json:"sourcegroup,omitempty"`
	SourceEntry int `json:"sourceentry,omitempty"`
	SourceId int `json:"sourceid,omitempty"`
	ElseGroup int `json:"elsegroup,omitempty"`
	ConditionTypeOrReference string `json:"conditiontypeorreference,omitempty"`
	ConditionTarget int8 `json:"conditiontarget,omitempty"`
	ConditionValue1 int `json:"conditionvalue1,omitempty"`
	ConditionValue2 int `json:"conditionvalue2,omitempty"`
	ConditionValue3 int `json:"conditionvalue3,omitempty"`
	NegativeCondition int8 `json:"negativecondition,omitempty"`
	ErrorType int `json:"errortype,omitempty"`
	ErrorTextId int `json:"errortextid,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type VwConditionsWithLabelsSlice []VwConditionsWithLabels

type DBVwConditionsWithLabels struct {
	bun.BaseModel `bun:"table:vw_conditions_with_labels,alias:vw_conditions_with_labels"`
	SourceTypeOrReferenceId string `bun:"SourceTypeOrReferenceId"`
	SourceGroup int `bun:"SourceGroup"`
	SourceEntry int `bun:"SourceEntry"`
	SourceId int `bun:"SourceId"`
	ElseGroup int `bun:"ElseGroup"`
	ConditionTypeOrReference string `bun:"ConditionTypeOrReference"`
	ConditionTarget int8 `bun:"ConditionTarget"`
	ConditionValue1 int `bun:"ConditionValue1"`
	ConditionValue2 int `bun:"ConditionValue2"`
	ConditionValue3 int `bun:"ConditionValue3"`
	NegativeCondition int8 `bun:"NegativeCondition"`
	ErrorType int `bun:"ErrorType"`
	ErrorTextId int `bun:"ErrorTextId"`
	ScriptName string `bun:"ScriptName"`
	Comment string `bun:"Comment"`
}

type DBVwConditionsWithLabelsSlice []DBVwConditionsWithLabels

func (entry *VwConditionsWithLabels) ToDB() *DBVwConditionsWithLabels {
	if entry == nil {
		return nil
	}
	return &DBVwConditionsWithLabels{
		SourceTypeOrReferenceId: entry.SourceTypeOrReferenceId,
		SourceGroup: entry.SourceGroup,
		SourceEntry: entry.SourceEntry,
		SourceId: entry.SourceId,
		ElseGroup: entry.ElseGroup,
		ConditionTypeOrReference: entry.ConditionTypeOrReference,
		ConditionTarget: entry.ConditionTarget,
		ConditionValue1: entry.ConditionValue1,
		ConditionValue2: entry.ConditionValue2,
		ConditionValue3: entry.ConditionValue3,
		NegativeCondition: entry.NegativeCondition,
		ErrorType: entry.ErrorType,
		ErrorTextId: entry.ErrorTextId,
		ScriptName: entry.ScriptName,
		Comment: entry.Comment,
	}
}

func (entry *DBVwConditionsWithLabels) ToWeb() *VwConditionsWithLabels {
	if entry == nil {
		return nil
	}
	return &VwConditionsWithLabels{
		SourceTypeOrReferenceId: entry.SourceTypeOrReferenceId,
		SourceGroup: entry.SourceGroup,
		SourceEntry: entry.SourceEntry,
		SourceId: entry.SourceId,
		ElseGroup: entry.ElseGroup,
		ConditionTypeOrReference: entry.ConditionTypeOrReference,
		ConditionTarget: entry.ConditionTarget,
		ConditionValue1: entry.ConditionValue1,
		ConditionValue2: entry.ConditionValue2,
		ConditionValue3: entry.ConditionValue3,
		NegativeCondition: entry.NegativeCondition,
		ErrorType: entry.ErrorType,
		ErrorTextId: entry.ErrorTextId,
		ScriptName: entry.ScriptName,
		Comment: entry.Comment,
	}
}

func (data VwConditionsWithLabelsSlice) ToDB() DBVwConditionsWithLabelsSlice {
	result := make(DBVwConditionsWithLabelsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVwConditionsWithLabelsSlice) ToWeb() VwConditionsWithLabelsSlice {
	result := make(VwConditionsWithLabelsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
