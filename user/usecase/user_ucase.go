package usecase

import (
	"context"
	"go-gorm-echo/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

// NewUserUseCase will create new User Object representation of domain.UserUseCase
func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepository: userRepo,
	}
}

func (u *userUseCase) FetchAll(ctx context.Context) (res []domain.User, err error) {
	res, err = u.userRepository.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (u *userUseCase) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	res, err = u.userRepository.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return
}

func (u *userUseCase) Store(c context.Context, data *domain.User) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 12)
	data.Password = string(hashedPassword)
	err = u.userRepository.Store(c, data)
	return
}

func (u *userUseCase) Update(c context.Context, data *domain.User, id int64) (res domain.User, err error) {
	res, err = u.userRepository.Update(c, data, id)
	if err != nil {
		return domain.User{}, err
	}
	return res, err
}

func (u *userUseCase) Delete(c context.Context, id int64) (err error) {
	err = u.userRepository.Delete(c, id)

	return
}

func (u *userUseCase) Login(c context.Context, data *domain.User) (token string, err error) {
	res, err := u.userRepository.Login(c, data)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(data.Password))

	token, err = u.createToken(data)

	return token, err

}

func (u *userUseCase) createToken(data *domain.User) (token string, err error) {
	secretToken := "insertyourjwthere"
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user"] = data
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(secretToken))

	return

}
