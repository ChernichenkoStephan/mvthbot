package solving

func (s Statement) Copy() *Statement {
	vs := make([]string, len(s.Variables))
	copy(vs, s.Variables)
	return &Statement{
		Variables: vs,
		Equation:  s.Equation,
		Value:     s.Value,
	}
}
