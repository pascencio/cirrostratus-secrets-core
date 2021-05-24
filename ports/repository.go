package ports

import (
	"cirrostratus-secrets-core/command"
	"cirrostratus-secrets-core/domain"
)

// SecretRepository Realiza acciones sobre un puerto de salida a partir de la lógica de negocio de los Secretos
type SecretRepository interface {
	GetAll(command.PageRequest) ([]domain.Secret, error)
	Get(string) (domain.Secret, error)
	Create(domain.Secret) (domain.Secret, error)
	Delete(string) error
	Update(domain.Secret) error
}

// KeyRepository Almacena datos de los secretos
type KeyRepository interface {
	GetAll(command.PageRequest) ([]domain.Key, error)
	Get(string) (domain.Key, error)
	Create(domain.Key) (domain.Key, error)
	Delete(string) error
	Update(domain.Secret) error
}
