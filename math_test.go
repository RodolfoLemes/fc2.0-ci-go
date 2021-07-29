package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(10, 15);

	if total != 25 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, 25)
	}
}