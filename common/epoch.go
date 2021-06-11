package common

import (
	"strconv"
	"time"
)

const epochToMillis = 1000000
const epochToSeconds = 1000

// InNanos Conversión de texto en formato epoch en nanosegundos
const InNanos = 1

// InMillis Conversión de texto en formato epoch en milisegundos
const InMillis = epochToMillis

// InSeconds Conversión de texto en formato epoch en segundos
const InSeconds = InMillis * epochToSeconds

type epoch struct {
	nanos int64
}

func (e epoch) AsMillis() int64 {
	return e.nanos / epochToMillis
}

func (e epoch) AsSeconds() int64 {
	return e.AsMillis() / epochToSeconds
}

func (e epoch) AsMillisString() string {
	return strconv.FormatInt(e.AsMillis(), 10)
}

func (e epoch) AsSecondsString() string {
	return strconv.FormatInt(e.AsSeconds(), 10)
}

func (e epoch) AsNanosString() string {
	return strconv.FormatInt(e.nanos, 10)
}

// Epoch Crea una instancia de Unix Time en nano segundos
func Epoch() epoch {
	return epoch{
		nanos: time.Now().UnixNano(),
	}
}

// EpochFrom Crea una instancia de Unixt Time a partir de segundos, milisegundos o nanosegundos
func EpochFrom(value int64, from int64) epoch {
	return epoch{
		nanos: (value * from),
	}
}

// ParseEpoch Crea una instancia de Unix Time a partir de texto
func ParseEpoch(value string, in int64) (epoch, error) {
	i, e := strconv.ParseInt(value, 10, 64)
	if e != nil {
		return epoch{}, e
	}
	return epoch{
		nanos: (i * in),
	}, nil
}
