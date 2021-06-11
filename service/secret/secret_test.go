package secret_test

import (
	"cirrostratus-secrets-core/application"
	"cirrostratus-secrets-core/common"
	"cirrostratus-secrets-core/domain"
	"cirrostratus-secrets-core/mock"
	"cirrostratus-secrets-core/ports"
	"cirrostratus-secrets-core/service/auditory"
	"cirrostratus-secrets-core/service/secret"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSecretCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mngr := application.GetInstance()
	defer ctrl.Finish()
	secret.Configure()
	auditory.Configure()
	mockSecretRepo := mockSecretRepo(ctrl, mngr)
	mockKeySrv(ctrl, mngr)
	mockAuditoryRepo := mockAuditoryRepo(ctrl, mngr)
	mockSecretRepo.EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(s domain.Secret) (domain.Secret, error) {
			return s, nil
		})
	mockAuditoryRepo.
		EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(a domain.Auditory) (domain.Auditory, error) {
			return a, nil
		})
	cmpt, e := mngr.GetComponent(ports.SecretServiceComponent)
	if e != nil {
		t.Fail()
	}
	iden := newIdentity()
	sec := newSecret("", "", "")
	secretSrv := cmpt.(ports.SecretService)
	secID, e := secretSrv.Create(sec, iden)
	if e != nil {
		t.Fail()
	}
	if secID.Name != sec.Name {
		t.Fail()
	}
}

func newIdentity() domain.IdentityID {
	return domain.IdentityID{
		UserName:  "Dummy",
		FullName:  "Admin Dummy",
		CreatedAt: common.Epoch().AsMillis(),
	}
}

func newSecret(name string, value string, publicKey string) domain.Secret {
	return domain.Secret{
		Name:        name,
		Value:       value,
		PublicKeyID: publicKey,
		Enabled:     true,
	}
}

func mockSecretRepo(ctrl *gomock.Controller, mngr application.ComponentManager) *mock.MockSecretRepository {
	m := mock.NewMockSecretRepository(ctrl)
	mngr.Register(ports.SecretRepositoryComponent, func(cc application.ComponentContainer) (interface{}, error) {
		return m, nil
	})
	return m
}

func mockKeySrv(ctrl *gomock.Controller, mngr application.ComponentManager) *mock.MockKeyService {
	m := mock.NewMockKeyService(ctrl)
	mngr.Register(ports.KeyServiceComponent, func(cc application.ComponentContainer) (interface{}, error) {
		return m, nil
	})
	return m
}

func mockAuditoryRepo(ctrl *gomock.Controller, mngr application.ComponentManager) *mock.MockAuditoryRepository {
	m := mock.NewMockAuditoryRepository(ctrl)
	mngr.Register(ports.AuditoryRepositoryComponent, func(cc application.ComponentContainer) (interface{}, error) {
		return m, nil
	})
	return m
}
