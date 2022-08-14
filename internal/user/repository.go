package user

import (
	"context"
	"fmt"
	"time"

	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

func getUserFromCache(cache *Cache, uID int64) (*User, error) {
	it, ok := cache.Get(fmt.Sprintf("%v", uID))
	if !ok {
		return nil, fmt.Errorf("UserNotFound")
	}

	u, ok := it.(*User)
	return u, nil
}

func addUserToCache(cache *Cache, user *User) error {
	cache.Set(fmt.Sprintf("%v", user.ID), user, time.Hour)
	return nil
}

type userRepository struct {
	cache *Cache
	// db *Connection
}

func NewUserRepository(c *Cache) *userRepository {
	return &userRepository{
		cache: c,
		//db: db,
	}
}

func (repo *userRepository) GetAll(ctx context.Context) (*[]User, error) {
	return nil, fmt.Errorf("Method forbidden.")
}

func (repo *userRepository) Get(ctx context.Context, userID int64) (*User, error) {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return nil, fmt.Errorf("UserNotFound")
	}

	return u, nil
}

func (repo *userRepository) Add(ctx context.Context, user *User) error {
	addUserToCache(repo.cache, user)
	return nil
}

func (repo *userRepository) Update(ctx context.Context, user *User) error {
	repo.cache.Set(fmt.Sprintf("%v", user.ID), user, time.Hour)
	return nil
}

func (repo *userRepository) Delete(ctx context.Context, userID int64) error {
	err := repo.cache.Delete(fmt.Sprintf("%v", userID))
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}
	return nil
}

func (repo *userRepository) GetHistory(ctx context.Context, userID int64) (*History, error) {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return nil, fmt.Errorf("UserNotFound")
	}
	return u.History, nil
}

func (repo *userRepository) AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	*u.History = append(*u.History, *statement)

	addUserToCache(repo.cache, u)

	return nil
}

func (repo *userRepository) DeleteHistory(ctx context.Context, userID int64) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	*u.History = make([]solv.Statement, 0)

	addUserToCache(repo.cache, u)

	return nil
}

func (repo *userRepository) Exist(ctx context.Context, userID int64) bool {
	return repo.cache.Exist(fmt.Sprintf("%v", userID))
}

func (repo *userRepository) Clear(ctx context.Context, userID int64) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	*u.History = make([]solv.Statement, 0)
	u.Variables = make(VMap)

	addUserToCache(repo.cache, u)

	return nil
}

//
//
//
//
//

type variableRepository struct {
	cache *Cache
	// db *Connection
}

func NewVariableRepository(c *Cache) *variableRepository {
	return &variableRepository{
		cache: c,
		//db: db,
	}
}

func (repo *variableRepository) Add(ctx context.Context, userID int64, name string, value float64) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	u.Variables[name] = value

	addUserToCache(repo.cache, u)

	return nil
}

func (repo *variableRepository) AddWithNames(
	ctx context.Context,
	userID int64,
	names []string,
	value float64,
) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	for _, n := range names {
		u.Variables[n] = value
	}

	addUserToCache(repo.cache, u)

	return nil

}

func (repo *variableRepository) Get(ctx context.Context, userID int64, name string) (float64, error) {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return 0.0, fmt.Errorf("UserNotFound")
	}

	v, ok := u.Variables[name]
	if !ok {
		return 0.0, fmt.Errorf("VariableNotFound")
	}

	return v, nil

}

func (repo *variableRepository) GetAll(ctx context.Context, userID int64) (VMap, error) {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return nil, fmt.Errorf("UserNotFound")
	}
	return u.Variables, nil
}

func (repo *variableRepository) GetWithNames(ctx context.Context, userID int64, names []string) (VMap, error) {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return nil, fmt.Errorf("UserNotFound")
	}

	res := make(VMap)
	for _, n := range names {
		v, ok := u.Variables[n]
		if !ok {
			return nil, fmt.Errorf("VariableNotFound")
		}
		res[n] = v
	}

	return res, nil

}

func (repo *variableRepository) Update(ctx context.Context, userID int64, name string, value float64) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	u.Variables[name] = value

	addUserToCache(repo.cache, u)

	return nil
}

func (repo *variableRepository) UpdateWithNames(ctx context.Context, userID int64, names []string, values []float64) error {
	if len(names) != len(values) {
		return fmt.Errorf("InputVarError")
	}
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	for i, n := range names {
		u.Variables[n] = values[i]
	}

	addUserToCache(repo.cache, u)
	return nil
}

func (repo *variableRepository) Delete(ctx context.Context, userID int64, name string) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	delete(u.Variables, name)

	addUserToCache(repo.cache, u)

	return nil
}

func (repo *variableRepository) DeleteWithNames(ctx context.Context, userID int64, names []string) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	for _, n := range names {
		delete(u.Variables, n)
	}

	addUserToCache(repo.cache, u)
	return nil
}

func (repo *variableRepository) DeleteAll(ctx context.Context, userID int64) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return fmt.Errorf("UserNotFound")
	}

	u.Variables = make(VMap)

	addUserToCache(repo.cache, u)

	return nil
}
