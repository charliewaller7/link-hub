package src

import (
	"main/data"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Happy Path Test
	path := "test_resources/happy.yaml"
	expected := data.Config{
		Apps: []data.App{
			{
				Name: "bbc",
				Url: "www.bbc.com",
			},
			{
				Name: "youtube",
				Url: "www.youtube.com",
			},
		},
	}

	actual, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("Failed to load config|%v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Actual does not match expected.\nGot: %v\nExpected: %v", actual, expected)
	}

	// Unhappy path test
	path = "test_resources/unhappy"
	if _, err := LoadConfig(path); err == nil {
		t.Fatalf("Expected error when unmarshalling %q, got nil|", path)
	}
}