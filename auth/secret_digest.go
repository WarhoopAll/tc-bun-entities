package model

import "github.com/uptrace/bun"

type SecretDigest struct {
	Id int `json:"id,omitempty"`
	Digest string `json:"digest,omitempty"`
}

type SecretDigestSlice []SecretDigest

type DBSecretDigest struct {
	bun.BaseModel `bun:"table:secret_digest,alias:secret_digest"`
	Id int `bun:"id"`
	Digest string `bun:"digest"`
}

type DBSecretDigestSlice []DBSecretDigest

func (entry *SecretDigest) ToDB() *DBSecretDigest {
	if entry == nil {
		return nil
	}
	return &DBSecretDigest{
		Id: entry.Id,
		Digest: entry.Digest,
	}
}

func (entry *DBSecretDigest) ToWeb() *SecretDigest {
	if entry == nil {
		return nil
	}
	return &SecretDigest{
		Id: entry.Id,
		Digest: entry.Digest,
	}
}

func (data SecretDigestSlice) ToDB() DBSecretDigestSlice {
	result := make(DBSecretDigestSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSecretDigestSlice) ToWeb() SecretDigestSlice {
	result := make(SecretDigestSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
