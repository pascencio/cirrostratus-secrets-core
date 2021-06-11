package application

import (
	"fmt"
	"sync"

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
	Register(string, ComponentFactory) error
	GetComponent(string) (interface{}, error)
}

var once sync.Once
var instance *defaultComponentManager

type defaultComponentManager struct {
	container ComponentContainer
}

// Register Registra un componente en el contenedor
func (c *defaultComponentManager) Register(name string, factory ComponentFactory) error {
	logrus.WithFields(
		logrus.Fields{
			"componentName": name,
		}).Info("Registrando componente")
	if e := c.container.add(name, factory); e != nil {
		return e
	}
	logrus.WithFields(
		logrus.Fields{
			"componentName": name,
		},
	).Info("Componente registrado")
	return nil
}

// GetComponent Retorna un instancia del componente registrado
func (c *defaultComponentManager) GetComponent(name string) (interface{}, error) {
	return c.container.Get(name)
}

func GetInstance() *defaultComponentManager {
	once.Do(func() {
		instance = &defaultComponentManager{
			container: ComponentContainer{
				Factories: map[string]ComponentFactory{},
				Cache:     make(map[string]interface{}),
			},
		}
	})
	return instance
}
