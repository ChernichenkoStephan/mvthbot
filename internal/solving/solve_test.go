package solving

import "testing"

func TestSum(t *testing.T) {
	ref := 4.0
	res, err := Solve([]string{"2", "2", "+"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestSub(t *testing.T) {
	ref := 0.0
	res, err := Solve([]string{"2", "2", "-"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestMul(t *testing.T) {
	ref := 4.0
	res, err := Solve([]string{"2", "2", "*"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestDiv(t *testing.T) {
	ref := 1.0
	res, err := Solve([]string{"2", "2", "/"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestSpo(t *testing.T) {
	ref := 4.0
	res, err := Solve([]string{"2", "2", "^"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestPov(t *testing.T) {
	ref := 4.0
	res, err := Solve([]string{"2", "2", "pow"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestLog(t *testing.T) {
	ref := 1.0
	res, err := Solve([]string{"2", "2", "log"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestMod(t *testing.T) {
	ref := 0.0
	res, err := Solve([]string{"2", "2", "mod"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestExp(t *testing.T) {
	ref := 7.38905609893065
	res, err := Solve([]string{"2", "exp"}, map[string]float64{})
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

func TestVariablesSum(t *testing.T) {
	ref := 4.2
	vars := map[string]float64{"a": 2.2}
	res, err := Solve([]string{"2", "a", "+"}, vars)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if res != ref {
		t.Errorf("got %f, wanted %f", res, ref)
	}
}

/*
func TestDiv0(t *testing.T) {
	c := "1/0"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := Solve(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > '%v' ", c)
    }
		if !equals(res, ref) {
			t.Errorf("got %f, wanted %f", res, ref)
		}
	})
  c := "(1/(1-1))"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := Solve(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > '%v' ", c)
    }
		if !equals(res, ref) {
			t.Errorf("got %f, wanted %f", res, ref)
		}
	})
}
*/
