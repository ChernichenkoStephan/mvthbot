package user

import (
	"context"
	"fmt"
	"time"

	"emperror.dev/errors"
	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func (db Database) getVariableFromCache(tgID int64, name string) (float64, error) {
	u, err := db.getUserFromCache(tgID)
	if err != nil {
		return 0.0, errors.Wrap(err, `User fetch from cache for var fail`)
	}

	v, ok := u.Variables[name]
	if !ok {
		return 0.0, errors.New(`Variiable not found`)
	}
	return v, nil
}

func (db Database) getUserFromCache(tgID int64) (*User, error) {
	it, err := db.cache.Get(fmt.Sprintf("%v", tgID))
	if err != nil {
		return nil, err
	}

	u, ok := it.(*User)
	if !ok {
		return nil, errors.New("Wrong type")
	}
	return u, nil
}

func (db Database) userToCache(user *User) error {
	db.cache.Set(fmt.Sprintf("%v", user.TelegramID), user, time.Hour)
	return nil
}

func NewDB(usersRepo UserRepository, varsRepo VariableRepository, cache Cache, conn Connector, logger *zap.SugaredLogger) *Database {
	return &Database{
		usersRepo: usersRepo,
		varsRepo:  varsRepo,
		cache:     cache,
		conn:      conn,
		lg:        logger,
	}
}

func (db Database) Add(ctx context.Context, tgID int64, password string) error {
	err := db.usersRepo.Add(ctx, tgID, password)
	if err != nil {
		return errors.Wrap(err, "Got error with user add to cache")
	}

	err = db.userToCache(NewUserWithPassword(tgID, password))
	if err != nil {
		db.lg.Errorf("Save to cache failed %s", err.Error())
	}

	return nil
}

func (db Database) GetAll(ctx context.Context) (*[]User, error) {
	users, err := db.usersRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with user add to cache")
	}

	for _, u := range *users {
		err = db.userToCache(&u)
		if err != nil {
			db.lg.Errorf("Save to cache failed %s", err.Error())
		}
	}

	return users, nil
}

func (db Database) Get(ctx context.Context, tgID int64) (*User, error) {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		db.lg.Info("From cache")
		return u, nil
	} else if !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	u, err = db.usersRepo.Get(ctx, tgID)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with getting user from cache")
	}

	err = db.userToCache(u)
	if err != nil {
		db.lg.Errorf("Save to cache failed %s", err.Error())
	}

	return u, nil
}

func (db Database) Update(ctx context.Context, tgID int64, password string) error {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		u.TelegramID = tgID
		u.Password = password
		db.userToCache(u)
		db.lg.Info("Cache user updated")
	} else if !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	err = db.usersRepo.Update(ctx, tgID, password)
	if err != nil {
		return errors.Wrap(err, "Got error with updating user in cache")
	}
	return nil
}

func (db Database) Delete(ctx context.Context, tgID int64) error {
	err := db.cache.Delete(fmt.Sprintf("%v", tgID))
	if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Delete from cache fail: %v", err)
	} else {
		db.lg.Info("User unchached")
	}

	err = db.usersRepo.Delete(ctx, tgID)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting user from cahce")
	}
	return nil
}

func (db Database) AddStatement(ctx context.Context, tgID int64, statement *solv.Statement) error {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		*u.History = append(*u.History, *statement)
		db.userToCache(u)
		db.lg.Info("Cache statement updated")
	} else if !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	err = db.usersRepo.AddStatement(ctx, tgID, statement)
	if err != nil {
		return errors.Wrap(err, "Got error with adding statement to cahce")
	}
	return nil
}

func (db Database) AddStatements(ctx context.Context, tgID int64, statements *[]solv.Statement) error {
	return errors.New("Method forbiden.")
}

func (db Database) GetHistory(ctx context.Context, tgID int64) (*History, error) {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		db.lg.Info("From cache")
		return u.History, nil
	} else if !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	h, err := db.usersRepo.GetHistory(ctx, tgID)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with getting user history")
	}
	return h, nil
}

