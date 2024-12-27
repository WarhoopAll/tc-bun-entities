package model

import "github.com/uptrace/bun"

type GroupMember struct {
	Guid int `json:"guid,omitempty"`
	MemberGuid int `json:"memberguid,omitempty"`
	MemberFlags int8 `json:"memberflags,omitempty"`
	Subgroup int8 `json:"subgroup,omitempty"`
	Roles int8 `json:"roles,omitempty"`
}

type GroupMemberSlice []GroupMember

type DBGroupMember struct {
	bun.BaseModel `bun:"table:group_member,alias:group_member"`
	Guid int `bun:"guid"`
	MemberGuid int `bun:"memberGuid"`
	MemberFlags int8 `bun:"memberFlags"`
	Subgroup int8 `bun:"subgroup"`
	Roles int8 `bun:"roles"`
}

type DBGroupMemberSlice []DBGroupMember

func (entry *GroupMember) ToDB() *DBGroupMember {
	if entry == nil {
		return nil
	}
	return &DBGroupMember{
		Guid: entry.Guid,
		MemberGuid: entry.MemberGuid,
		MemberFlags: entry.MemberFlags,
		Subgroup: entry.Subgroup,
		Roles: entry.Roles,
	}
}

func (entry *DBGroupMember) ToWeb() *GroupMember {
	if entry == nil {
		return nil
	}
	return &GroupMember{
		Guid: entry.Guid,
		MemberGuid: entry.MemberGuid,
		MemberFlags: entry.MemberFlags,
		Subgroup: entry.Subgroup,
		Roles: entry.Roles,
	}
}

func (data GroupMemberSlice) ToDB() DBGroupMemberSlice {
	result := make(DBGroupMemberSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGroupMemberSlice) ToWeb() GroupMemberSlice {
	result := make(GroupMemberSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
