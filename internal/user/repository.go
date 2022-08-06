package user

/*
import (
	"context"
	"fmt"
	"sync"

	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)
type userRepository struct {
	mx sync.Mutex
	db *storage
}

func NewUserRepository() *userRepository {
	db := _default_storage
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetUsers(ctx context.Context) (*[]User, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	resp := repo.db.getUsers()
	return resp, nil
}

func (repo *userRepository) GetUser(ctx context.Context, userID int64) (*User, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	u, ok := repo.db.getUser(userID)
	if !ok {
		return &User{}, fmt.Errorf("UserNotFound")
	}
	return u, nil
}

func (repo *userRepository) AddUser(ctx context.Context, user *User) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	repo.db.addUser(user)
	return nil
}

func (repo *userRepository) UpdateUser(ctx context.Context, user *User) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.updateUser(user)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *userRepository) DeleteUser(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	repo.db.deleteUser(userID)
	return nil
}

func (repo *userRepository) GetHistory(ctx context.Context, userID int64) (*[]solv.Statement, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	h, err := repo.db.getHistory(userID)
	if err != nil {
		return &[]solv.Statement{}, fmt.Errorf("UserNotFound")
	}
	return h, nil
}

func (repo *userRepository) AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.addStatement(userID, statement)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *userRepository) DeleteHistory(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.deleteHistory(userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

type variableRepository struct {
	mx sync.Mutex
	db *storage
}

func NewVariableRepository() *variableRepository {
	db := _default_storage
	return &variableRepository{
		db: db,
	}
}

func (repo *variableRepository) CreateUserVariable(ctx context.Context, userID int64, name string, value float64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.addVariable(name, value, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *variableRepository) CreateUserVariables(
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

func (repo *variableRepository) GetUserVariable(ctx context.Context, userID int64, name string) (float64, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	v, err := repo.db.getVariable(name, userID)
	if err != nil {
		return 0.0, fmt.Errorf("UserNotFound")
	}
	return v, nil

}

func (repo *variableRepository) GetUserVariables(
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

func (repo *variableRepository) GetAllUserVariables(ctx context.Context, userID int64) (VMap, error) {
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

func (repo *variableRepository) UpdateUserVariable(ctx context.Context, userID int64, name string, value float64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.updateVariable(name, value, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *variableRepository) UpdateUserVariables(
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

func (repo *variableRepository) DeleteUserVariable(ctx context.Context, userID int64, name string) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.removeVariable(name, userID)
	if err != nil {
		return fmt.Errorf("Got error: %v", err)
	}
	return nil
}

func (repo *variableRepository) DeleteUserVariables(ctx context.Context, userID int64) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	err := repo.db.clearUserVariables(userID)
	if err != nil {
		return fmt.Errorf("Got error: %v", err)
	}
	return nil
}
*/
