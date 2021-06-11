package domain

import (
	"cirrostratus-secrets-core/common"
)

const keyIDSeparator = "/"

// KeyID Utiliza datos de la llave para generar una clave única
type KeyID struct {
	Name      string
	Algorithm string
	CreatedBy string
	CreatedAt int64
}

// Hash Genera identificador único a partir de atributos del ID
func (k KeyID) Hash() (string, error) {
	epoch := common.EpochFrom(k.CreatedAt, common.InMillis)
	return HashValues(k.Name, k.Algorithm, epoch.AsMillisString(), k.CreatedBy), nil
}

// Encode Codifica los datos del ID de la Llave
func (k KeyID) Encode() (string, error) {
	epoch := common.EpochFrom(k.CreatedAt, common.InMillis)
	return EncodeValuesToB64(keyIDSeparator, k.Name, k.Algorithm, k.CreatedBy, epoch.AsMillisString()), nil
}

// DecodeKey Crea instancia del ID de la Llave a partir de una cadena de texto
func DecodeKey(r string) (KeyID, error) {
	values, e := Decode64ToValues(keyIDSeparator, r)
	if e != nil {
		return KeyID{}, e
	}
	name := values[0]
	algorithm := values[1]
	createdBy := values[2]
	createdAt, e := common.ParseEpoch(values[3], common.InMillis)
	if e != nil {
		return KeyID{}, e
	}
	return KeyID{
		Name:      name,
		Algorithm: algorithm,
		CreatedBy: createdBy,
		CreatedAt: createdAt.AsMillis(),
	}, nil
}

// Key Almacena datos de la llave
type Key struct {
	ID        string
	Name      string
	Algorithm string
	Value     string
	Private   bool
	Auditory  Auditory
}
