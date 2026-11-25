package frete

import (
	"testing"

	"github.com/google/uuid"
)

func TestIdentificadorDeCenarioOSV(t *testing.T) {
	const identificador = "c41d7e9a-2f3b-4c8d-a5e6-0b1c2d3e4f5a"

	id, err := uuid.Parse(identificador)
	if err != nil {
		t.Fatalf("identificador do cenário inválido: %v", err)
	}
	if id.String() != identificador {
		t.Fatalf("identificador esperado %q, obtido %q", identificador, id)
	}
}
