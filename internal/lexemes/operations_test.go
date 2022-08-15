package lexemes

import (
	"math"
	"testing"
)

func TestNothing(t *testing.T) {
	ref := 0.0
	res, err := Nothing(1.1, 2.2, 3.3)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestSum(t *testing.T) {
	ref := 4.0
	res, err := Sum(2.0, 2.0)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestSub(t *testing.T) {
	ref := 0.0
	res, err := Sub(2.0, 2.0)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestMult(t *testing.T) {
	ref := 4.0
	res, err := Mult(2.0, 2.0)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestDiv(t *testing.T) {
	ref := 1.0
	res, err := Div(2.0, 2.0)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestDiv0(t *testing.T) {
	res, err := Div(2.0, 0.0)
	if err == nil {
		t.Errorf("Shuld be div0 error. Gor res %f and no error", res)
	}
}

func TestPow(t *testing.T) {
	ref := 4.0
	res, err := Pow(2.0, 2.0)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestLog(t *testing.T) {
	arg0, arg1 := 2.0, 2.71828
	ref := math.Log(arg0) / math.Log(arg1)
	res, err := Log(arg0, arg1)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestMod(t *testing.T) {
	ref := 0.0
	res, err := Mod(2.0, 2.0)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestExp(t *testing.T) {
	arg := 2.0
	ref := math.Exp(arg)
	res, err := Exp(arg)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}
func TestSqrt(t *testing.T) {
	arg := 2.0
	ref := math.Sqrt(arg)
	res, err := Sqrt(arg)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}
