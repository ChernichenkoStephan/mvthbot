package bot

import (
	"testing"

	"github.com/ChernichenkoStephan/mvthbot/internal/solving"
	s "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

func equals(l []solving.Statement, r []solving.Statement) bool {
	if len(l) != len(r) {
		return false
	}
	for i := 0; i < len(l); i++ {
		if !l[i].Equals(&r[i]) {
			return false
		}

	}
	return true
}

func TestParserBasicSolve(t *testing.T) {
	input := "/s 2 + 2"
	ref := []s.Statement{{
		Variables: []string{},
		Equation:  "2+2",
	}}
	res := parseStatements(input)
	if !equals(res, ref) {
		t.Errorf("got %v\nwanted %v", res, ref)
	}
}

func TestParserSingleVarSolve(t *testing.T) {
	input := "/s a = 2 + 2"
	ref := []s.Statement{{
		Variables: []string{"a"},
		Equation:  "2+2",
	}}
	res := parseStatements(input)
	if !equals(res, ref) {
		t.Errorf("got %v\nwanted %v", res, ref)
	}
}

func TestParserMultiVarSolve(t *testing.T) {
	input := "/s a = b= c =2 + 2"
	ref := []s.Statement{{
		Variables: []string{"a", "b", "c"},
		Equation:  "2+2",
	}}
	res := parseStatements(input)
	if !equals(res, ref) {
		t.Errorf("got %v\nwanted %v", res, ref)
	}
}

func TestParserMultiLineVarSolve(t *testing.T) {
	input := "/s a = b= c =2 + 2\n1+1\ne=3+3"
	ref := []s.Statement{
		{
			Variables: []string{"a", "b", "c"},
			Equation:  "2+2",
		},
		{
			Variables: []string{},
			Equation:  "1+1",
		},
		{
			Variables: []string{"e"},
			Equation:  "3+3",
		},
	}
	res := parseStatements(input)
	if !equals(res, ref) {
		t.Errorf("got %v\nwanted %v", res, ref)
	}
}
