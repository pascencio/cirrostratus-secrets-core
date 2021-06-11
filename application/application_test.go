package application_test

import (
	"cirrostratus-secrets-core/application"
	"testing"
)

func TestComponentManager(t *testing.T) {
	m := application.GetInstance()
	m.Register("mockComponent", func(container application.ComponentContainer) (interface{}, error) {
		m, e := container.Get("mockDependency")
		if e != nil {
			return nil, e
		}
		return mockComponent{MockDependency: m.(mockDependency)}, nil
	})
	m.Register("mockDependency", func(container application.ComponentContainer) (interface{}, error) {
		return mockDependency{}, nil
	})
	_, e := m.GetComponent("mockComponent")
	if e != nil {
		t.Log("Componente no encontrado")
		t.Fail()
	}
}

type mockDependency struct{}
type mockComponent struct {
	MockDependency mockDependency
}
