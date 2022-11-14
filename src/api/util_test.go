package api

import (
	"testing"
)

func TestConvertDatetime(t *testing.T) {
	s := "2022-10-31T01:42:35.677938797Z"
	got, _ := ConvertDatetime(s)
	want := "2022-10-31 01:42:35"
	if got.Format("2006-01-02 15:04:05") != want {
		t.Errorf("Error Got: %v Want: %v", got, want)
	}
}
func TestSpliceSlice(t *testing.T) {
	s := []string{"hello", "golang"}
	want := "[hello, golang]"
	got := SpliceSlice(s, ",")
	if got != want {
		t.Errorf("Error Got: %v Want: %v", got, want)
	}
}
