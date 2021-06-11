package domain

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

const hashSeparator = "/"
const hashBytesLength = 24

// HashValues Genera identificador único en base 64 para valores entregados usando algoritmo SHA1 tomando los primeros 24 bytes
func HashValues(values ...string) string {
	hash := sha1.New()
	hashInBytes := hash.Sum([]byte(strings.Join(values, hashSeparator)))
	return base64.StdEncoding.EncodeToString(hashInBytes[:hashBytesLength])
}

// EncodeValuesToB64 Genera un código hexadecimal a partir de elementos de un arreglo
func EncodeValuesToB64(separator string, values ...string) string {
	base := strings.Join(values, separator)
	encoded := base64.StdEncoding.EncodeToString([]byte(base))
	return encoded
}

// Decode64ToValues Genera un arreglo de datos a partir de valores codificados como hexadecimal
func Decode64ToValues(separator string, encoded string) ([]string, error) {
	b, e := base64.StdEncoding.DecodeString(encoded)
	if e != nil {
		return nil, e
	}
	base := string(b)
	return strings.Split(base, separator), nil

}
