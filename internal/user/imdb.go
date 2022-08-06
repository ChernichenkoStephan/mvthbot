package user

import (
	"fmt"

	"github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

var _default_storage *storage = &storage{
	data: map[int64]*User{},
}

type storage struct {
	data map[int64]*User
}

func (s *storage) addUser(u *User) {
	s.data[u.ID] = u.Copy()
}

func (s storage) getUsers() *[]User {
	res := make([]User, len(s.data))
	for _, u := range s.data {
		t := u.Copy()
		fmt.Println(t.String())
		res = append(res, *t)
	}
	return &res
}

func (s storage) getUser(id int64) (*User, bool) {
	if u, ok := s.data[id]; ok {

		return u.Copy(), ok
	}
	return &User{}, false

}

func (s *storage) updateUser(u *User) error {
	if _, ok := s.data[u.ID]; ok {
		s.data[u.ID] = u.Copy()
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *storage) deleteUser(id int64) {
	delete(s.data, id)
}

func (s storage) getVariable(varName string, userID int64) (float64, error) {
	if u, ok := s.data[userID]; ok {
		if v, ok := u.Variables[varName]; ok {
			return v, nil
		}
		return 0.0, fmt.Errorf("VariableNotFound")
	}
	return 0.0, fmt.Errorf("UserNotFound")
}

func (s *storage) addVariable(varName string, value float64, userID int64) error {
	if u, ok := s.data[userID]; ok {
		s.data[u.ID].Variables[varName] = value
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *storage) updateVariable(varName string, value float64, userID int64) error {
	if u, ok := s.data[userID]; ok {
		u.Variables[varName] = value
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *storage) removeVariable(varName string, userID int64) error {
	if u, ok := s.data[userID]; ok {
		delete(u.Variables, varName)
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *storage) clearUserVariables(userID int64) error {
	if u, ok := s.data[userID]; ok {
		u.Variables = make(map[string]float64)
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s storage) getHistory(userID int64) (*[]solving.Statement, error) {
	if u, ok := s.data[userID]; ok {
		resp := make([]solving.Statement, len(*u.History))
		copy(resp, *u.History)
		return &resp, nil
	}
	return &[]solving.Statement{}, fmt.Errorf("UserNotFound")
}

func (s *storage) addStatement(userID int64, st *solving.Statement) error {
	if u, ok := s.data[userID]; ok {
		*u.History = append(*u.History, *st.Copy())
		return nil
	}
	return fmt.Errorf("UserNotFound")
}

func (s *storage) deleteHistory(userID int64) error {
	if u, ok := s.data[userID]; ok {
		*u.History = make([]solving.Statement, 0)
		return nil
	}
	return fmt.Errorf("UserNotFound")
}


//
//
//    Realization of repo
//
//


type imdbUserRepository struct {
	mx sync.Mutex
	db *storage
}

func NewImdbUserRepository() *imdbUserRepository {
	db := _default_storage
	return &imdbUserRepository{
		db: db,
	}
}

func (repo *imdbUserRepository) GetUsers(ctx context.Context) (*[]User, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	resp := repo.db.getUsers()
	return resp, nil
}

func (repo *imdbUserRepository) GetUser(ctx context.Context, userID int64) (*User, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	u, ok := repo.db.getUser(userID)
	if !ok {
		return &User{}, fmt.Errorf("UserNotFound")
	}
	return u, nil
}

func (repo *imdbUserRepository) AddUser(ctx context.Context, user *User) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	repo.db.addUser(user)
	return nil
}

func (repo *imdbUserRepository) UpdateUser(ctx context.Context, user *User) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.updateUser(user)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *imdbUserRepository) DeleteUser(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	repo.db.deleteUser(userID)
	return nil
}

func (repo *imdbUserRepository) GetHistory(ctx context.Context, userID int64) (*[]solv.Statement, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	h, err := repo.db.getHistory(userID)
	if err != nil {
		return &[]solv.Statement{}, fmt.Errorf("UserNotFound")
	}
	return h, nil
}

func (repo *imdbUserRepository) AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.addStatement(userID, statement)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *imdbUserRepository) DeleteHistory(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.deleteHistory(userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

type imdbVariableRepository struct {
	mx sync.Mutex
	db *storage
}

func NewImdbVariableRepository() *imdbVariableRepository {
	db := _default_storage
	return &imdbVariableRepository{
		db: db,
	}
}

func (repo *imdbVariableRepository) CreateUserVariable(ctx context.Context, userID int64, name string, value float64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.addVariable(name, value, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *imdbVariableRepository) CreateUserVariables(
	ctx context.Context,
	userID int64,
	names []string,
	values []float64,
) error {

	repo.mx.Lock()
	defer repo.mx.Unlock()
	if len(names) != len(values) {
		return fmt.Errorf("Length of names and values doesn't match")
	}

	for i := 0; i < len(names); i++ {
		err := repo.db.addVariable(names[i], values[i], userID)
		if err != nil {
			return fmt.Errorf("MultyAdd error")
		}
	}

	return nil

}

func (repo *imdbVariableRepository) GetUserVariable(ctx context.Context, userID int64, name string) (float64, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	v, err := repo.db.getVariable(name, userID)
	if err != nil {
		return 0.0, fmt.Errorf("UserNotFound")
	}
	return v, nil

}

func (repo *imdbVariableRepository) GetUserVariables(
	ctx context.Context,
	userID int64,
	names []string,
) (VMap, error) {

	repo.mx.Lock()
	defer repo.mx.Unlock()
	vs := make(VMap)
	for _, n := range names {
		v, err := repo.db.getVariable(n, userID)
		if err != nil {
			return VMap{}, fmt.Errorf("Got error: %v", err)
		}
		vs[n] = v
	}
	return vs, nil
}

func (repo *imdbVariableRepository) GetAllUserVariables(ctx context.Context, userID int64) (VMap, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	vs := make(VMap)
	u, ok := repo.db.getUser(userID)
	if !ok {
		return VMap{}, fmt.Errorf("UserNotFound")
	}

	for n, v := range u.Variables {
		vs[n] = v
	}
	return vs, nil
}

func (repo *imdbVariableRepository) UpdateUserVariable(ctx context.Context, userID int64, name string, value float64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.updateVariable(name, value, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *imdbVariableRepository) UpdateUserVariables(
	ctx context.Context,
	userID int64,
	names []string,
	values []float64,
) error {

	repo.mx.Lock()
	defer repo.mx.Unlock()
	if len(names) != len(values) {
		return fmt.Errorf("Length of names and values doesn't match")
	}

	for i := 0; i < len(names); i++ {
		err := repo.db.updateVariable(names[i], values[i], userID)
		if err != nil {
			return fmt.Errorf("MultyAdd error")
		}
	}

	return nil

}

func (repo *imdbVariableRepository) DeleteUserVariable(ctx context.Context, userID int64, name string) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.removeVariable(name, userID)
	if err != nil {
		return fmt.Errorf("Got error: %v", err)
	}
	return nil
}

func (repo *imdbVariableRepository) DeleteUserVariables(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.clearUserVariables(userID)
	if err != nil {
		return fmt.Errorf("Got error: %v", err)
	}
	return nil
}
