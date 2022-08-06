package user

import (
	"fmt"

	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

func (u User) Copy() *User {
	h := make([]solv.Statement, 1)
	vs := make(map[string]float64)
	for _, s := range *u.History {
		h = append(h, *s.Copy())
	}
	for k, v := range u.Variables {
		vs[k] = v
	}
	return &User{
		ID:        u.ID,
		History:   &h,
		Variables: vs,
	}
}

func (u User) String() string {
	return fmt.Sprintf("{\n\tID: %d,\n\tHistory: %v,\n\tVariables: %v\n}",
		u.ID, u.History, u.Variables)
}
