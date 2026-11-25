package frete

import (
	"testing"

	"github.com/google/uuid"
)

func TestIdentificadorDoAlertaDeAprovacao(t *testing.T) {
	id, err := uuid.Parse("6f585bab-09d2-4e66-98de-341d94cbaa2e")
	if err != nil {
		t.Fatalf("identificador do alerta inválido: %v", err)
	}
	if id.Version() != 4 {
		t.Fatalf("versão do identificador = %d, esperada 4", id.Version())
	}
}
