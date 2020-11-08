package mysql

import (
	"context"
	"go-gorm-echo/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

//NewUserRepository will create a new UserRepository object represent domain.UserRepository interface
func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &userRepository{Conn}
}

func (u *userRepository) FetchAll(ctx context.Context) (res []domain.User, err error) {
	var users []domain.User

	u.Conn.Find(&users)

	return users, nil
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	var user domain.User

	u.Conn.First(&user, id)

	return user, nil
}

func (u *userRepository) Store(ctx context.Context, data *domain.User) (err error) {
	u.Conn.Create(&data)

	return nil
}

func (u *userRepository) Update(ctx context.Context, data *domain.User, id int64) (res domain.User, err error) {
	var user domain.User

	u.Conn.First(&user, id).Updates(&data)

	return user, nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) (err error) {
	var user domain.User

	u.Conn.Delete(&user, id)

	return nil
}

func (u *userRepository) Login(ctx context.Context, data *domain.User) (res domain.User, err error) {
	var user domain.User

	u.Conn.Where("email = ?", data.Email).First(&user)

	return user, err
}