func (db Database) Clear(ctx context.Context, tgID int64) error {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		u.History = &History{}
		u.Variables = VMap{}
		db.userToCache(u)
		db.lg.Info("Cache user history cleared")
	} else if !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	err = db.usersRepo.Clear(ctx, tgID)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting user history from cache")
	}
	return nil
}

func (db Database) Exist(ctx context.Context, tgID int64) (bool, error) {
	ok, err := db.usersRepo.Exist(ctx, tgID)
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

func (db Database) GetVariable(ctx context.Context, tgID int64, name string) (float64, error) {
	cv, err := db.getVariableFromCache(tgID, name)
	if err == nil {
		db.lg.Info("From cache")
		return cv, nil
	} else if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v\n", err)
	}

	v, err := db.varsRepo.Get(ctx, tgID, name)
	if err != nil {
		return 0.0, errors.Wrap(err, "Got error with getting variable from cache")
	}
	return v, nil
}

func (db Database) GetAllVariables(ctx context.Context, tgID int64) (VMap, error) {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		db.lg.Info("From cache")
		return u.Variables, nil
	} else if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	vs, err := db.varsRepo.GetAll(ctx, tgID)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with getting all variables from cache")
	}
	return vs, nil
}

func (db Database) GetVariablesWithNames(ctx context.Context, tgID int64, names []string) (VMap, error) {
	res := VMap{}

	u, err := db.getUserFromCache(tgID)
	if err == nil {
		for _, n := range names {
			if v, ok := u.Variables[n]; ok {
				res[n] = v
			}
		}
		if len(res) == len(names) {
			db.lg.Info("From cache")
			return u.Variables, nil
		}
	} else if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	vs, err := db.varsRepo.GetWithNames(ctx, tgID, names)
	if err != nil {
		return nil, errors.Wrap(err, "Got error with getting multiple variables from cache")
	}
	return vs, nil
}

func (db Database) UpdateVariable(ctx context.Context, tgID int64, name string, value float64) error {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		u.Variables[name] = value
		db.userToCache(u)
		db.lg.Info("Cache variable updated")
	} else if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	err = db.varsRepo.Update(ctx, tgID, name, value)
	if err != nil {
		return errors.Wrap(err, "Got error with updating variable in cache")
	}
	return nil
}

func (db Database) UpdateVariablesWithNames(ctx context.Context, tgID int64, names []string, value float64) error {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		for _, n := range names {
			u.Variables[n] = value
		}
		db.userToCache(u)
		db.lg.Info("Cache variables updated")
	} else if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	err = db.varsRepo.UpdateWithNames(ctx, tgID, names, value)
	if err != nil {
		return errors.Wrap(err, "Got error with updating multiple variables in cache")
	}
	return nil
}

func (db Database) DeleteVariable(ctx context.Context, tgID int64, name string) error {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		delete(u.Variables, name)
		db.userToCache(u)
		db.lg.Info("Cache variable deleted")
	} else if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	err = db.varsRepo.Delete(ctx, tgID, name)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting variable in cache")
	}
	return nil
}

func (db Database) DeleteVariablesWithNames(ctx context.Context, tgID int64, names []string) error {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		for _, n := range names {
			delete(u.Variables, n)
		}
		db.userToCache(u)
		db.lg.Info("Cache variables deleted")
	} else if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	err = db.varsRepo.DeleteWithNames(ctx, tgID, names)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting multiple variables from cache")
	}
	return nil
}

func (db Database) DeleteAllVariables(ctx context.Context, tgID int64) error {
	u, err := db.getUserFromCache(tgID)
	if err == nil {
		u.Variables = VMap{}
		db.userToCache(u)
		db.lg.Info("Cache variables deleted")
	} else if err != nil && !errors.As(err, &itemNotFoundErr) {
		db.lg.Errorf("Cache fetch fail: %v", err)
	}

	err = db.varsRepo.DeleteAll(ctx, tgID)
	if err != nil {
		return errors.Wrap(err, "Got error with deleting all variables from cache")
	}
	return nil
}

