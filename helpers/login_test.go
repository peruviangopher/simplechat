package helpers

import (
	"testing"
)

// Tests for CheckUserPass
func TestCheckUserPass(t *testing.T) {
	tests := []struct {
		username string
		password string
		expected bool
	}{
		{"Bob", "1111", true},      // valid user/pass
		{"Jimmy", "2222", true},    // valid user/pass
		{"Paul", "3333", true},     // valid user/pass
		{"Kat", "4444", true},      // valid user/pass
		{"Bob", "wrong", false},    // wrong password
		{"Unknown", "1234", false}, // user does not exist
		{"", "", false},            // empty credentials
	}

	for _, tt := range tests {
		result := CheckUserPass(tt.username, tt.password)
		if result != tt.expected {
			t.Errorf("CheckUserPass(%q, %q) = %v; want %v", tt.username, tt.password, result, tt.expected)
		}
	}
}

// Tests for EmptyUserPass
func TestEmptyUserPass(t *testing.T) {
	tests := []struct {
		username string
		password string
		expected bool
	}{
		{"", "", true},         // both empty
		{"Bob", "", true},      // empty password
		{"", "1111", true},     // empty username
		{"   ", "2222", true},  // username only spaces
		{"Jimmy", "   ", true}, // password only spaces
		{"Bob", "1111", false}, // valid username and password
	}

	for _, tt := range tests {
		result := EmptyUserPass(tt.username, tt.password)
		if result != tt.expected {
			t.Errorf("EmptyUserPass(%q, %q) = %v; want %v", tt.username, tt.password, result, tt.expected)
		}
	}
}
