package user

import (
	"context"

	"github.com/jmoiron/sqlx"

	"emperror.dev/errors"
	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

var itemNotFoundErr *ItemNotFoundError = &ItemNotFoundError{}

// Get variables
type dbVariable struct {
	Name  string  `db:"name"`
	Value float64 `db:"value"`
}

type dbUser struct {
	Id       int    `db:"id"`
	TgID     int64  `db:"tg_id"`
	Password string `db:"password"`
	Created  string `db:"created_at"`
}

type dbStatement struct {
	Id       int     `db:"id"`
	Equation string  `db:"equation"`
	Value    float64 `db:"value"`
	Created  string  `db:"created_at"`
}

type userRepository struct {
	cache Cache
	db    *sqlx.DB
}

func NewUserRepository(c Cache, db *sqlx.DB) *userRepository {
	return &userRepository{
		cache: c,
		db:    db,
	}
}
func (repo userRepository) getFilledStatement(ctx context.Context, dbs *dbStatement, variables *VMap) (*solv.Statement, error) {

	query := `SELECT DISTINCT variables.name as "name", variables.value as "value"
		FROM "variables" INNER JOIN
			"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
			"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
			"users"                 ON "statements".user_id                 = "users".id
		WHERE "statements".id = $1;`

	dbVars := []dbVariable{}
	err := repo.db.Select(&dbVars, query, dbs.Id)
	if err != nil {
		return nil, errors.Wrap(err, "Select vars request fail")
	}

	stVars := make([]string, 0)
	for _, v := range dbVars {
		(*variables)[v.Name] = v.Value
		stVars = append(stVars, v.Name)
	}

	return &solv.Statement{
		Id:        dbs.Id,
		Variables: stVars,
		Equation:  dbs.Equation,
		Value:     dbs.Value,
	}, nil
}

func (repo userRepository) getHistoryWithTgID(ctx context.Context, tgID int64) (*History, *VMap, error) {
	query := `SELECT  "statements".id, "statements".equation, "statements".value, "statements".created_at FROM "statements" INNER JOIN 
		"users" ON "statements".user_id = "users".id
		WHERE "users".tg_id = $1;`

	return repo.getHistory(ctx, query, tgID)
}

func (repo userRepository) getHistoryWithDBuID(ctx context.Context, uID int) (*History, *VMap, error) {
	query := `SELECT  "statements".id, "statements".equation, "statements".value, "statements".created_at FROM "statements" INNER JOIN 
		"users" ON "statements".user_id = "users".id
		WHERE "users".id = $1;`

	return repo.getHistory(ctx, query, uID)
}

func (repo userRepository) getHistory(ctx context.Context, query string, id interface{}) (*History, *VMap, error) {

	sts := []dbStatement{}

	err := repo.db.Select(&sts, query, id)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Select history request fail")
	}

	variables := VMap{}
	h := make(History, len(sts))
	for i, dbs := range sts {
		s, err := repo.getFilledStatement(ctx, &dbs, &variables)
		if err != nil {
			return nil, nil, err
		}
		h[i] = *s
	}
	return &h, &variables, nil
}

func (repo userRepository) getFilledUser(ctx context.Context, user *dbUser) (*User, error) {

	h, vs, err := repo.getHistoryWithDBuID(ctx, user.Id)
	if err != nil {
		return nil, errors.Wrap(err, "History fetch fail.")
	}

	return &User{user.Id, user.TgID, user.Password, h, *vs}, nil
}

func (repo userRepository) Add(ctx context.Context, tgID int64, password string) error {

	query := `INSERT INTO users (tg_id, password, created_at)
		VALUES ($1, $2, now());`

	repo.db.MustExec(query, tgID, password)

	return nil
}

func (repo userRepository) GetAll(ctx context.Context) (*[]User, error) {
	query := `SELECT * FROM "users";`

	dbUsers := []dbUser{}
	err := repo.db.Select(&dbUsers, query)
	if err != nil {
		return nil, errors.Wrap(err, "Get request fail")
	}

	users := make([]User, len(dbUsers))
	for i, dbu := range dbUsers {
		u, err := repo.getFilledUser(ctx, &dbu)
		if err != nil {
			return nil, errors.Wrap(err, "User fill request fail.")
		}
		users[i] = *u
	}
	return &users, nil
}

