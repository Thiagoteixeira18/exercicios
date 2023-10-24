package exercicios

import "testing"

func TestS(t *testing.T) {
	s := "ola"
	if s != "ola" {
	t.Errorf("esperado ola, recebido %s", s)
}
}