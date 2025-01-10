package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Johan", "")
		want := "Hello, Johan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jim", "French")
		want := "Bonjour, Jim"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this, when it fails, the line number reported will be in our function
	// call rather than inside our test helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
