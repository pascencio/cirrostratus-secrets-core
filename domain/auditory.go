package domain

import "fmt"

// IdentityID Contiene información que identifica a un usuario del sistema
type IdentityID struct {
	UserName    string
	CreatedAt   int64
	Permissions []string // TODO: Ver si los permisos podrían ser una estructura también
}

// Encode Genera identificador único de la Identidad
func (i IdentityID) Encode() (string, error) {
	// FIXME: Buscar algoritmo de codificación
	return fmt.Sprintf("%s:%d:%s", i.UserName, i.CreatedAt, i.Permissions), nil
}

// Decode Genera instancia de la Identidad a partir de una cadena de texto
func (i IdentityID) Decode(r string) (ID, error) {
	// FIXME: Buscar algoritmo de decodificación
	return IdentityID{}, nil
}

// Auditory Auditoria de dominios de la aplicación
type Auditory struct {
	CreatedBy string
	UpdatedBy string
	UpdatedAt int64
	CreatedAt int64
}
