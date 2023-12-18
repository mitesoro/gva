package utils

import "testing"

func TestTime(t *testing.T) {
	t.Log(GetTime(0))
	t.Log(GetTime(20231218))
}
