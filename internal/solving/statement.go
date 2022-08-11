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

func (this *Statement) Equals(other *Statement) bool {
	if this.Value != other.Value {
		return false
	}
	if this.Equation != other.Equation {
		return false
	}
	if len(this.Variables) != len(other.Variables) {
		return false
	}
	for i := 0; i < len(this.Variables); i++ {
		if this.Variables[i] != other.Variables[i] {
			return false
		}
	}
	return true
}
