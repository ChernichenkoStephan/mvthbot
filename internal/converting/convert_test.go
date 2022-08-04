package converting

/*
func equals(l []string, r []string) bool {
	if len(l) != len(r) {
		return false
	}
	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}

func TestBasic(t *testing.T){
    ref := []string{"2.0","2.0","+"}
    res, err := ToRPN("2+2")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestBasic(t *testing.T){
    ref := []string{"36.6","2.0","+"}
    res, err := ToRPN("36.6+2")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestBasic(t *testing.T){
    ref := []string{"0.54","2.0","+"}
    res, err := ToRPN(".54+2")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestBasic(t *testing.T){
    ref := []string{"221.0","2.0","+"}
    res, err := ToRPN("221.+2")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestSingle(t *testing.T) {
	c := "2+2"
	t.Run(c, func(t *testing.T) {
      ref := []string{"2.0","2.0","+"}
    res, err := ToRPN(c)
		if err != nil {
      t.Errorf("got error: %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	c := "2-2"
	t.Run(c, func(t *testing.T) {
      ref := []string{"2.0","2.0","-"}
    res, err := ToRPN(c)
		if err != nil {
      t.Errorf("got error: %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	c := "2*2"
	t.Run(c, func(t *testing.T) {
      ref := []string{"2.0","2.0","*"}
    res, err := ToRPN(c)
		if err != nil {
      t.Errorf("got error: %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	c := "2/2"
	t.Run(c, func(t *testing.T) {
      ref := []string{"2.0","2.0","/"}
    res, err := ToRPN(c)
		if err != nil {
      t.Errorf("got error: %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

  c := "2^2"
	t.Run(c, func(t *testing.T) {
      ref := []string{"2.0","2.0","^"}
    res, err := ToRPN(c)
		if err != nil {
      t.Errorf("got error: %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestMultiple(t *testing.T){
    ref := []string{"2.0","2.0","+","2.0","+","2.0","+"}
    res, err := ToRPN("2+2+2+2")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestPriority(t *testing.T){
    ref := []string{"2.0", "2.0", "2.0", "/", "+", "2.0", "+"}
    res, err := ToRPN("2+2/2+2")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestBrackets(t *testing.T){
    ref := []string{"3.0","1.0","2.0","+","*"}
    res, err := ToRPN("3*(1+2)")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestBrackets(t *testing.T){
    ref := []string{"3.0","1.0","2.0","+","*"}
    res, err := ToRPN("3*(1+2)")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestBracketsSingle(t *testing.T){
    ref := []string{"3.0",}
    res, err := ToRPN("(3)")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestNestedBrackets(t *testing.T){
    ref := []string{"3.0","1.0","2.0","4.0","2.0","2.0","+","*","*","+","*"}
    res, err := ToRPN("3*(1+2*(4*(2+2)))")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestTwoPlaceFunc(t *testing.T){
    ref := []string{"2.0","2.0","pow"}
    res, err := ToRPN("pow(2;2)")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


func TestSinglePlaceFunc(t *testing.T){
    ref := []string{"2.0","2.0","+","exp"}
    res, err := ToRPN("exp(2+2)")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestNestedFuncs(t *testing.T){
    ref := []string{"2.0","2.0","pow","2.0","pow"}
    res, err := ToRPN("pow(pow(2;2);2)")
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestMixed(t *testing.T) {
	c := "(1+2)*(3+4)-5"
	t.Run(c, func(t *testing.T) {
      ref := []string{"1.0","2.0","+","3.0","4.0","+","*","5.0","-"}
    res, err := ToRPN(c)
		if err != nil {
      t.Errorf("got error: %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

  c := "3+4*2/(1-5)^2"
  t.Run(c, func(t *testing.T) {
      ref := []string{"3.0","4.0","2.0","*","1.0","5.0","-","2.0","^","/","+"}
    res, err := ToRPN(c)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
      t.Errorf("got %q, wanted %q", got, want)
    }
  })

  c := "pow(3+4*2/(1-5)^2;2)"
  t.Run(c, func(t *testing.T) {
      ref := []string{"3.0","4.0","2.0","*","1.0","5.0","-","2.0","^","/","+","2.0","pow"}
    res, err := ToRPN(c)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    if !equals(res, ref) {
      t.Errorf("got %q, wanted %q", got, want)
    }
  })
}

func TestLeftBracketsError(t *testing.T) {
	c := "2+2)"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > '%v' ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
  c := "2+2))))"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > '%v' ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
  c := "(2+2))"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > '%v' ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestRightBracketsError(t *testing.T) {
	c := "(2+2"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > '%v' ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
  c := "(((2+2"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > '%v' ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
  c := "((2+2)"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > '%v' ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestTwoPlaceFunc(t *testing.T){
  ref := []string{}
  res, err := ToRPN("2+xyz(2;2)")
  if err == nil {
    t.Errorf("got no error, but there is one here > '%v' ", c)
  }
  if _, ok := err.(XXX); !ok {
    t.Errorf("Wrong error type. %v", err)
  }
  if !equals(res, ref) {
    t.Errorf("got %q, wanted %q", got, want)
  }
}
*/
