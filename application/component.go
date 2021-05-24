package application

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Component Definición de un componente de la aplicación. Usado para administrar su instancia
type Component interface {
	GetInstance(ComponentContainer) interface{}
}

// ComponentContainer Almacena en "Cache" las instancias de los componentes de la aplicación
type ComponentContainer struct {
	Factories map[string]ComponentFactory
	Cache     map[string]interface{}
}

func (c *ComponentContainer) add(name string, factory ComponentFactory) error {
	if c.Factories == nil {
		return fmt.Errorf("No se asigno un valor al atributo Factories")
	}
	c.Factories[name] = factory
	return nil
}

// Get Obtiene instancia de un componente almacenado en el contenedor
func (c ComponentContainer) Get(name string) (interface{}, error) {
	if instance, ok := c.Cache[name]; ok {
		return instance, nil
	}
	instance, e := c.Factories[name](c)
	if e != nil {
		return nil, e
	}
	c.Cache[name] = instance

	return instance, nil
}

// ComponentFactory Crea una instancia de un componente
type ComponentFactory func(ComponentContainer) (interface{}, error)

// ComponentManager Administra e instancia los componentes
type ComponentManager interface {
	Register(string, Component) error
	GetInstance(string) (interface{}, error)
}

// DefaultComponentManager Implementación por defecto de ComponentManager
type DefaultComponentManager struct {
	Container ComponentContainer
}

// Register Registra un componente en el contenedor
func (c *DefaultComponentManager) Register(name string, factory ComponentFactory) error {
	logrus.WithFields(
		logrus.Fields{
			"componentName": name,
		}).Info("Registrando componente")
	if e := c.Container.add(name, factory); e != nil {
		return e
	}
	logrus.WithFields(
		logrus.Fields{
			"componentName": name,
		},
	).Info("Componente registrado")
	return nil
}

//GetInstance Retorna un instancia del componente registrado
func (c *DefaultComponentManager) GetInstance(name string) (interface{}, error) {
	return c.Container.Get(name)
}