func (repo *userRepository) Get(ctx context.Context, tgID int64) (*User, error) {

	query := `SELECT * FROM "users" WHERE "users".tg_id = $1;`

	dbu := dbUser{}
	err := repo.db.Get(&dbu, query, tgID)
	if err != nil {
		return nil, errors.Wrap(err, "Get request fail")
	}

	u, err := repo.getFilledUser(context.TODO(), &dbu)
	if err != nil {
		return nil, errors.Wrap(err, `User additional requests fail`)
	}

	return u, nil
}

func (repo *userRepository) Update(ctx context.Context, tgID int64, password string) error {

	query := `UPDATE "users"
		SET password = $2
		WHERE "users".tg_id = $1;`

	repo.db.MustExec(query, tgID, password)

	return nil
}

func (repo *userRepository) Delete(ctx context.Context, tgID int64) error {
	tx := repo.db.MustBegin()

	query := `DELETE FROM "variables" 
		WHERE "variables".id IN (
			SELECT "variables".id
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id
			WHERE "users".tg_id = $1
		);`

	tx.MustExec(query, tgID)

	query = `DELETE FROM "statements" 
		WHERE "statements".id IN (
			SELECT "statements".id
			FROM "statements" INNER JOIN "users" ON "statements".user_id = "users".id
			WHERE "users".tg_id = $1
		)`

	tx.MustExec(query, tgID)

	query = `DELETE FROM "users" WHERE "users".tg_id = $1`

	tx.MustExec(query, tgID)

	tx.Commit()

	return nil
}

func (repo *userRepository) GetHistory(ctx context.Context, tgID int64) (*History, error) {
	h, _, err := repo.getHistoryWithTgID(ctx, tgID)
	if err != nil {
		return nil, errors.Wrap(err, "Get history from DB fail.")
	}

	return h, nil
}

func (repo *userRepository) AddStatement(ctx context.Context, tgID int64, statement *solv.Statement) error {
	tx := repo.db.MustBegin()

	query := `INSERT INTO statements (user_id, equation, value, created_at)
			VALUES ((SELECT id FROM "users" WHERE "users".tg_id = $1), $2, $3, now()) RETURNING id;`

	var newStID int
	err := tx.Get(&newStID, query, tgID, statement.Equation, statement.Value)
	if err != nil {
		return errors.Wrap(err, "Add statement request to db fail")
	}

	query = `select * from set_var($1, $2, $3, $4);`
	for _, varName := range statement.Variables {
		tx.MustExec(query, tgID, newStID, varName, statement.Value)
	}

	tx.Commit()

	return nil
}

func (repo *userRepository) Exist(ctx context.Context, tgID int64) (bool, error) {
	query := `SELECT 1
		FROM "users"
		WHERE "users".tg_id = $1;
		`
	var res int
	repo.db.Get(&res, query, tgID)

	return res == 1, nil
}

func (repo *userRepository) Clear(ctx context.Context, tgID int64) error {

	query := `DELETE FROM "variables" 
		WHERE "variables".id IN (
			SELECT "variables".id
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id
			WHERE "users".tg_id = $1
		);`

	repo.db.MustExec(query, tgID)

	query = `DELETE FROM "statements" 
		WHERE "statements".id IN (
			SELECT "statements".id
			FROM "statements" INNER JOIN "users" ON "statements".user_id = "users".id
			WHERE "users".tg_id = $1
		);`

	repo.db.MustExec(query, tgID)

	return nil
}

//
//
//
//
//

type variableRepository struct {
	cache Cache
	db    *sqlx.DB
}

func NewVariableRepository(c Cache, db *sqlx.DB) *variableRepository {
	return &variableRepository{
		cache: c,
		db:    db,
	}
}

func (repo *variableRepository) Get(ctx context.Context, tgID int64, name string) (float64, error) {

	query := `SELECT DISTINCT variables.name as "name", variables.value as "value"
		FROM variables INNER JOIN
			"statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
			statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
			users                 ON statements.user_id                 = users.id
		WHERE variables.name = $1 AND users.tg_id = $2 AND variables.value IS NOT NULL;`

	v := dbVariable{}
	err := repo.db.Get(&v, query, name, tgID)
	if err != nil {
		return 0.0, errors.Wrap(err, "Get request fail")
	}

	return v.Value, nil
}

