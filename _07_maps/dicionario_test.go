package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	word := "test"
	definition := "isso é apenas um teste"
	dictionary := Dictionary{word: definition}

	t.Run("palavra existente", func(t *testing.T) {
		got, _ := dictionary.Search(word)
		want := definition

		compareStrings(t, got, want)
	})

	t.Run("palavra nova", func(t *testing.T) {
		word = "palavra nova"
		_, err := dictionary.Search(word)
		want := ErrorNotFound

		confirmError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "palavra"

	t.Run("palavra nova", func(t *testing.T) {
		want := "isso é apenas um teste"
		dictionary.Add(word, want)

		got, err := dictionary.Search(word)
		checkErrorUnspected(t, err)
		compareStrings(t, got, want)
	})

	t.Run("palavra existente", func(t *testing.T) {
		want := ErrorAlreadyExists
		err := dictionary.Add(word, "definição")
		confirmError(t, err, want)
	})
}

func TestUpdate(t *testing.T) {
	word := "palavra"
	definition := "definição"
	dictionary := Dictionary{word: definition}
	newDefinition := "nova definição"

	t.Run("palavra existente", func(t *testing.T) {
		err := dictionary.Update(word, newDefinition)
		checkErrorUnspected(t, err)
		compareDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("palavra nova", func(t *testing.T) {
		word = "palavra nova"
		want := ErrorNotExists
		err := dictionary.Update(word, "definição")
		confirmError(t, err, want)
	})
}

func TestDelete(t *testing.T) {
	word := "palavra"
	definition := "definição"
	dictionary := Dictionary{word: definition}

	t.Run("palavra existente", func(t *testing.T) {
		err := dictionary.Delete(word)
		checkErrorUnspected(t, err)
	})

	t.Run("palavra nova", func(t *testing.T) {
		word = "palavra nova"
		want := ErrorNotFound
		err := dictionary.Delete(word)
		confirmError(t, err, want)
	})
}

func checkErrorUnspected(t *testing.T, err error) {
	if err != nil {
		t.Errorf("erro inesperado")
	}
}

func confirmError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("é esperado um erro '%s'", got)
	}
}

func compareStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf(" obteve '%s', queria '%s'", got, want)
	}
}

func compareDefinition(t *testing.T, dictionary Dictionary, word string, want string) {
	t.Helper()

	got, _ := dictionary.Search(word)
	if got != want {
		t.Errorf(" obteve '%s', queria '%s'", got, want)
	}
}
