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

func (us userService) Add(ctx context.Context, user *User) error {
	err := us.userReposiory.Add(ctx, user)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us userService) GetAll(ctx context.Context) (*[]User, error) {
	return nil, fmt.Errorf("Method forbidden.")
}

func (us userService) Get(ctx context.Context, userID int64) (*User, error) {
	u, err := us.userReposiory.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("Got error from db: %v", err)
	}
	return u, nil
}
func (us userService) Update(ctx context.Context, user *User) error {
	err := us.userReposiory.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us userService) Delete(ctx context.Context, userID int64) error {
	err := us.userReposiory.Delete(ctx, userID)
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

func (us userService) GetHistory(ctx context.Context, userID int64) (*History, error) {
	h, err := us.userReposiory.GetHistory(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("Got error from db: %v", err)
	}
	return h, nil
}

func (us userService) DeleteHistory(ctx context.Context, userID int64) error {
	err := us.userReposiory.DeleteHistory(ctx, userID)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us userService) Exist(ctx context.Context, userID int64) (bool, error) {
	return us.userReposiory.Exist(ctx, userID), nil
}

func (us userService) Clear(ctx context.Context, userID int64) error {
	err := us.userReposiory.Clear(ctx, userID)
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

func (us variableService) Add(ctx context.Context, userID int64, name string, value float64) error {
	err := us.variableReposiory.Add(ctx, userID, name, value)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) AddWithNames(ctx context.Context, userID int64, names []string, value float64) error {
	err := us.variableReposiory.AddWithNames(ctx, userID, names, value)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) Get(ctx context.Context, userID int64, name string) (float64, error) {
	v, err := us.variableReposiory.Get(ctx, userID, name)
	if err != nil {
		return 0.0, fmt.Errorf("Got error from db: %v", err)
	}
	return v, nil
}

func (us variableService) GetAll(ctx context.Context, userID int64) (VMap, error) {
	vs, err := us.variableReposiory.GetAll(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("Got error from db: %v", err)
	}
	return vs, nil
}

func (us variableService) GetWithNames(ctx context.Context, userID int64, names []string) (VMap, error) {
	vs, err := us.variableReposiory.GetWithNames(ctx, userID, names)
	if err != nil {
		return nil, fmt.Errorf("Got error from db: %v", err)
	}
	return vs, nil
}

func (us variableService) Update(ctx context.Context, userID int64, name string, value float64) error {
	err := us.variableReposiory.Update(ctx, userID, name, value)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) UpdateWithNames(ctx context.Context, userID int64, names []string, value float64) error {
	values := make([]float64, len(names))
	for i := range names {
		values[i] = value
	}
	err := us.variableReposiory.UpdateWithNames(ctx, userID, names, values)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) Delete(ctx context.Context, userID int64, name string) error {
	err := us.variableReposiory.Delete(ctx, userID, name)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) DeleteWithNames(ctx context.Context, userID int64, names []string) error {
	err := us.variableReposiory.DeleteWithNames(ctx, userID, names)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}

func (us variableService) DeleteAll(ctx context.Context, userID int64) error {
	err := us.variableReposiory.DeleteAll(ctx, userID)
	if err != nil {
		return fmt.Errorf("Got error from db: %v", err)
	}
	return nil
}
