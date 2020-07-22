package main

import (
	"testing"
)

func TestOla(t *testing.T) {

	checkCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper() //indica que é um método auxiliar (necessário)
		if result != expected {
			t.Errorf("resultado '%s', esperado '%s'", result, expected)
		}
	}

	t.Run(" diz olá para as pessoas", func(t *testing.T) {
		result := Hello("Chris", "português")
		expected := "Olá, Chris"

		checkCorrectMessage(t, result, expected)
	})

	t.Run(" diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		result := Hello("", "")
		expected := "Olá, mundo"

		checkCorrectMessage(t, result, expected)
	})

	t.Run(" em espanhol", func(t *testing.T) {
		result := Hello("Elodie", "espanhol")
		expected := "Hola, Elodie"

		checkCorrectMessage(t, result, expected)
	})

	t.Run(" em francês", func(t *testing.T) {
		result := Hello("Maria", "francês")
		expected := "Bonjour, Maria"

		checkCorrectMessage(t, result, expected)
	})

	t.Run(" em italiano", func(t *testing.T) {
		result := Hello("Pepe", "italiano")
		expected := "Ciao, Pepe"

		checkCorrectMessage(t, result, expected)
	})
}
