package lib

import "testing"

func TestWorld(t *testing.T) {
	if World != "World" {
		t.Fatal("The universe is broken")
	}
}

func TestGetWorld(t *testing.T) {
	if GetWorld() != "World" {
		t.Fatal("The universe is broken")
	}
}
