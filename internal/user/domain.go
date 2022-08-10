package user

import (
	"context"

	solv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

type VMap map[string]float64

type User struct {

	// Telegram user id
	ID int64

	// Randomly generated password
	Password string

	// History of solving
	History *[]solv.Statement

	// User personal variables
	Variables VMap
}

// Our use-case or service will implement these methods.
type UserService interface {
	FetchUsers(ctx context.Context) (*[]User, error)
	FetchUser(ctx context.Context, userID int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	SetUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, userID int64) error

	AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error
	FetchHistory(ctx context.Context, userID int64) (*[]solv.Statement, error)
	ClearHistory(ctx context.Context, userID int64) error
}

type UserRepository interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, userID int64) (*User, error)
	AddUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, userID int64) error

	GetHistory(ctx context.Context, userID int64) (*[]solv.Statement, error)
	AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error
	DeleteHistory(ctx context.Context, userID int64) error
}

type VariableService interface {
	AddUserVariable(ctx context.Context, userID int64, name string, value float64) error
	AddUserVariables(ctx context.Context, userID int64, names []string, value float64) error

	FetchUserVariable(ctx context.Context, userID int64, name string) (float64, error)
	FetchUserVariables(ctx context.Context, userID int64, names []string) (VMap, error)
	FetchAllUserVariables(ctx context.Context, userID int64) (VMap, error)

	SetUserVariable(ctx context.Context, userID int64, name string, value float64) error
	SetUserVariables(ctx context.Context, userID int64, names []string, value float64) error

	DeleteUserVariable(ctx context.Context, userID int64, name string) error
	ClearUserVariables(ctx context.Context, userID int64) error
}

type VariableRepository interface {
	CreateUserVariable(ctx context.Context, userID int64, name string, value float64) error
	CreateUserVariables(ctx context.Context, userID int64, names []string, values float64) error

	GetUserVariable(ctx context.Context, userID int64, name string) (float64, error)
	GetUserVariables(ctx context.Context, userID int64, names []string) (VMap, error)
	GetAllUserVariables(ctx context.Context, userID int64) (VMap, error)

	UpdateUserVariable(ctx context.Context, userID int64, name string, value float64) error
	UpdateUserVariables(ctx context.Context, userID int64, names []string, values []float64) error

	DeleteUserVariable(ctx context.Context, userID int64, name string) error
	DeleteUserVariables(ctx context.Context, userID int64) error
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
