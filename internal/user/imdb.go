package user

/*

func (s *IMStorage) addUser(u *User) {
	s.data[u.ID] = u.Copy()
}

func (s IMStorage) getUsers() *[]User {
	res := make([]User, len(s.data))
	for _, u := range s.data {
		t := u.Copy()
		fmt.Println(t.String())
		res = append(res, *t)
	}
	return &res
}

func (s IMStorage) getUser(id int64) (*User, bool) {
	if u, ok := s.data[id]; ok {

		return u.Copy(), ok
	}
	return &User{}, false
}

func (s IMStorage) GetAuth(id int64) (string, bool) {
	if u, ok := s.data[id]; ok {
		return u.Password, ok
	}
	return "", false
}

func (s *IMStorage) updateUser(u *User) error {
	if _, ok := s.data[u.ID]; ok {
		s.data[u.ID] = u.Copy()
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *IMStorage) deleteUser(id int64) {
	delete(s.data, id)
}

func (s IMStorage) getVariable(varName string, userID int64) (float64, error) {
	if u, ok := s.data[userID]; ok {
		if v, ok := u.Variables[varName]; ok {
			return v, nil
		}
		return 0.0, fmt.Errorf("VariableNotFound")
	}
	return 0.0, fmt.Errorf("UserNotFound")
}

func (s *IMStorage) addVariable(varName string, value float64, userID int64) error {
	if u, ok := s.data[userID]; ok {
		s.data[u.ID].Variables[varName] = value
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *IMStorage) updateVariable(varName string, value float64, userID int64) error {
	if u, ok := s.data[userID]; ok {
		u.Variables[varName] = value
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *IMStorage) removeVariable(varName string, userID int64) error {
	if u, ok := s.data[userID]; ok {
		delete(u.Variables, varName)
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *IMStorage) clearUserVariables(userID int64) error {
	if u, ok := s.data[userID]; ok {
		u.Variables = make(map[string]float64)
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s IMStorage) getHistory(userID int64) (*[]solving.Statement, error) {
	if u, ok := s.data[userID]; ok {
		resp := make([]solving.Statement, len(*u.History))
		copy(resp, *u.History)
		return &resp, nil
	}
	return &[]solving.Statement{}, fmt.Errorf("UserNotFound")
}

func (s *IMStorage) addStatement(userID int64, st *solving.Statement) error {
	if u, ok := s.data[userID]; ok {
		*u.History = append(*u.History, *st.Copy())
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *IMStorage) deleteHistory(userID int64) error {
	if u, ok := s.data[userID]; ok {
		*u.History = make([]solving.Statement, 0)
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

*/
