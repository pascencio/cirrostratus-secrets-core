package auditory

import (
	"cirrostratus-secrets-core/common"
	"cirrostratus-secrets-core/domain"
	"cirrostratus-secrets-core/ports"

	"github.com/sirupsen/logrus"
)

func emptyAuditory() domain.Auditory {
	return emptyAuditory()
}

type service struct {
	auditoryRepository ports.AuditoryRepository
}

func (s service) New(identityID domain.IdentityID, action domain.Action) (domain.Auditory, error) {
	createdBy, e := identityID.Hash()
	if e != nil {
		return emptyAuditory(), e
	}

	return domain.Auditory{
		ID:         common.UUID(),
		LastAction: action.String(),
		CreatedBy:  createdBy,
		CreatedAt:  common.Epoch().AsMillis(),
	}, nil
}

func (s service) NewFrom(auditory domain.Auditory, identityID domain.IdentityID, action domain.Action) (domain.Auditory, error) {
	updatedBy, e := identityID.Hash()
	if e != nil {
		return emptyAuditory(), e
	}
	newAuditory := domain.Auditory{
		ID:         common.UUID(),
		LastAction: action.Action,
		CreatedBy:  auditory.CreatedBy,
		CreatedAt:  auditory.CreatedAt,
		UpdatedAt:  common.Epoch().AsMillis(),
		UpdatedBy:  updatedBy,
	}
	return newAuditory, nil
}

func (s service) Add(auditory domain.Auditory) {
	if _, e := s.auditoryRepository.Create(auditory); e != nil {
		logrus.Warn("Error al crear auditoria: ", e)
	}
}
