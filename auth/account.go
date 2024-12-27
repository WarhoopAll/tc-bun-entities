package model

import "github.com/uptrace/bun"

type Account struct {
	Id int `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Salt []byte `json:"salt,omitempty"`
	Verifier []byte `json:"verifier,omitempty"`
	SessionKeyAuth []byte `json:"session_key_auth,omitempty"`
	SessionKeyBnet []byte `json:"session_key_bnet,omitempty"`
	TotpSecret []byte `json:"totp_secret,omitempty"`
	Email string `json:"email,omitempty"`
	RegMail string `json:"reg_mail,omitempty"`
	Joindate time.Time `json:"joindate,omitempty"`
	LastIp string `json:"last_ip,omitempty"`
	LastAttemptIp string `json:"last_attempt_ip,omitempty"`
	FailedLogins int `json:"failed_logins,omitempty"`
	Locked int8 `json:"locked,omitempty"`
	LockCountry string `json:"lock_country,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty"`
	Online int8 `json:"online,omitempty"`
	Expansion int8 `json:"expansion,omitempty"`
	Mutetime int `json:"mutetime,omitempty"`
	Mutereason string `json:"mutereason,omitempty"`
	Muteby string `json:"muteby,omitempty"`
	Locale int8 `json:"locale,omitempty"`
	Os string `json:"os,omitempty"`
	TimezoneOffset int16 `json:"timezone_offset,omitempty"`
	Recruiter int `json:"recruiter,omitempty"`
}

type AccountSlice []Account

type DBAccount struct {
	bun.BaseModel `bun:"table:account,alias:account"`
	Id int `bun:"id"`
	Username string `bun:"username"`
	Salt []byte `bun:"salt"`
	Verifier []byte `bun:"verifier"`
	SessionKeyAuth []byte `bun:"session_key_auth"`
	SessionKeyBnet []byte `bun:"session_key_bnet"`
	TotpSecret []byte `bun:"totp_secret"`
	Email string `bun:"email"`
	RegMail string `bun:"reg_mail"`
	Joindate time.Time `bun:"joindate"`
	LastIp string `bun:"last_ip"`
	LastAttemptIp string `bun:"last_attempt_ip"`
	FailedLogins int `bun:"failed_logins"`
	Locked int8 `bun:"locked"`
	LockCountry string `bun:"lock_country"`
	LastLogin time.Time `bun:"last_login"`
	Online int8 `bun:"online"`
	Expansion int8 `bun:"expansion"`
	Mutetime int `bun:"mutetime"`
	Mutereason string `bun:"mutereason"`
	Muteby string `bun:"muteby"`
	Locale int8 `bun:"locale"`
	Os string `bun:"os"`
	TimezoneOffset int16 `bun:"timezone_offset"`
	Recruiter int `bun:"recruiter"`
}

type DBAccountSlice []DBAccount

func (entry *Account) ToDB() *DBAccount {
	if entry == nil {
		return nil
	}
	return &DBAccount{
		Id: entry.Id,
		Username: entry.Username,
		Salt: entry.Salt,
		Verifier: entry.Verifier,
		SessionKeyAuth: entry.SessionKeyAuth,
		SessionKeyBnet: entry.SessionKeyBnet,
		TotpSecret: entry.TotpSecret,
		Email: entry.Email,
		RegMail: entry.RegMail,
		Joindate: entry.Joindate,
		LastIp: entry.LastIp,
		LastAttemptIp: entry.LastAttemptIp,
		FailedLogins: entry.FailedLogins,
		Locked: entry.Locked,
		LockCountry: entry.LockCountry,
		LastLogin: entry.LastLogin,
		Online: entry.Online,
		Expansion: entry.Expansion,
		Mutetime: entry.Mutetime,
		Mutereason: entry.Mutereason,
		Muteby: entry.Muteby,
		Locale: entry.Locale,
		Os: entry.Os,
		TimezoneOffset: entry.TimezoneOffset,
		Recruiter: entry.Recruiter,
	}
}

func (entry *DBAccount) ToWeb() *Account {
	if entry == nil {
		return nil
	}
	return &Account{
		Id: entry.Id,
		Username: entry.Username,
		Salt: entry.Salt,
		Verifier: entry.Verifier,
		SessionKeyAuth: entry.SessionKeyAuth,
		SessionKeyBnet: entry.SessionKeyBnet,
		TotpSecret: entry.TotpSecret,
		Email: entry.Email,
		RegMail: entry.RegMail,
		Joindate: entry.Joindate,
		LastIp: entry.LastIp,
		LastAttemptIp: entry.LastAttemptIp,
		FailedLogins: entry.FailedLogins,
		Locked: entry.Locked,
		LockCountry: entry.LockCountry,
		LastLogin: entry.LastLogin,
		Online: entry.Online,
		Expansion: entry.Expansion,
		Mutetime: entry.Mutetime,
		Mutereason: entry.Mutereason,
		Muteby: entry.Muteby,
		Locale: entry.Locale,
		Os: entry.Os,
		TimezoneOffset: entry.TimezoneOffset,
		Recruiter: entry.Recruiter,
	}
}

func (data AccountSlice) ToDB() DBAccountSlice {
	result := make(DBAccountSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAccountSlice) ToWeb() AccountSlice {
	result := make(AccountSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
