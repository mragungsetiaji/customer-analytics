package main

import "testing"

func TestHello(t * testing.T){
	got := Hello("Chris")
	want := "Helo, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}