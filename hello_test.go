package main

import "testing"

func TestMessage(t *testing.T) {
	msg := message("guys")
	if msg != "Hello guys" {
		t.Errorf("Message was incorrect, got: %s, want: %s.", msg, "Hello guys")
	}
}
