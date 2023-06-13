package main

import "testing"

func TestIntroducao(t *testing.T) {
	resultado := exibeIntroducao()
	esperado := "Paula"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
