package user

import (
	"context"

	"emperror.dev/errors"
	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

type userService struct {
	userReposiory UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{repo}
}

func (us userService) Add(ctx context.Context, tgID int64, password string) error {
	err := us.userReposiory.Add(ctx, tgID, password)
	if err != nil {
		return errors.Wrap(err, "Got error with user add to cache")
	}
	return nil
}

func (us userService) GetAll(ctx context.Context) (*[]User, error) {
	return nil, errors.New("Method forbidden.")
}

func (us userService) Get(ctx context.Context, tgID int64) (*User, error) {
	u, err := us.userReposiory.Get(ctx, tgID)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with getting user from cache")
	}
	return u, nil
}
func (us userService) Update(ctx context.Context, tgID int64, password string) error {
	err := us.userReposiory.Update(ctx, tgID, password)
	if err != nil {
		return errors.Wrap(err, "Got error with updating user in cache")
	}
	return nil
}

func (us userService) Delete(ctx context.Context, tgID int64) error {
	err := us.userReposiory.Delete(ctx, tgID)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting user from cahce")
	}
	return nil
}

func (us userService) AddStatement(ctx context.Context, tgID int64, statement *solv.Statement) error {
	err := us.userReposiory.AddStatement(ctx, tgID, statement)
	if err != nil {
		return errors.Wrap(err, "Got error with adding statement to cahce")
	}
	return nil
}

func (us userService) AddStatements(ctx context.Context, tgID int64, statements *[]solv.Statement) error {
	err := us.userReposiory.AddStatements(ctx, tgID, statements)
	if err != nil {
		return errors.Wrap(err, "Got error with adding statement to cahce")
	}
	return nil
}

func (us userService) GetHistory(ctx context.Context, tgID int64) (*History, error) {
	h, err := us.userReposiory.GetHistory(ctx, tgID)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with getting user history")
	}
	return h, nil
}

func (us userService) Clear(ctx context.Context, tgID int64) error {
	err := us.userReposiory.Clear(ctx, tgID)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting user history from cache")
	}
	return nil
}

func (us userService) Exist(ctx context.Context, tgID int64) (bool, error) {
	ok, err := us.userReposiory.Exist(ctx, tgID)
	if err != nil {
		return false, errors.Wrap(err, "Exist req fail")
	}
	return ok, nil
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

func (us variableService) Get(ctx context.Context, tgID int64, name string) (float64, error) {
	v, err := us.variableReposiory.Get(ctx, tgID, name)
	if err != nil {
		return 0.0, errors.Wrap(err, "Got error with getting variable from cache")
	}
	return v, nil
}

func (us variableService) GetAll(ctx context.Context, tgID int64) (VMap, error) {
	vs, err := us.variableReposiory.GetAll(ctx, tgID)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with getting all variables from cache")
	}
	return vs, nil
}

func (us variableService) GetWithNames(ctx context.Context, tgID int64, names []string) (VMap, error) {
	vs, err := us.variableReposiory.GetWithNames(ctx, tgID, names)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with getting multiple variables from cache")
	}
	return vs, nil
}

func (us variableService) Update(ctx context.Context, tgID int64, name string, value float64) error {
	err := us.variableReposiory.Update(ctx, tgID, name, value)
	if err != nil {
		return errors.Wrap(err, "Got error with updating variable in cache")
	}
	return nil
}

func (us variableService) UpdateWithNames(ctx context.Context, tgID int64, names []string, value float64) error {
	err := us.variableReposiory.UpdateWithNames(ctx, tgID, names, value)
	if err != nil {
		return errors.Wrap(err, "Got error with updating multiple variables in cache")
	}
	return nil
}

func (us variableService) Delete(ctx context.Context, tgID int64, name string) error {
	err := us.variableReposiory.Delete(ctx, tgID, name)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting variable in cache")
	}
	return nil
}

func (us variableService) DeleteWithNames(ctx context.Context, tgID int64, names []string) error {
	err := us.variableReposiory.DeleteWithNames(ctx, tgID, names)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting multiple variables from cache")
	}
	return nil
}

func (us variableService) DeleteAll(ctx context.Context, tgID int64) error {
	err := us.variableReposiory.DeleteAll(ctx, tgID)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting all variables from cache")
	}
	return nil
}
