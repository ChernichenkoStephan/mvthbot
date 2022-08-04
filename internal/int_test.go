package tests

/*
// -------------- PlusTests --------------

func TestPlus(t *testing.T){
    ref := 4.0
    res := FixEquation("2+2")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestPlusSpace(t *testing.T){
    ref := 4.0
    res := FixEquation(" 2 + 2 ")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestPlusMany(t *testing.T){
    ref := 10.0
    res := FixEquation("2+2+2+2+2")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

// -------------- MultTests --------------

func TestMult(t *testing.T){
    ref := 4.0
    res := FixEquation("2*2")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestMultSpace(t *testing.T){
    ref := 4.0
    res := FixEquation(" 2 * 2 ")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestMultMany(t *testing.T){
    ref := 32.0
    res := FixEquation("2*2*2*2*2")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

// -------------- MultTests --------------

func TestDiv(t *testing.T){
    ref := 1.0
    res := FixEquation("2/2")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestDivSpace(t *testing.T){
    ref := 1.0
    res := FixEquation(" 2 / 2 ")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestDivMany(t *testing.T){
    ref := 2.0
    res := FixEquation("32/2/2/2/2")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

// -------------- BracketsTests --------------

func TestBrackets(t *testing.T){
    ref := 4.0
    res := FixEquation("(2+2)")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestBracketsSpace(t *testing.T){
    ref := 4.0
    res := FixEquation("( 2 + 2 )")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestBracketsMany(t *testing.T){
    ref := 10.0
    res := FixEquation("(((2+2)+2)+(2+2))")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

// -------------- BracketsTests --------------

func TestPlus(t *testing.T){
    ref := 4.0
    res := FixEquation("pow(2;2)")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}



func TestPlusSpace(t *testing.T){
    ref := 16.0
    res := FixEquation("pow((2+2);2)")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}



func TestPlusMany(t *testing.T){
    ref := 16.0
    res := FixEquation("pow(pow(2;2);2)")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

// -------------- MixedTests --------------

func TestCase0(t *testing.T){
    ref := 4.1
    res := FixEquation(".1 + pow(2;2)")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestCase1(t *testing.T){
    ref := 3.5
    res := FixEquation("3+4*2/(1-5)^2")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestCase2(t *testing.T){
    ref := 99.0
    res := FixEquation("3*(1+2*(4*(2+2)))")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestCase3(t *testing.T){
    ref := 12.25
    res := FixEquation("pow(3+4*2/(1-5)^2;2)")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestCase4(t *testing.T){
    ref := 16.0
    res := FixEquation("(1+2)*(3+4)-5")
    eq, err := ToRPN(res)
    if err != nil {
      t.Errorf("got error: %v", err)
    }
    res, err := Solve(eq)
    if err != nil {
      t.Errorf("Got error: %v", err)
    }
    if res == ref {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
*/
