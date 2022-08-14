package user

import (
	"fmt"
)

func (u User) Copy() *User {
	h := make(History, 0)
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

func NewUser(id int64) *User {
	h := make(History, 0)
	vs := make(map[string]float64)

	return &User{
		ID:        id,
		History:   &h,
		Variables: vs,
	}
}

func (u User) Recipient() string {
	return fmt.Sprintf("%v", u.ID)
}
