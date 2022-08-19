package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"emperror.dev/errors"
	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

func getUserFromCache(cache *Cache, uID int64) (*User, error) {
	it, err := cache.Get(fmt.Sprintf("%v", uID))
	if err != nil {
		return nil, err
	}

	u, ok := it.(*User)
	if !ok {
		return nil, errors.New("Wrong type")
	}
	return u, nil
}

func addUserToCache(cache *Cache, user *User) error {
	cache.Set(fmt.Sprintf("%v", user.id), user, time.Hour)
	return nil
}

type userRepository struct {
	cache *Cache
	db    *sqlx.DB
}

func NewUserRepository(c *Cache, db *sqlx.DB) *userRepository {
	return &userRepository{
		cache: c,
		db:    db,
	}
}

func (repo *userRepository) GetAll(ctx context.Context) (*[]User, error) {

	return nil, errors.New("Method forbidden.")
}

func (repo *userRepository) Get(ctx context.Context, userID int64) (*User, error) {
	u, err := getUserFromCache(repo.cache, userID)
	if !errors.As(err, &ItemNotFoundError{}) {
		return nil, errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
	}

	return u, nil
}

func (repo *userRepository) Add(ctx context.Context, user *User) error {
	addUserToCache(repo.cache, user)
	return nil
}

func (repo *userRepository) Update(ctx context.Context, user *User) error {
	repo.cache.Set(fmt.Sprintf("%v", user.id), user, time.Hour)
	return nil
}

func (repo *userRepository) Delete(ctx context.Context, userID int64) error {
	err := repo.cache.Delete(fmt.Sprintf("%v", userID))
	if !errors.As(err, &ItemNotFoundError{}) {
		return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
	}
	return nil
}

func (repo *userRepository) GetHistory(ctx context.Context, userID int64) (*History, error) {
	u, err := getUserFromCache(repo.cache, userID)
	if !errors.As(err, &ItemNotFoundError{}) {
		return nil, errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
	}
	return u.History, nil
}

func (repo *userRepository) AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error {
	u, err := getUserFromCache(repo.cache, userID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
	}

	*u.History = append(*u.History, *statement)

	addUserToCache(repo.cache, u)

	return nil
}

func (repo *userRepository) AddStatements(ctx context.Context, userID int64, statement []*solv.Statement) error {
	u, err := getUserFromCache(repo.cache, userID)
	if !errors.As(err, &ItemNotFoundError{}) {
		return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
	}

	//*u.History = append(*u.History, *statement)

	addUserToCache(repo.cache, u)

	return nil
}

