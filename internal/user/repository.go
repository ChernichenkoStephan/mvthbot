package user

import (
	"context"
	"fmt"
	"sync"

	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

type IMUserRepository struct {
	mx sync.Mutex
	db *IMStorage
}

func NewImdbUserRepository() *IMUserRepository {
	db := GetDefaultStorage()
	return &IMUserRepository{
		db: db,
	}
}

func (repo *IMUserRepository) GetUsers(ctx context.Context) (*[]User, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	resp := repo.db.getUsers()
	return resp, nil
}

func (repo *IMUserRepository) GetUser(ctx context.Context, userID int64) (*User, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	u, ok := repo.db.getUser(userID)
	if !ok {
		return &User{}, fmt.Errorf("UserNotFound")
	}
	return u, nil
}

func (repo *IMUserRepository) AddUser(ctx context.Context, user *User) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	repo.db.addUser(user)
	return nil
}

func (repo *IMUserRepository) UpdateUser(ctx context.Context, user *User) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.updateUser(user)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *IMUserRepository) DeleteUser(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	repo.db.deleteUser(userID)
	return nil
}

func (repo *IMUserRepository) GetHistory(ctx context.Context, userID int64) (*[]solv.Statement, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	h, err := repo.db.getHistory(userID)
	if err != nil {
		return &[]solv.Statement{}, fmt.Errorf("UserNotFound")
	}
	return h, nil
}

func (repo *IMUserRepository) AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.addStatement(userID, statement)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *IMUserRepository) DeleteHistory(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.deleteHistory(userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

//
//
//
//
//

type IMVariableRepository struct {
	mx sync.Mutex
	db *IMStorage
}

func NewImdbVariableRepository() *IMVariableRepository {
	db := GetDefaultStorage()
	return &IMVariableRepository{
		db: db,
	}
}

func (repo *IMVariableRepository) CreateUserVariable(ctx context.Context, userID int64, name string, value float64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.addVariable(name, value, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *IMVariableRepository) CreateUserVariables(
	ctx context.Context,
	userID int64,
	names []string,
	value float64,
) error {

	repo.mx.Lock()
	defer repo.mx.Unlock()

	for i := 0; i < len(names); i++ {
		err := repo.db.addVariable(names[i], value, userID)
		if err != nil {
			return fmt.Errorf("MultyAdd error")
		}
	}

	return nil

}

func (repo *IMVariableRepository) GetUserVariable(ctx context.Context, userID int64, name string) (float64, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	v, err := repo.db.getVariable(name, userID)
	if err != nil {
		return 0.0, fmt.Errorf("UserNotFound")
	}
	return v, nil

}

func (repo *IMVariableRepository) GetUserVariables(
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

func (repo *IMVariableRepository) GetAllUserVariables(ctx context.Context, userID int64) (VMap, error) {
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

func (repo *IMVariableRepository) UpdateUserVariable(ctx context.Context, userID int64, name string, value float64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.updateVariable(name, value, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *IMVariableRepository) UpdateUserVariables(
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

func (repo *IMVariableRepository) DeleteUserVariable(ctx context.Context, userID int64, name string) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.removeVariable(name, userID)
	if err != nil {
		return fmt.Errorf("Got error: %v", err)
	}
	return nil
}

func (repo *IMVariableRepository) DeleteUserVariables(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.clearUserVariables(userID)
	if err != nil {
		return fmt.Errorf("Got error: %v", err)
	}
	return nil
}