func (repo *variableRepository) GetAll(ctx context.Context, tgID int64) (VMap, error) {
	res := VMap{}

	query := `SELECT DISTINCT variables.name as "name", variables.value as "value"
		FROM "variables" INNER JOIN
			"statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
			statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
			users                 ON statements.user_id                 = users.id
		WHERE users.tg_id = $1 AND variables.value IS NOT NULL;`

	vs := []dbVariable{}
	err := repo.db.Select(&vs, query, tgID)
	if err != nil {
		return res, errors.Wrap(err, "Select request fail")
	}
	for _, v := range vs {
		res[v.Name] = v.Value
	}

	return res, nil
}

func (repo *variableRepository) GetWithNames(ctx context.Context, tgID int64, names []string) (VMap, error) {
	res := VMap{}

	query := `SELECT DISTINCT variables.name as "name", variables.value as "value"
		FROM variables INNER JOIN
			"statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
			statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
			users                 ON statements.user_id                 = users.id
		WHERE variables.name IN (?) AND users.tg_id = ? AND variables.value IS NOT NULL;`

	query, args, err := sqlx.In(query, names, tgID)
	if err != nil {
		return res, errors.Wrap(err, "Query build fail")
	}

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = repo.db.Rebind(query)

	vs := []dbVariable{}
	err = repo.db.Select(&vs, query, args...)
	if err != nil {
		return res, errors.Wrap(err, "Select request fail")
	}

	for _, v := range vs {
		res[v.Name] = v.Value
	}

	return res, nil

}

func (repo *variableRepository) Update(ctx context.Context, tgID int64, name string, value float64) error {

	query := `UPDATE variables
			SET value = $1
			WHERE id = (
				SELECT DISTINCT variables.id
				FROM "variables" INNER JOIN
					"statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
					statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
					users                 ON statements.user_id                 = users.id
			WHERE users.tg_id = $2 AND variables.name = $3
			) RETURNING id;`

	res := repo.db.MustExec(query, value, tgID, name)
	am, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Update request fail")
	} else if am == 0 {
		return errors.New("Variable not found")
	}

	return nil
}

func (repo *variableRepository) UpdateWithNames(ctx context.Context, tgID int64, names []string, value float64) error {

	query := `UPDATE "variables"
		SET value = ?
		WHERE id IN (
			SELECT "variables".id
			FROM "variables" INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
				"statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
				"users"                 ON "statements".user_id                 = "users".id
		WHERE "users".tg_id = ? AND "variables".name IN (?)
		) RETURNING id;`

	query, args, err := sqlx.In(query, value, tgID, names)
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

func (repo *variableRepository) Delete(ctx context.Context, tgID int64, name string) error {

	query := `UPDATE variables
		SET value = NULL
		WHERE variables.id IN (
			SELECT variables.id
			FROM variables INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
				statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
				users                 ON statements.user_id                 = users.id

			WHERE users.tg_id = $1 AND variables.name = $2
		) RETURNING id;`

	res := repo.db.MustExec(query, tgID, name)
	am, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Delete [Update] request fail")
	} else if am == 0 {
		return errors.New("Variable not found")
	}

	return nil
}

func (repo *variableRepository) DeleteWithNames(ctx context.Context, tgID int64, names []string) error {

	query := `UPDATE variables
		SET value = NULL
		WHERE variables.id IN (
			SELECT variables.id
			FROM variables INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
				statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
				users                 ON statements.user_id                 = users.id

			WHERE users.tg_id = ? AND (variables.name IN (?))
		) RETURNING id;`

	query, args, err := sqlx.In(query, tgID, names)
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

func (repo *variableRepository) DeleteAll(ctx context.Context, tgID int64) error {

	query := `UPDATE variables
		SET value = NULL
		WHERE id IN (
			SELECT  variables.id
			FROM variables INNER JOIN
				"statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
				statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
				users                 ON statements.user_id                 = users.id
			WHERE users.tg_id = $1
		);`

	res := repo.db.MustExec(query, tgID)
	am, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Update request fail")
	} else if am == 0 {
		return errors.New("Variable not found")
	}

	return nil
}
