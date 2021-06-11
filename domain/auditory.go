package domain

import (
	"cirrostratus-secrets-core/common"
)

// Auditory Auditoria de dominios de la aplicación
type Auditory struct {
	ID         string
	LastAction string
	CreatedBy  string
	UpdatedBy  string
	UpdatedAt  int64
	CreatedAt  int64
}

// NewAuditory Crea una instancia de auditoría para un dominio nuevo
func NewAuditory(identity IdentityID, action Action) (Auditory, error) {
	createdBy, e := identity.Hash()
	if e != nil {
		return Auditory{}, e
	}

	return Auditory{
		LastAction: action.Action,
		CreatedBy:  createdBy,
		CreatedAt:  common.Epoch().AsMillis(),
	}, nil
}
