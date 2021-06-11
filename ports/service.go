package ports

import (
	"cirrostratus-secrets-core/command"
	"cirrostratus-secrets-core/domain"
)

// SecretServiceComponent Nombre del componente Service
const SecretServiceComponent = "SecretService"
const KeyServiceComponent = "KeyService"
const AuditoryServiceComponent = "AuditoryService"

// SecretService Lógica de negocios del manejo de secretos
type SecretService interface {
	Create(domain.Secret, domain.IdentityID) (domain.SecretID, error)
	Get(domain.SecretID, domain.IdentityID) (domain.Secret, error)
	GetAll(command.PageRequest, domain.IdentityID) ([]domain.Secret, error)
	Update(domain.SecretID, domain.Secret, domain.IdentityID) (domain.Secret, error)
	Delete(domain.SecretID, domain.IdentityID) error
	Reveal(domain.SecretID, domain.IdentityID) (string, error)
}

// KeyService Lógica de negocio para el manejo de llaves
type KeyService interface {
	Get(domain.KeyID, domain.IdentityID) (domain.Key, error)
	Create(domain.Key, domain.IdentityID) (domain.KeyID, error)
}

type AuditoryService interface {
	New(domain.IdentityID, domain.Action) (domain.Auditory, error)
	NewFrom(domain.Auditory, domain.IdentityID, domain.Action) (domain.Auditory, error)
	Add(domain.Auditory)
}
