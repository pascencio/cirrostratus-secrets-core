package domain

import "fmt"

// SecretID Identificador generado a través de atributos de la structura
type SecretID struct {
	Name      string
	CreatedAt int64
	CreatedBy string
}

// Encode Genera identificador único del Secreto
func (s SecretID) Encode() (string, error) {
	// FIXME: Buscar algoritmo de codificación
	return fmt.Sprintf("%s:%s:%d", s.Name, s.CreatedBy, s.CreatedAt), nil
}

// Decode Crea instancia de ID a partir de una cadena de texto
func (s SecretID) Decode(r string) (ID, error) {
	// FIXME: Buscar algoritmo de decodificación
	return SecretID{}, nil
}

// Secret Contiene datos del secreto
type Secret struct {
	ID             string
	Name           string
	EncryptedValue string
	KeyID          string
	Auditory       Auditory
	Enabled        bool
}
