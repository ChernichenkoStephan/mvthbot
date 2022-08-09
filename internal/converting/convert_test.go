package converting

import "testing"

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

func TestBasic(t *testing.T) {
	ref := []string{"2", "2", "+"}
	res, err := ToRPN("2+2")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestNormalDot(t *testing.T) {
	ref := []string{"36.6", "2", "+"}
	res, err := ToRPN("36.6+2")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestBeforeDot(t *testing.T) {
	ref := []string{".54", "2", "+"}
	res, err := ToRPN(".54+2")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestAfterDot(t *testing.T) {
	ref := []string{"221.", "2", "+"}
	res, err := ToRPN("221.+2")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestSingle(t *testing.T) {
	c := "2+2"
	t.Run(c, func(t *testing.T) {
		ref := []string{"2", "2", "+"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "2-2"
	t.Run(c, func(t *testing.T) {
		ref := []string{"2", "2", "-"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "2*2"
	t.Run(c, func(t *testing.T) {
		ref := []string{"2", "2", "*"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "2/2"
	t.Run(c, func(t *testing.T) {
		ref := []string{"2", "2", "/"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "2^2"
	t.Run(c, func(t *testing.T) {
		ref := []string{"2", "2", "^"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
}

func TestMultiple(t *testing.T) {
	ref := []string{"2", "2", "+", "2", "+", "2", "+"}
	res, err := ToRPN("2+2+2+2")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestPriority(t *testing.T) {
	ref := []string{"2", "2", "2", "/", "+", "2", "+"}
	res, err := ToRPN("2+2/2+2")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestBrackets(t *testing.T) {
	ref := []string{"3", "1", "2", "+", "*"}
	res, err := ToRPN("3*(1+2)")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestBracketsSingle(t *testing.T) {
	ref := []string{"3"}
	res, err := ToRPN("(3)")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestBracketsRecursion(t *testing.T) {
	ref := []string{"3", "1", "2", "4", "2", "2", "+", "*", "*", "+", "*"}
	res, err := ToRPN("3*(1+2*(4*(2+2)))")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestNestedBrackets(t *testing.T) {
	ref := []string{"3", "1", "2", "4", "2", "2", "+", "*", "*", "+", "*"}
	res, err := ToRPN("3*(1+2*(4*(2+2)))")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestTwoPlaceFunc(t *testing.T) {
	ref := []string{"2", "2", "pow"}
	res, err := ToRPN("pow(2;2)")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestSinglePlaceFunc(t *testing.T) {
	ref := []string{"2", "2", "+", "exp"}
	res, err := ToRPN("exp(2+2)")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestNestedFuncs(t *testing.T) {
	ref := []string{"2", "2", "pow", "2", "pow"}
	res, err := ToRPN("pow(pow(2;2);2)")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if !equals(res, ref) {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}

func TestMixed(t *testing.T) {
	c := "(1+2)*(3+4)-5"
	t.Run(c, func(t *testing.T) {
		ref := []string{"1", "2", "+", "3", "4", "+", "*", "5", "-"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "3+4*2/(1-5)^2"
	t.Run(c, func(t *testing.T) {
		ref := []string{"3", "4", "2", "*", "1", "5", "-", "2", "^", "/", "+"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "pow(3+4*2/(1-5)^2;2)"
	t.Run(c, func(t *testing.T) {
		ref := []string{"3", "4", "2", "*", "1", "5", "-", "2", "^", "/", "+", "2", "pow"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
}

func TestSingleVariables(t *testing.T) {
	c := "a+b"
	t.Run(c, func(t *testing.T) {
		ref := []string{"a", "b", "+"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "a-b"
	t.Run(c, func(t *testing.T) {
		ref := []string{"a", "b", "-"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "a*b"
	t.Run(c, func(t *testing.T) {
		ref := []string{"a", "b", "*"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "a/b"
	t.Run(c, func(t *testing.T) {
		ref := []string{"a", "b", "/"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "a^b"
	t.Run(c, func(t *testing.T) {
		ref := []string{"a", "b", "^"}
		res, err := ToRPN(c)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
}

/*
func TestLeftBracketsError(t *testing.T) {
	c := "2+2)"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > "%v" ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
  c := "2+2))))"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > "%v" ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
  c := "(2+2))"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > "%v" ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
}

func TestRightBracketsError(t *testing.T) {
	c := "(2+2"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > "%v" ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
  c := "(((2+2"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > "%v" ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
  c := "((2+2)"
	t.Run(c, func(t *testing.T) {
    ref := []string{}
    res, err := ToRPN(c)
		if err == nil {
      t.Errorf("got no error, but there is one here > "%v" ", c)
    }
    if _, ok := err.(XXX); !ok {
      t.Errorf("Wrong error type. %v", err)
    }
		if !equals(res, ref) {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
}

func TestTwoPlaceFunc(t *testing.T){
  ref := []string{}
  res, err := ToRPN("2+xyz(2;2)")
  if err == nil {
    t.Errorf("got no error, but there is one here > "%v" ", c)
  }
  if _, ok := err.(XXX); !ok {
    t.Errorf("Wrong error type. %v", err)
  }
  if !equals(res, ref) {
    t.Errorf("got %q, wanted %q", res, ref)
  }
}
*/
