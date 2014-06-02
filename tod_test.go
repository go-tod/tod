package tod

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewTime(t *testing.T) {
	nt := NewTime()
	if nt.hours != 0 {
		t.Errorf("expected 0, got %d", nt.hours)
	}
}

func TestSetHours(t *testing.T) {
	nt := NewTime()
	err := nt.SetHours(24)
	if err != ErrInvalidTime {
		t.Errorf("expected ErrInvalidTime, got %v", err)
	}
	if nt.hours != 0 {
		t.Errorf("expected 0, got %d", nt.hours)
	}
	err = nt.SetHours(17)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if nt.hours != 17 {
		t.Errorf("expected 17, got %d", nt.Hours())
	}
}

func TestStr(t *testing.T) {
	nt := NewTime()
	nt.SetTime(20, 32)
	s := fmt.Sprintf("%s", nt)
	if s != "20:32" {
		t.Errorf("expected \"20:32\", got \"%s\"", s)
	}
}

func TestParseString(t *testing.T) {
	nt := NewTime()
	if err := nt.ParseString("X"); err != ErrInvalidString {
		t.Errorf("expected ErrInvalidString, got %v", err)
	}
	if err := nt.ParseString(" 23:12 "); err != ErrInvalidString {
		t.Errorf("expected ErrInvalidString, got %v", err)
	}

	if err := nt.ParseString("20:32"); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if nt.Hours() != 20 {
		t.Errorf("expected 20, got %d", nt.Hours())
	}
	if nt.Minutes() != 32 {
		t.Errorf("expected 32, got %d", nt.Minutes())
	}
}

func TestMashalJSON(t *testing.T) {
	nt := NewTime()
	nt.SetTime(12, 34)
	type jTime struct {
		T Time `json:"t"`
	}
	j := jTime{
		T: nt,
	}
	b, err := json.Marshal(j)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if string(b) != "{\"t\":\"12:34\"}" {
		t.Errorf("expected \"12:34\", got %s", b)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	nt := NewTime()
	j := []byte("{\"T\" : \"12:34\"}")
	type jsontime struct {
		T Time `json:"t"`
	}
	jsonTime := jsontime{}
	err := json.Unmarshal(j, &jsonTime)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if jsonTime.T.Hours() != 12 {
		t.Errorf("expected 12, got %d", nt.Hours())
	}
}
