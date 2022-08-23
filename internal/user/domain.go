package user

import (
	"context"
	"time"

	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"go.uber.org/zap"
)

type VMap map[string]float64

type History []solv.Statement

type User struct {

	// Internal DB id
	Id int `db:"id"`

	// Telegram user id
	TelegramID int64 `db:"tg_id"`

	// Randomly generated password
	Password string `db:"password"`

	// History of solving
	History *History

	// User personal variables
	Variables VMap
}

type Database struct {
	usersRepo UserRepository
	varsRepo  VariableRepository

	cache Cache

	lg *zap.SugaredLogger
}

//
//
//
//
//
type UserRepository interface {
	Add(ctx context.Context, tgID int64, password string) error

	Get(ctx context.Context, tgID int64) (*User, error)
	GetAll(ctx context.Context) (*[]User, error)

	Update(ctx context.Context, tgID int64, password string) error

	Delete(ctx context.Context, tgID int64) error

	AddStatement(ctx context.Context, tgID int64, statement *solv.Statement) error

	GetHistory(ctx context.Context, tgID int64) (*History, error)

	Exist(ctx context.Context, tgID int64) (bool, error)
	Clear(ctx context.Context, tgID int64) error
}

//
//
//
//
//
type VariableRepository interface {
	Get(ctx context.Context, tgID int64, name string) (float64, error)
	GetWithNames(ctx context.Context, tgID int64, names []string) (VMap, error)
	GetAll(ctx context.Context, tgID int64) (VMap, error)

	Update(ctx context.Context, tgID int64, name string, value float64) error
	UpdateWithNames(ctx context.Context, tgID int64, names []string, value float64) error

	Delete(ctx context.Context, tgID int64, name string) error
	DeleteWithNames(ctx context.Context, tgID int64, names []string) error
	DeleteAll(ctx context.Context, tgID int64) error
}

//
//
// Database Proxy
//
//

type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, error)
	Delete(key string) error
	Count() int
	Rename(prewKey, newKey string) error
	Exist(key string) bool
	FlushAll() int
}

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

type ItemNotFoundError struct{}

func (e ItemNotFoundError) Error() string {
	return "Item not found"
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
