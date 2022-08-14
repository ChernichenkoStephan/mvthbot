package user

import (
	"context"
	"sync"
	"time"

	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

type VMap map[string]float64

type History []solv.Statement

type User struct {

	// Telegram user id
	ID int64

	// Randomly generated password
	Password string

	// History of solving
	History *History

	// User personal variables
	Variables VMap
}

//
//
//
//
//
type UserService interface {
	Add(ctx context.Context, user *User) error

	Get(ctx context.Context, userID int64) (*User, error)
	GetAll(ctx context.Context) (*[]User, error)

	Update(ctx context.Context, user *User) error

	Delete(ctx context.Context, userID int64) error

	AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error

	GetHistory(ctx context.Context, userID int64) (*History, error)

	DeleteHistory(ctx context.Context, userID int64) error

	Exist(ctx context.Context, userID int64) (bool, error)
	Clear(ctx context.Context, userID int64) error
}

//
//
//
//
//
type UserRepository interface {
	Add(ctx context.Context, user *User) error

	Get(ctx context.Context, userID int64) (*User, error)
	GetAll(ctx context.Context) (*[]User, error)

	Update(ctx context.Context, user *User) error

	Delete(ctx context.Context, userID int64) error

	AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error

	GetHistory(ctx context.Context, userID int64) (*History, error)

	DeleteHistory(ctx context.Context, userID int64) error

	Exist(ctx context.Context, userID int64) bool
	Clear(ctx context.Context, userID int64) error
}

//
//
//
//
//
type VariableService interface {
	Add(ctx context.Context, userID int64, name string, value float64) error
	AddWithNames(ctx context.Context, userID int64, names []string, value float64) error

	Get(ctx context.Context, userID int64, name string) (float64, error)
	GetWithNames(ctx context.Context, userID int64, names []string) (VMap, error)
	GetAll(ctx context.Context, userID int64) (VMap, error)

	Update(ctx context.Context, userID int64, name string, value float64) error
	UpdateWithNames(ctx context.Context, userID int64, names []string, value float64) error

	Delete(ctx context.Context, userID int64, name string) error
	DeleteWithNames(ctx context.Context, userID int64, names []string) error
	DeleteAll(ctx context.Context, userID int64) error
}

//
//
//
//
//
type VariableRepository interface {
	Add(ctx context.Context, userID int64, name string, value float64) error
	AddWithNames(ctx context.Context, userID int64, names []string, value float64) error

	Get(ctx context.Context, userID int64, name string) (float64, error)
	GetWithNames(ctx context.Context, userID int64, names []string) (VMap, error)
	GetAll(ctx context.Context, userID int64) (VMap, error)

	Update(ctx context.Context, userID int64, name string, value float64) error
	UpdateWithNames(ctx context.Context, userID int64, names []string, values []float64) error

	Delete(ctx context.Context, userID int64, name string) error
	DeleteWithNames(ctx context.Context, userID int64, names []string) error
	DeleteAll(ctx context.Context, userID int64) error
}

//
//
// Database Proxy
//
//

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item

	vr VariableRepository
	ur UserRepository
}

//
//
//  Data types for views
//
//

type StatementDTO struct {
	names    []string `validate:"required"`
	equation string   `validate:"required"`
}

type VariablesPackDTO struct {
	statements []StatementDTO `validate:"required"`
}
