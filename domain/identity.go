package domain

import (
	"cirrostratus-secrets-core/common"
)

const identityIDSeparator = "/"

// IdentityID Contiene información que identifica a un usuario del sistema
type IdentityID struct {
	UserName  string
	FullName  string
	CreatedAt int64
}

// Encode Genera identificador único de la Identidad
func (i IdentityID) Encode() (string, error) {
	epoch := common.EpochFrom(i.CreatedAt, common.InMillis)
	encoded := EncodeValuesToB64(identityIDSeparator, i.UserName, epoch.AsMillisString(), i.FullName)
	return encoded, nil
}

// Hash Genera ID a partir de atributos de la identidad
func (i IdentityID) Hash() (string, error) {
	epoch := common.EpochFrom(i.CreatedAt, common.InMillis)
	return HashValues(i.UserName, epoch.AsMillisString(), i.FullName), nil
}

// DecodeIdentity Genera instancia de la Identidad a partir de una cadena de texto
func DecodeIdentity(r string) (IdentityID, error) {
	values, e := Decode64ToValues(identityIDSeparator, r)
	if e != nil {
		return IdentityID{}, e
	}
	userName := values[0]
	createdAt, e := common.ParseEpoch(values[1], common.InMillis)
	if e != nil {
		return IdentityID{}, e
	}
	fullName := values[2]
	return IdentityID{
		UserName:  userName,
		CreatedAt: createdAt.AsMillis(),
		FullName:  fullName,
	}, nil
}

// NewIdentityID Crea un nuevo ID de Identidad
func NewIdentityID(userName string, createdAt int64, fullName string) IdentityID {
	return IdentityID{
		UserName:  userName,
		CreatedAt: createdAt,
		FullName:  fullName,
	}
}

type Identity struct {
	ID           IdentityID
	UserName     string
	FullName     string
	Auditory     Auditory
	Enabled      bool
	PrivateKeyID KeyID
	Permissions  []Permission
}
