package user

import (
	"context"
	"sync"
	"time"

	"database/sql"

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

type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	Get(dest interface{}, query string, args ...interface{}) error
	Rebind(query string) string
}

type connKey struct{}

type Connection struct {
	mx       sync.Mutex
	refCount uint

	DB
}

type Connector interface {
	WithConnection(ctx context.Context) (*Connection, context.Context)
}

type Database struct {
	usersRepo UserRepository
	varsRepo  VariableRepository

	cache Cache

	conn Connector

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
//  DTO's for views
//
//

// swagger:model Statement
type StatementDTO struct {
	// Variables names in statement to set
	// Example: a=b=2+2, names=["a","b"]
	// Could be empty
	// in: string[]
	Names []string `json:"names"`
	// Equation to solve
	// Example: a=b=2+2, equation="2+2"
	// in: string
	Equation string `json:"equation"`
}

// swagger:model Pack
type PackDTO struct {
	// Request for muliple solving
	// in: Statement[]
	Statements []StatementDTO `json:"statements"`
}

// type variable struct {
// 	Name  string
// 	Value float64
// }

//
//
// DTO's for DB requests
//
//

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
