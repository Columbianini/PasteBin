package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2025, 5, 25, 10, 56, 0, 0, time.UTC)
	hd := humanDate(tm)

	if hd != "25 May 2025 at 10:56" {
		t.Errorf("Expected '25 May 2025 at 10:56', got '%s'", hd)
	}
}
