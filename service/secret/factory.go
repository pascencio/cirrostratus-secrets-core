package secret

import (
	"cirrostratus-secrets-core/application"
	"cirrostratus-secrets-core/ports"
)

func Configure() {
	componentManager := application.GetInstance()
	componentManager.Register(ports.SecretServiceComponent, newService)
}

func newService(container application.ComponentContainer) (interface{}, error) {
	component, e := container.Get(ports.SecretRepositoryComponent)
	if e != nil {
		return nil, e
	}
	secretRepository := component.(ports.SecretRepository)
	component, e = container.Get(ports.KeyServiceComponent)
	if e != nil {
		return nil, e
	}
	keyService := component.(ports.KeyService)
	component, e = container.Get(ports.AuditoryServiceComponent)
	if e != nil {
		return nil, e
	}
	auditoryService := component.(ports.AuditoryService)
	var secretService ports.SecretService = service{
		secretRepository: secretRepository,
		keyService:       keyService,
		auditoryService:  auditoryService,
	}
	return secretService, nil
}
