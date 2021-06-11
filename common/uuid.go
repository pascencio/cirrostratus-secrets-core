package common

import (
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
