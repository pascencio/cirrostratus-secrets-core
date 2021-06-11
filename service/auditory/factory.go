package auditory

import (
	"cirrostratus-secrets-core/application"
	"cirrostratus-secrets-core/ports"
)

func Configure() {
	componentManager := application.GetInstance()
	componentManager.Register(ports.AuditoryServiceComponent, newService)
}

func newService(container application.ComponentContainer) (interface{}, error) {
	component, e := container.Get(ports.AuditoryRepositoryComponent)
	if e != nil {
		return nil, e
	}
	auditoryRepository := component.(ports.AuditoryRepository)
	var auditoryService ports.AuditoryService = service{
		auditoryRepository: auditoryRepository,
	}
	return auditoryService, nil
}