func (repo *userRepository) DeleteHistory(ctx context.Context, userID int64) error {
	u, err := getUserFromCache(repo.cache, userID)
	if !errors.As(err, &ItemNotFoundError{}) {
		return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
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
	if !errors.As(err, &ItemNotFoundError{}) {
		return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
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
	db    *sqlx.DB
}

func NewVariableRepository(c *Cache, db *sqlx.DB) *variableRepository {
	return &variableRepository{
		cache: c,
		db:    db,
	}
}

func (repo *variableRepository) Get(ctx context.Context, userID int64, name string) (float64, error) {
	/*
		u, err := getUserFromCache(repo.cache, userID)

			var itemNotFoundError *ItemNotFoundError
			if !errors.As(err, &itemNotFoundError) {
				return 0.0, errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
			} else if err == nil {
				v, ok := u.Variables[name]
				if !ok {
					return 0.0, errors.Wrap(err, fmt.Sprintf("Getting variable with id %s fail", name))
				}
				return v, nil
			}
	*/

	type variable struct {
		Value float64
	}

	query := `SELECT DISTINCT variables.value as "value"
		FROM "variables" INNER JOIN
			"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
			"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
			"users"                 ON "statements".user_id                 = "users".id
		WHERE "variables".name = $1 AND "users".tg_id = $2;`

	v := variable{}
	err := repo.db.Get(&v, query, name, userID)
	if err != nil {
		return 0.0, errors.Wrap(err, "Get request fail")
	}

	return v.Value, nil
}

func (repo *variableRepository) GetAll(ctx context.Context, userID int64) (VMap, error) {
	res := VMap{}

	/*
		u, err := getUserFromCache(repo.cache, userID)

		var itemNotFoundError *ItemNotFoundError
		if !errors.As(err, &itemNotFoundError) {
			return res, errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
		} else if err == nil {
			vs := u.Variables
			return vs, nil
		}
	*/

	type variable struct {
		Name  string
		Value float64
	}

	query := `SELECT DISTINCT variables.name as "name", variables.value as "value"
		FROM "variables" INNER JOIN
			"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
			"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
			"users"                 ON "statements".user_id                 = "users".id
		WHERE "users".tg_id = $1;`

	vs := []variable{}
	err := repo.db.Select(&vs, query, userID)
	if err != nil {
		return res, errors.Wrap(err, "Select request fail")
	}
	for _, v := range vs {
		res[v.Name] = v.Value
	}

	return res, nil
}

func (repo *variableRepository) GetWithNames(ctx context.Context, userID int64, names []string) (VMap, error) {
	res := VMap{}

	/*
		u, err := getUserFromCache(repo.cache, userID)

		var itemNotFoundError *ItemNotFoundError
		if !errors.As(err, &itemNotFoundError) {
			return res, errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
		} else if err == nil {
			for _, n := range names {
				v, ok := u.Variables[n]
				if !ok {
					return nil, errors.Wrap(err, fmt.Sprintf("Getting variable with id %s fail", n))
				}
				res[n] = v
			}

			return res, nil
		}
	*/

	type variable struct {
		Name  string
		Value float64
	}

	query := `SELECT DISTINCT variables.name as "name", variables.value as "value"
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id
			WHERE "variables".name IN (?) AND "users".tg_id = ?;`

	query, args, err := sqlx.In(query, names, userID)
	if err != nil {
		return res, errors.Wrap(err, "Query build fail")
	}

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = repo.db.Rebind(query)

	vs := []variable{}
	err = repo.db.Select(&vs, query, args...)
	if err != nil {
		return res, errors.Wrap(err, "Select request fail")
	}

	for _, v := range vs {
		res[v.Name] = v.Value
	}

	return res, nil

}

func (repo *variableRepository) Update(ctx context.Context, userID int64, name string, value float64) error {
	/*
			cache
		u.Variables[name] = value

		addUserToCache(repo.cache, u)


	*/

	query := `UPDATE "variables"
		SET value = $1
		WHERE EXISTS (
			SELECT  *
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id
		WHERE "users".tg_id = $2
		) AND ("variables".name = $3);`

	res := repo.db.MustExec(query, value, userID, name)
	am, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Update request fail")
	} else if am == 0 {
		return errors.New("Variable not found")
	}

	return nil
}

func (repo *variableRepository) UpdateWithNames(ctx context.Context, userID int64, names []string, value float64) error {
	/*
		u, err := getUserFromCache(repo.cache, userID)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
		}

		for i, n := range names {
			u.Variables[n] = values[i]
		}

		addUserToCache(repo.cache, u)
		return nil
	*/

	query := `UPDATE "variables"
		SET value = ?
		WHERE EXISTS (
			SELECT  *
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id
		WHERE "users".tg_id = ?
		) AND "variables".name IN (?);`

	query, args, err := sqlx.In(query, value, userID, names)
	if err != nil {
		return errors.Wrap(err, "Query build fail")
	}

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = repo.db.Rebind(query)

	res := repo.db.MustExec(query, args...)
	am, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Update request fail")
	} else if am == 0 {
		return errors.New("Variables not found")
	}

	return nil
}

func (repo *variableRepository) Delete(ctx context.Context, userID int64, name string) error {
	/*
		u, err := getUserFromCache(repo.cache, userID)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
		}

		delete(u.Variables, name)

		addUserToCache(repo.cache, u)

		return nil
	*/

	query := `UPDATE "variables"
		SET value = NULL
		WHERE EXISTS (
			SELECT  *
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id

			WHERE "users".tg_id = $1
		) AND "variables".name=$2;`

	res := repo.db.MustExec(query, userID, name)
	am, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Delete [Update] request fail")
	} else if am == 0 {
		return errors.New("Variable not found")
	}

	return nil
}

func (repo *variableRepository) DeleteWithNames(ctx context.Context, userID int64, names []string) error {
	/*
		u, err := getUserFromCache(repo.cache, userID)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
		}

		for _, n := range names {
			delete(u.Variables, n)
		}

		addUserToCache(repo.cache, u)
		return nil
	*/

	query := `UPDATE "variables"
		SET value = NULL
		WHERE EXISTS (
			SELECT  *
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id
		WHERE "users".tg_id = ?
		) AND "variables".name IN (?);`

	query, args, err := sqlx.In(query, userID, names)
	if err != nil {
		return errors.Wrap(err, "Query build fail")
	}

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = repo.db.Rebind(query)

	res := repo.db.MustExec(query, args...)
	am, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Delete [Update] request fail")
	} else if am == 0 {
		return errors.New("Variables not found")
	}

	return nil
}

func (repo *variableRepository) DeleteAll(ctx context.Context, userID int64) error {
	/*
		u, err := getUserFromCache(repo.cache, userID)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", userID))
		}

		u.Variables = make(VMap)

		addUserToCache(repo.cache, u)

		return nil
	*/
	query := `UPDATE "variables"
		SET value = NULL
		WHERE EXISTS (
			SELECT  *
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id
		WHERE "users".tg_id = $1
		);`

	res := repo.db.MustExec(query, userID)
	am, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Update request fail")
	} else if am == 0 {
		return errors.New("Variable not found")
	}

	return nil
}
