package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a basic test"}
	t.Run("known word", func(t *testing.T) {
		searchKey := "test"
		want := "this is just a basic test"

		got, _ := dictionary.Search(searchKey)

		assertStrings(t, got, want, searchKey)
	})
	t.Run("unknown word", func(t *testing.T) {
		searchKey := "unknown"

		_, err := dictionary.Search(searchKey)

		if err == nil {
			t.Fatal("Expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Basic add", func(t *testing.T) {
		searchKey := "test"
		searchDefinition := "this is a test"

		dictionary := Dictionary{}
		err := dictionary.Add(searchKey, searchDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, searchKey, searchDefinition)
	})
	t.Run("existing word", func(t *testing.T) {
		searchKey := "test"
		searchDefinition := "this is a test"

		dictionary := Dictionary{searchKey: searchDefinition}
		err := dictionary.Add(searchKey, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, searchKey, searchDefinition)

	})
	// t.Run("", func(t *testing.T) {})
}

func TestUpdate(t *testing.T) {
	t.Run("Basic update", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("basic delete", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)

		assertError(t, err, nil)

		_, err = dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})
	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition, word)
}
