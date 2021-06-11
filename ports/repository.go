package ports

import (
	"cirrostratus-secrets-core/command"
	"cirrostratus-secrets-core/domain"
)

// SecretRepositoryComponent  Nombre del componente SecretRepository
const SecretRepositoryComponent = "SecretRepository"
const AuditoryRepositoryComponent = "AuditoryRepository"

// SecretRepository Realiza acciones sobre un puerto de salida a partir de la lógica de negocio de los Secretos
type SecretRepository interface {
	GetAll(command.PageRequest) ([]domain.Secret, error)
	Get(string) (domain.Secret, error)
	Create(domain.Secret) (domain.Secret, error)
	Delete(string) error
	Update(string, domain.Secret) error
}

// KeyRepository Almacena datos de los secretos
type KeyRepository interface {
	GetAll(command.PageRequest) ([]domain.Key, error)
	Get(string) (domain.Key, error)
	Create(string) (domain.Key, error)
	Delete(string) error
	Update(string, domain.Key) error
}

// AuditoryRepository Almacena datos de auditoria
type AuditoryRepository interface {
	Create(domain.Auditory) (domain.Auditory, error)
	Get(string) (domain.Auditory, error)
	GetAll(command.PageRequest) ([]domain.Auditory, error)
}
