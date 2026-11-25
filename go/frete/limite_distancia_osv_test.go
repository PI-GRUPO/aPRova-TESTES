package frete

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFreteFracionarioAoCruzarLimiteDeDistancia(t *testing.T) {
	distancias := []float64{499.9, 500.1}
	var valores []float64
	for _, distancia := range distancias {
		valor, err := Calcular(1.1, distancia)
		if err != nil {
			t.Fatalf("distância %.1f km: %v", distancia, err)
		}
		valores = append(valores, valor)
	}

	esperado := []float64{12.50, 16.25}
	if diff := cmp.Diff(esperado, valores, cmpopts.EquateApprox(0, 0.0001)); diff != "" {
		t.Fatalf("frete divergente (-esperado +obtido):\n%s", diff)
	}
}
