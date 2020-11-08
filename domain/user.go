package domain

import (
	"context"
	"time"
)

//User Model
type User struct {
	ID              int       `json:"id"`
	Name            string    `gorm:"type:varchar(30)" json:"name"`
	Email           string    `gorm:"type:varchar(90)" json:"email"`
	Password        string    `gorm:"type:varchar(255)" json:"password"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

//UserUseCase represent the logic and the business flow
type UserUseCase interface {
	FetchAll(ctx context.Context) (res []User, err error)
	Store(ctx context.Context, data *User) (err error)
	GetByID(ctx context.Context, id int64) (res User, err error)
	Update(ctx context.Context, data *User, id int64) (res User, err error)
	Delete(ctx context.Context, id int64) (err error)
	Login(ctx context.Context, data *User) (token string, err error)
}

//UserRepository represent the sql query
type UserRepository interface {
	FetchAll(ctx context.Context) (res []User, err error)
	Store(ctx context.Context, data *User) (err error)
	GetByID(ctx context.Context, id int64) (res User, err error)
	Update(ctx context.Context, data *User, id int64) (res User, err error)
	Delete(ctx context.Context, id int64) (err error)
	Login(ctx context.Context, data *User) (res User, err error)
}
