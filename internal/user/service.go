package user

import (
	"context"
	"fmt"

	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

type userService struct {
	userReposiory UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{repo}
}

func (us userService) FetchUsers(ctx context.Context) (*[]User, error) {
	usrs, err := us.userReposiory.GetUsers(ctx)
	if err != nil {
		return &[]User{}, fmt.Errorf("Got error from db: %v", err)
	}
	return usrs, nil
}

func (us userService) FetchUser(ctx context.Context, userID int64) (*User, error) {
	u, err := us.userReposiory.GetUser(ctx, userID)
	if err != nil {
		return &User{}, fmt.Errorf("Got error from db: %v", err)
	}
	return u, nil
}

func (us userService) CreateUser(ctx context.Context, user *User) error {
	err := us.userReposiory.AddUser(ctx, user)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us userService) SetUser(ctx context.Context, user *User) error {
	err := us.userReposiory.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us userService) DeleteUser(ctx context.Context, userID int64) error {
	err := us.userReposiory.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us userService) AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error {
	err := us.userReposiory.AddStatement(ctx, userID, statement)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us userService) FetchHistory(ctx context.Context, userID int64) (*[]solv.Statement, error) {
	h, err := us.userReposiory.GetHistory(ctx, userID)
	if err != nil {
		return &[]solv.Statement{}, fmt.Errorf("Got error from db: %v", err)
	}
	return h, nil
}

func (us userService) ClearHistory(ctx context.Context, userID int64) error {
	err := us.userReposiory.DeleteHistory(ctx, userID)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

//
//
// Variables part
//
//

type variableService struct {
	variableReposiory VariableRepository
}

func NewVariableService(repo VariableRepository) *variableService {
	return &variableService{repo}
}

func (us variableService) AddUserVariable(ctx context.Context, userID int64, name string, value float64) error {
	err := us.variableReposiory.CreateUserVariable(ctx, userID, name, value)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) AddUserVariables(ctx context.Context, userID int64, names []string, value float64) error {
	err := us.variableReposiory.CreateUserVariables(ctx, userID, names, value)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) FetchUserVariable(ctx context.Context, userID int64, name string) (float64, error) {
	v, err := us.variableReposiory.GetUserVariable(ctx, userID, name)
	if err != nil {
		return 0.0, fmt.Errorf("Got error from db: %v", err)
	}
	return v, nil
}

func (us variableService) FetchUserVariables(ctx context.Context, userID int64, names []string) (VMap, error) {
	vs, err := us.variableReposiory.GetUserVariables(ctx, userID, names)
	if err != nil {
		return VMap{}, fmt.Errorf("Got error from db: %v", err)
	}
	return vs, nil
}

func (us variableService) FetchAllUserVariables(ctx context.Context, userID int64) (VMap, error) {
	vs, err := us.variableReposiory.GetAllUserVariables(ctx, userID)
	if err != nil {
		return VMap{}, fmt.Errorf("Got error from db: %v", err)
	}
	return vs, nil
}

func (us variableService) SetUserVariable(ctx context.Context, userID int64, name string, value float64) error {
	err := us.variableReposiory.UpdateUserVariable(ctx, userID, name, value)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) SetUserVariables(ctx context.Context, userID int64, names []string, value float64) error {
	values := make([]float64, len(names))
	for i := range names {
		values[i] = value
	}
	err := us.variableReposiory.UpdateUserVariables(ctx, userID, names, values)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) DeleteUserVariable(ctx context.Context, userID int64, name string) error {
	err := us.variableReposiory.DeleteUserVariable(ctx, userID, name)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) ClearUserVariables(ctx context.Context, userID int64) error {
	err := us.variableReposiory.DeleteUserVariables(ctx, userID)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}
