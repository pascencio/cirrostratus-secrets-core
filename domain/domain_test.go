package domain_test

import (
	"cirrostratus-secrets-core/common"
	"cirrostratus-secrets-core/domain"
	"testing"
)

func TestSecretIDEncoding(t *testing.T) {
	name := "s3cr3t"
	createdBy := "dummy"
	createdAt := common.Epoch()
	secretID := &domain.SecretID{
		Name:      name,
		CreatedAt: createdAt.AsMillis(),
		CreatedBy: createdBy,
	}
	encoded, e := secretID.Encode()
	if e != nil {
		t.Log("Error al codificar ID")
		t.Fail()
	}
	t.Logf("SecretID códificado: [encoded='%s']", encoded)
	secretID, e = domain.DecodeSecret(encoded)
	if secretID.Name != name {
		t.Logf("El nombre debe ser igual a '%s'", name)
		t.Fail()
	}
	if secretID.CreatedAt != createdAt.AsMillis() {
		t.Logf("La fecha de creación debe ser igual a '%d'", createdAt)
		t.Fail()
	}
}

func TestIdentityIDEncoding(t *testing.T) {
	userName := "dummy"
	createdAt := common.Epoch()
	fullName := "Dummy User"
	identityID := domain.NewIdentityID(userName, createdAt.AsMillis(), fullName)
	encoded, e := identityID.Encode()
	if e != nil {
		t.Log("Error al crear identificación de usuario")
		t.Fail()
	}
	t.Logf("IdentityID codificado: [encoded='%s']", encoded)
	identityID, e = domain.DecodeIdentity(encoded)
	if e != nil {
		t.Error("Error al decodificar usuario", e)
		t.Fail()
	}
}

func TestSecretIDHash(t *testing.T) {
	name := "s3cr3t"
	createdAt := common.Epoch()
	createdBy := "dummy"
	secretID := &domain.SecretID{
		Name:      name,
		CreatedAt: createdAt.AsMillis(),
		CreatedBy: createdBy,
	}
	hash, e := secretID.Hash()
	if e != nil {
		t.Error("Error al generar hash de Secreto", e)
		t.Fail()
	}
	t.Logf("Hash generado: [hash='%s']", hash)
	if hash != domain.HashValues(name, createdAt.AsMillisString(), createdBy) {
		t.Fail()
	}
}

func TestIdentityIDHash(t *testing.T) {
	userName := "dummy"
	createdAt := common.Epoch()
	fullName := "Dummy User Of Santiago Pérez"
	identity := domain.NewIdentityID(userName, createdAt.AsMillis(), fullName)
	hash, e := identity.Hash()
	if e != nil {
		t.Error("Error al generar hash: ", e)
		t.Fail()
	}
	t.Logf("Hash generado: [hash='%s']", hash)
	if hash != domain.HashValues(userName, createdAt.AsMillisString(), fullName) {
		t.Fail()
	}
}
