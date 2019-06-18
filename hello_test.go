package main

import "testing"

func TestMessage(t *testing.T) {
	msg := message("guys")
	if msg != "Hello guys" {
		t.Errorf("Message was incorrect, got: %s, want: %s.", msg, "Hello guys")
	}
}

/**func TestQuerying(t *testing.T) {
	id, name, email := checkDb(2)
	if id != 2 && name != "dua" && email != "dua@mail.com" {
		t.Errorf("query incorrect, got: %d, %s, %s, want: %s", id, name, email, "2, dua, dua@mail.com")
	}
}**/
