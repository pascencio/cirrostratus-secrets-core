package domain

import "fmt"

// KeyID Utiliza datos de la llave para generar una clave única
type KeyID struct {
	Name      string
	Algorithm string
	CreatedBy string
	CreatedAt int64
}

// Encode Codifica los datos del ID de la Llave
func (k KeyID) Encode() (string, error) {
	// FIXME: Buscar algoritmo de codificación
	return fmt.Sprintf("%s:%s:%s:%d", k.Name, k.Algorithm, k.CreatedBy, k.CreatedAt), nil
}

// Decode Crea instancia del ID de la Llave a partir de una cadena de texto
func (k KeyID) Decode(r string) (ID, error) {
	// FIXME: Buscar algoritmo de decodificación
	return KeyID{}, nil
}

// Key Almacena datos de la llave
type Key struct {
	ID        string
	Name      string
	Algorithm string
	PublicKey string
	Auditory  Auditory
}
