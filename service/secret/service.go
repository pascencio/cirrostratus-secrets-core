package secret

import (
	"cirrostratus-secrets-core/command"
	"cirrostratus-secrets-core/domain"
	"cirrostratus-secrets-core/ports"

	"github.com/sirupsen/logrus"
)

const createAction = "create"
const getAction = "get"
const getAllAction = "get-all"
const deleteAction = "delete"
const updateAction = "update"
const revealAction = "reveal"
const secretContext = "secret"

type service struct {
	secretRepository ports.SecretRepository
	keyService       ports.KeyService
	auditoryService  ports.AuditoryService
}

// Create Crea un nuevo Secreto
func (s service) Create(secret domain.Secret, identityID domain.IdentityID) (domain.SecretID, error) {
	action := domain.Action{Context: secretContext, Action: createAction}
	auditory, e := s.auditoryService.New(identityID, action)
	if e != nil {
		return domain.SecretID{}, e
	}
	logrus.WithFields(logrus.Fields{
		"action":     action.String(),
		"identityID": auditory.CreatedBy,
	}).Info("Creando secreto")
	id := domain.SecretID{
		Name:      secret.Name,
		CreatedAt: auditory.CreatedAt,
		CreatedBy: auditory.CreatedBy,
	}
	hash, e := id.Hash()
	if e != nil {
		return domain.SecretID{}, e
	}
	secret.ID = hash
	secret.Auditory = auditory
	// FIXME: Implementar lógica para desencryptar
	if _, e = s.secretRepository.Create(secret); e != nil {
		return domain.SecretID{}, e
	}
	logrus.WithFields(logrus.Fields{
		"secretID":   hash,
		"identityID": auditory.CreatedBy,
	}).Info("Secreto creado correctamente")
	s.auditoryService.Add(auditory)
	return id, nil
}

func (s service) Get(secretID domain.SecretID, identityID domain.IdentityID) (domain.Secret, error) {
	action := domain.Action{Context: secretContext, Action: getAction}
	auditory, e := s.auditoryService.New(identityID, action)
	if e != nil {
		return domain.Secret{}, e
	}
	id, e := secretID.Hash()
	if e != nil {
		return domain.Secret{}, e
	}
	secret, e := s.secretRepository.Get(id)
	if e != nil {
		return domain.Secret{}, e
	}
	s.auditoryService.Add(auditory)
	return secret, nil
}

func (s service) GetAll(paging command.PageRequest, identityID domain.IdentityID) ([]domain.Secret, error) {
	action := domain.Action{Context: secretContext, Action: getAllAction}
	auditory, e := s.auditoryService.New(identityID, action)
	if e != nil {
		return nil, e
	}
	secrets, e := s.secretRepository.GetAll(paging)
	if e != nil {
		return nil, e
	}
	s.auditoryService.Add(auditory)
	return secrets, nil
}

func (s service) Update(secretID domain.SecretID, secret domain.Secret, identityID domain.IdentityID) (domain.Secret, error) {
	action := domain.Action{Context: secretContext, Action: updateAction}
	auditory, e := s.auditoryService.NewFrom(secret.Auditory, identityID, action)
	if e != nil {
		return domain.Secret{}, e
	}
	secret.Auditory = auditory
	hash, e := secretID.Hash()
	if e != nil {
		return domain.Secret{}, e
	}
	e = s.secretRepository.Update(hash, secret)
	if e != nil {
		return domain.Secret{}, e
	}
	s.auditoryService.Add(auditory)
	return secret, nil
}

func (s service) Delete(secretID domain.SecretID, identityID domain.IdentityID) error {
	action := domain.Action{Context: secretContext, Action: deleteAction}
	auditory, e := s.auditoryService.New(identityID, action)
	if e != nil {
		return e
	}
	hash, e := secretID.Hash()
	if e != nil {
		return e
	}
	e = s.secretRepository.Delete(hash)
	if e != nil {
		return e
	}
	s.auditoryService.Add(auditory)
	return nil
}

func (s service) Reveal(secretID domain.SecretID, identityID domain.IdentityID) (string, error) {
	action := domain.Action{Context: secretContext, Action: revealAction}
	auditory, e := s.auditoryService.New(identityID, action)
	if e != nil {
		return "", e
	}
	hash, e := secretID.Hash()
	if e != nil {
		return "", e
	}
	_, e = s.secretRepository.Get(hash)
	if e != nil {
		return "", e
	}
	// FIXME: Implementar lógica para desencryptar
	s.auditoryService.Add(auditory)
	return "", e
}
