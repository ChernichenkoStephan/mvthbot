package fixing

import "testing"

func TestDoupleOp(t *testing.T) {
	c := "2+-2"
	t.Run(c, func(t *testing.T) {
    f := New()
		res := f.Fix(c)
		ref := "2-2"
		if res != ref {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "2-+2"
	t.Run(c, func(t *testing.T) {
    f := New()
		res := f.Fix(c)
		ref := "2-2"
		if res != ref {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "2--2"
	t.Run(c, func(t *testing.T) {
    f := New()
		res := f.Fix(c)
		ref := "2+2"
		if res != ref {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "2++2"
	t.Run(c, func(t *testing.T) {
    f := New()
		res := f.Fix(c)
		ref := "2+2"
		if res != ref {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
}

func TestBracket(t *testing.T) {
	c := "(+2+2)"
	t.Run(c, func(t *testing.T) {
    f := New()
		res := f.Fix(c)
		ref := "(2+2)"
		if res != ref {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})

	c = "(-2)+2"
	t.Run(c, func(t *testing.T) {
    f := New()
		res := f.Fix(c)
		ref := "(0-2)+2"
		if res != ref {
			t.Errorf("got %q, wanted %q", res, ref)
		}
	})
}

func TestSpaces(t *testing.T) {
  f := New()
	res := f.Fix(" ( 2 + 2 ) ")
	ref := "(2+2)"
	if res != ref {
		t.Errorf("got %q, wanted %q", res, ref)
	}
}
