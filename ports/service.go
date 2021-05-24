package ports

import (
	"cirrostratus-secrets-core/command"
	"cirrostratus-secrets-core/domain"
)

// SecretService Lógica de negocios del manejo de secretos
type SecretService interface {
	Create(domain.Secret) (domain.SecretID, error)
	Get(domain.SecretID) (domain.Secret, error)
	GetAll(command.PageRequest) ([]domain.Secret, error)
	Update(domain.SecretID, domain.Secret) (domain.Secret, error)
	Delete(domain.SecretID) error
	Reveal(domain.SecretID) (string, error)
}
