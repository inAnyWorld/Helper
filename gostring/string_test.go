package gostring

import (
	"reflect"
	"strconv"
	"testing"
)

func TestByteToString(t *testing.T) {
	b := []byte("123abc")
	s := ByteToString(&b)
	got := *s
	want := `123abc`
	if !reflect.DeepEqual(got, want) {
		t.Errorf("The values of %v is not %v\n", got, want)
	}
}

func TestConvString(t *testing.T) {
	// int
	var i int = 8
	got8 := strconv.Itoa(i)
	want8 := ConvString(i)
	if !reflect.DeepEqual(got8, want8) {
		t.Errorf("The values of int %v is not %v\n", got8, want8)
	}

	// unit64
	var n uint64 = 8
	got64u := strconv.Itoa(int(n))
	want64u := ConvString(n)
	if !reflect.DeepEqual(got64u, want64u) {
		t.Errorf("The values of unit64 %v is not %v\n", got64u, want64u)
	}

	// int64
	var ts int64 = 8
	got64 := strconv.FormatInt(ts, 10)
	want64 := ConvString(ts)
	if !reflect.DeepEqual(got64, want64) {
		t.Errorf("The values of int64 %v is not %v\n", got64, want64)
	}

	// float
	var f float64 = 3.1415
	got64f := strconv.FormatFloat(f, 'f', -1, 64)
	want64f := ConvString(f)
	if !reflect.DeepEqual(got64f, want64f) {
		t.Errorf("The values of float64 %v is not %v\n", got64f, want64f)
	}

	// string
	var s string = `is string`
	wants := ConvString(s)
	if !reflect.DeepEqual(s, wants) {
		t.Errorf("The values of string %v is not %v\n", s, wants)
	}

	// other
	var o bool
	want0 := ConvString(o)
	if reflect.DeepEqual(o, want0) {
		t.Errorf("The values of bool %v is not %v\n", o, wants)
	}

	var ss string = ""
	wantSS := ConvString(ss)
	if !reflect.DeepEqual(ss, wantSS) {
		t.Errorf("The values of empty string %v is not %v\n", ss, wantSS)
	}
}

func TestBase64Encode(t *testing.T)  {
	s := 1234
	s1 := strconv.Itoa(s)
	gotInt := Base64Encode(s1)
	wantInt := `MTIzNA==`
	if gotInt != wantInt {
		t.Errorf("The int values of %v is not %v\n", gotInt, wantInt)
	}

	s2 := `1234abc`
	gotStr := Base64Encode(s2)
	wantStr := `MTIzNGFiYw==`
	if gotStr != wantStr {
		if gotStr != wantStr {
			t.Errorf("The str values of %v is not %v\n", gotStr, wantStr)
		}
	}
}
