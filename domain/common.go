package domain

// ID Representa el ID de un dominio. El ID debe ser generado a través de atributos de un dominio
type ID interface {
	Encode() (string, error)
	Decode(string) (ID, error)
}