func (db Database) WithinTransaction(ctx context.Context, f func(ctx context.Context) error) error {
	c, ctx := db.conn.WithConnection(ctx)
	c.Bind()
	defer c.Release()

	err := f(ctx)
	if err != nil {
		return errors.Wrap(err, `Error during transaction`)
	}
	return nil
}

func step(e error, opName string, res interface{}) bool {
	if e != nil {
		fmt.Println(e)
		return false
	} else {
		fmt.Println(opName, " SUCCESS")
	}
	fmt.Println(res)
	fmt.Println("next?")
	o, err := fmt.Scanln()
	if err != nil {
		fmt.Println(err)
		return false
	} else if o != 0 {
		return false
	}
	fmt.Println(o)
	return true
}

func users(ctx context.Context, db *Database) error {
	db.lg.Info("Urepo")

	s := &solv.Statement{
		Variables: []string{`r`, `g`},
		Equation:  "777",
		Value:     2.2,
	}

	us, err := db.GetAll(ctx)
	if !step(err, "GetAll", us) {
		return err
	}

	err = db.Add(ctx, 33333, "33333")
	if !step(err, "Add", "ok") {
		return err
	}

	u, err := db.Get(ctx, 11111)
	if !step(err, "Get", u) {
		return err
	}

	err = db.Update(ctx, 33333, "newpass")
	if !step(err, "Update", "ok") {
		return err
	}

	err = db.Delete(ctx, 33333)
	if !step(err, "Delete", `ok`) {
		return err
	}

	err = db.AddStatement(ctx, 11111, s)
	if !step(err, "AddStatement", `ok`) {
		return err
	}

	h, err := db.GetHistory(ctx, 11111)
	if !step(err, "GetHistory", h) {
		return err
	}

	ok, err := db.Exist(ctx, 11111)
	if !step(err, "Exist", ok) {
		return err
	}

	err = db.Clear(ctx, 11111)
	if !step(err, "Clear", `ok`) {
		return err
	}

	return nil
}

func vars(ctx context.Context, db *Database) error {
	_, err := db.GetAll(ctx)
	if err != nil {
		return err
	}

	db.lg.Info("Vrepo")

	v, err := db.GetVariable(ctx, 11111, `a`)
	if !step(err, "Get", v) {
		return err
	}

	vs, err := db.GetVariablesWithNames(ctx, 11111, []string{`a`, `c`})
	if !step(err, "GetWithNames", vs) {
		return err
	}

	vs, err = db.GetAllVariables(ctx, 11111)
	if !step(err, "GetAll", vs) {
		return err
	}

	err = db.UpdateVariable(ctx, 11111, `b`, 321.0)
	if !step(err, "Update", "ok") {
		return err
	}

	err = db.UpdateVariablesWithNames(ctx, 11111, []string{`c`, `a`}, 123.4)
	if !step(err, "UpdateWithNames", "ok") {
		return err
	}

	err = db.DeleteVariable(ctx, 11111, `c`)
	if !step(err, "Delete", "ok") {
		return err
	}
	err = db.DeleteVariablesWithNames(ctx, 11111, []string{`a`})
	if !step(err, "DeleteWithNames", "ok") {
		return err
	}
	err = db.DeleteAllVariables(ctx, 11111)
	if !step(err, "DeleteAll", "ok") {
		return err
	}

	return nil
}

func trans(ctx context.Context, db *Database) error {
	return db.WithinTransaction(ctx, func(ctx context.Context) error {
		log.Info("WithinTransaction")
		return users(ctx, db)
	})
}

const (
	DUMMY_VAR_RUN = iota
	DUMMY_USER_RUN
	DUMMY_TRANS_RUN
)

func DummyRun(ctx context.Context, db *Database, choice int) error {
	switch choice {
	case DUMMY_USER_RUN:
		err := users(ctx, db)
		if err != nil {
			return err
		}
	case DUMMY_VAR_RUN:
		err := vars(ctx, db)
		if err != nil {
			return err
		}
	case DUMMY_TRANS_RUN:
		err := trans(ctx, db)
		if err != nil {
			return err
		}
	}

	return nil
}
