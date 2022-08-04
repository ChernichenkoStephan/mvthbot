package solving

/*
func NewMockConverter() func(line string) []string {
  cases := map[string][]string{
    "2+2": ("2.0","2.0","+"),
    "2-2": ("2.0","2.0","-"),
    "2*2": ("2.0","2.0","*"),
    "2/2": ("2.0","2.0","/"),
    "2^2": ("2.0","2.0","^"),
    "pow(2;2)": ("2","2","pow"),
    "log(2;2)": ("2","2","log"),
    "mod(2;2)": ("2","2","mod"),
    "exp(2)": ("2","exp"),
  }

  return func(line string) []string {
    if v, ok := cases[line]; ok {
      return v
    }
    return []string{}
  }
}
*/

/*
func TestSum(t *testing.T){
    ref := 4.0
    res, err := Solve([]string{"2.0","2.0","+"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestSub(t *testing.T){
    ref := 0.0
    res, err := Solve([]string{"2.0","2.0","-"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestMul(t *testing.T){
    ref := 4.0
    res, err := Solve([]string{"2.0","2.0","*"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestDiv(t *testing.T){
    ref := 1.0
    res, err := Solve([]string{"2.0","2.0","/"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestSpo(t *testing.T){
    ref := 4.0
    res, err := Solve([]string{"2.0","2.0","^"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestPov(t *testing.T){
    ref := 4.0
    res, err := Solve("p[]string{"2","2","pow"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestLog(t *testing.T){
    ref := 1.0
    res, err := Solve("l[]string{"2","2","log"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestMod(t *testing.T){
    ref := 0.0
    res, err := Solve("m[]string{"2","2","mod"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestExp(t *testing.T){
    ref := 7.38905609893065
    res, err := Solve([]string{"2","exp"})
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
*/
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
			t.Errorf("got %q, wanted %q", got, want)
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
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
*/
