package domain

import (
	"fmt"
	"strings"
)

// Action Acción realizada sobre algún componente de la capa de lógica de negocio
type Action struct {
	Context string
	Action  string
}

func (a Action) String() string {
	action := fmt.Sprintf("%s:%s", a.Context, a.Action)
	return strings.ToLower(action)
}

type Permission struct {
	Context string
	Action  string
}

func (p Permission) ToString() string {
	permission := fmt.Sprintf("%s:%s", p.Context, p.Action)
	return strings.ToLower(permission)
}

func (p Permission) IsPermited(action Action) bool {
	return p.ToString() == action.String()
}

func HasPermission(action Action, permission ...Permission) bool {
	for _, p := range permission {
		if p.IsPermited(action) {
			return true
		}
	}
	return false
}
