package gostring

import (
	"reflect"
	"strconv"
	"testing"
)

type Example struct {
	Examples string
}

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
	i := 8
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
	f := 3.1415
	got64f := strconv.FormatFloat(f, 'f', -1, 64)
	want64f := ConvString(f)
	if !reflect.DeepEqual(got64f, want64f) {
		t.Errorf("The values of float64 %v is not %v\n", got64f, want64f)
	}

	// string
	s := `is string`
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

	ss := ""
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

func TestStructName(t *testing.T) {
	e := `Example`
	structName := StructName(&Example{})
	if e != structName {
		t.Errorf("The struct name values of %v is not %v\n", structName, e)
	}
}

func TestCalculatePercentage(t *testing.T) {
	e1 := `50%`
	p1 := CalculatePercentage(1, 2)
	if e1 != p1 {
		t.Errorf("The percentage values of %v is not %v\n", e1, p1)
	}

	e2 := `33.33%`
	p2 := CalculatePercentage(1, 3)

	if e2 != p2 {
		t.Errorf("The percentage values of %v is not %v\n", e2, p2)
	}

	e3 := `0.00%`
	p3 := CalculatePercentage(1, 0)
	if e3 != p3 {
		t.Errorf("The percentage values of %v is not %v\n", e3, p3)
	}
}

func TestRandStringRunes(t *testing.T) {
	n1 := 1
	s1 := RandStringRunes(n1)
	if len(s1) != n1 {
		t.Errorf("The round len values of %v is not %v\n", n1, len(s1))
	}
}

func TestRandomString(t *testing.T) {
	var n1 uint8 = 4
	s1 := RandomString(n1, 1)

	if len(s1) != 4 {
		t.Errorf("The round type values of 1 len of %v is not %v\n", n1, len(s1))
	}

	s2 := RandomString(4, 2)
	if len(s2) != 4 {
		t.Errorf("The round type values of values2 len of %v is not %v\n", n1, len(s2))
	}

	s3 := RandomString(4, 3)
	if len(s3) != 4 {
		t.Errorf("The round type values of 3 len of %v is not %v\n", n1, len(s3))
	}

	s4 := RandomString(4, 4)
	if len(s3) != 4 {
		t.Errorf("The round type values of 4 len of %v is not %v\n", n1, len(s4))
	}

	s5 := RandomString(4, 5)
	if len(s3) != 4 {
		t.Errorf("The round type values of 5 len of %v is not %v\n", n1, len(s5))
	}
}

func TestGetBetweenStr(t *testing.T) {
	s := "@abc456"
	e := GetBetweenStr(`@abc456%`, "@", "%")
	if s != e {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e, s)
	}

	e1 := GetBetweenStr(`vcx@abc456%aio`, "@", "%")
	if s != e1 {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e1, s)
	}

	e2 := GetBetweenStr(`@abc456%aio`, "@", "%")
	if s != e1 {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e2, s)
	}

	e3 := GetBetweenStr(`aaaa@abc456%`, "@", "%")
	if s != e3 {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e3, s)
	}

	e4 := GetBetweenStr(`@abc456%`, "@", "%")
	if s != e3 {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e4, s)
	}
}

func TestSubstr(t *testing.T) {
	s := `a`
	e := Substr(`abc@123&`, 0, 1)
	if s != e {
		t.Errorf("The Substr values of %v is not %v\n", e, s)
	}

	s1 := `3&`
	e1 := Substr(`abc@123&`, -1, 2)
	if s1 != e1 {
		t.Errorf("The Substr values of %v is not %v\n", e1, s1)
	}

	s2 := `abc@123&`
	e2 := Substr(`abc@123&`, 0, 9999)
	if s2 != e2 {
		t.Errorf("The Substr values of %v is not %v\n", e2, s2)
	}
}