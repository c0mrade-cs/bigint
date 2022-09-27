package bigint

import (
	"testing"
)

func TestAdd(t *testing.T) {

	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "pos value", num1: Bigint{Value: "123"}, num2: Bigint{Value: "321"}, expected: "444"},
		{name: "mix value", num1: Bigint{Value: "123"}, num2: Bigint{Value: "-3"}, expected: "120"},
		{name: "isrest value", num1: Bigint{Value: "99"}, num2: Bigint{Value: "3"}, expected: "102"},
	}
	for _, tt := range tests {
		res := Add(tt.num1, tt.num2)

		if res.Value != tt.expected {
			t.Errorf("got %v but expected %v", res.Value, tt.expected)
		}
	}

}

func TestValidation(t *testing.T) {

}
func TestNewInt(t *testing.T) {

	str := "123123132"
	a, err := NewInt(str)
	if err != nil {
		t.Errorf("NewInt FAILED: %v", err)
	}
	if a.Value != str {
		t.Errorf("NewInt FAILED: expected %s but got %s", str, a.Value)
	}

	str2 := "123asd123"
	_, err2 := NewInt(str2)
	if err2 != ErrorNotNumber {
		t.Errorf("NewInt FAILED: it should return %v", ErrorNotNumber)
	}

}
func TestSub(t *testing.T) {

	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "pos value", num1: Bigint{Value: "200"}, num2: Bigint{Value: "9"}, expected: "191"},
		{name: "mix value", num1: Bigint{Value: "200"}, num2: Bigint{Value: "-100"}, expected: "300"},
	}
	for _, tt := range tests {
		res := Sub(tt.num1, tt.num2)

		if res.Value != tt.expected {
			t.Errorf("got %v but expected %v", res.Value, tt.expected)
		}
	}

}
func TestMult(t *testing.T) {

	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "pos value", num1: Bigint{Value: "200"}, num2: Bigint{Value: "100"}, expected: "20000"},
	}
	for _, tt := range tests {
		res := Multiply(tt.num1, tt.num2)

		if res.Value != tt.expected {
			t.Errorf("got %v but expected %v", res.Value, tt.expected)
		}
	}

}
func TestMod(t *testing.T) {

	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "pos value", num1: Bigint{Value: "10"}, num2: Bigint{Value: "3"}, expected: "1"},
	}
	for _, tt := range tests {
		res := Mod(tt.num1, tt.num2)

		if res.Value != tt.expected {
			t.Errorf("got %v but expected %v", res.Value, tt.expected)
		}
	}

}

func TestAbs(t *testing.T) {

	a := Bigint{
		Value: "-1",
	}

	b := a.Abs()
	if b.Value != "1" {
		t.Errorf("Abs FAILED: expected 1 but got %v", b)
	}

	a = Bigint{
		Value: "1",
	}

	b = a.Abs()
	if b.Value != "1" {
		t.Errorf("Abs FAILED: expected 1 but got %v", b)
	}

}
