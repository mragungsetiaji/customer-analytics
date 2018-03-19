package main

import "testing"

func TestHello(t * testing.T){

	assertCorrectMesage := func(t *testing.T, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("to a person", func(t *testing.T){
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMesage(t, got, want)
	})

	t.Run("empty string", func(t *testing .T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMesage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMesage(t, got, want)
	})

	t.Run("in French", func(t * testing.T){
		got := Hello("Lauren", french)
		want := "Bonjour, Lauren"
		assertCorrectMesage(t, got, want)
	})
}