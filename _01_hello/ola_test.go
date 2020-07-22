package hello

import (
	"testing"
)

func TestOla(t *testing.T) {

	checkCorrectMessage := func(t *testing.T, expected, result string) {
		t.Helper() //indica que é um método auxiliar (necessário)
		if result != expected {
			t.Errorf("esperado '%s', resultado '%s'", expected, result)
		}
	}

	t.Run(" diz olá para as pessoas", func(t *testing.T) {
		result := Hello("Chris", "português")
		expected := "Olá, Chris"

		checkCorrectMessage(t, expected, result)
	})

	t.Run(" diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		result := Hello("", "")
		expected := "Olá, mundo"

		checkCorrectMessage(t, expected, result)
	})

	t.Run(" em espanhol", func(t *testing.T) {
		result := Hello("Elodie", "espanhol")
		expected := "Hola, Elodie"

		checkCorrectMessage(t, expected, result)
	})

	t.Run(" em francês", func(t *testing.T) {
		result := Hello("Maria", "francês")
		expected := "Bonjour, Maria"

		checkCorrectMessage(t, expected, result)
	})

	t.Run(" em italiano", func(t *testing.T) {
		result := Hello("Pepe", "italiano")
		expected := "Ciao, Pepe"

		checkCorrectMessage(t, expected, result)
	})
}
