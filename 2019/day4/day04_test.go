package main

import "testing"

func TestCheckPassword(t *testing.T) {
	if !checkPassword(112233) {
		t.Errorf("112233 should pass check")
	}
	if !checkPassword(111122) {
		t.Errorf("111122 should pass check")
	}
	if checkPassword(123444) {
		t.Errorf("123444 should fail check")
	}
	if checkPassword(111111) {
		t.Errorf("111111 should fail check")
	}
	if checkPassword(223450) {
		t.Errorf("223450 should fail check")
	}
	if checkPassword(123789) {
		t.Errorf("123789 should fail check")
	}
}
