package main

import "testing"

func TestMessage(t *testing.T) {
	msg := message("guys")
	if msg != "Hello guys" {
		t.Errorf("Message was incorrect, got: %s, want: %s.", msg, "Hello guys")
	}
}
func TestQuerying(t *testing.T) {
	queryResult := checkDb(2)
	if queryResult != "2 dua dua@mail.com" {
		t.Errorf("query incorrect, got: %s, want: %s", queryResult, "2 dua dua@mail.com")
	}
}
