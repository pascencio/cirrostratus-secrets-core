package domain

import (
	"cirrostratus-secrets-core/common"
)

const secretIDSeparator = "/"

// SecretID Identificador generado a través de atributos de la structura
type SecretID struct {
	Name      string
	CreatedAt int64
	CreatedBy string
}

// Hash Genera identificado único del ID del secreto
func (s SecretID) Hash() (string, error) {
	epoch := common.EpochFrom(s.CreatedAt, common.InMillis)
	return HashValues(s.Name, epoch.AsMillisString(), s.CreatedBy), nil
}

// Encode Codifica en base 64 atributos del ID del secreto
func (s SecretID) Encode() (string, error) {
	epoch := common.EpochFrom(s.CreatedAt, common.InMillis)
	encoded := EncodeValuesToB64(secretIDSeparator, s.Name, epoch.AsMillisString(), s.CreatedBy)
	return encoded, nil
}

// DecodeSecret Decodifica string hexadecimal y genera instancia de Secreto
func DecodeSecret(r string) (*SecretID, error) {
	parts, e := Decode64ToValues(secretIDSeparator, r)
	if e != nil {
		return nil, e
	}
	createdAt, e := common.ParseEpoch(parts[1], common.InMillis)
	if e != nil {
		return nil, e
	}
	name := parts[0]
	createdBy := parts[2]
	return &SecretID{
		Name:      name,
		CreatedAt: createdAt.AsMillis(),
		CreatedBy: createdBy,
	}, nil
}

// Secret Contiene datos del secreto
type Secret struct {
	ID          string
	Name        string
	Value       string
	PublicKeyID string
	Auditory    Auditory
	Enabled     bool
}
