package main

import (
	"fmt"
	"go-gorm-echo/domain"

	_userHttpDelivery "go-gorm-echo/user/delivery/http"
	_userRepository "go-gorm-echo/user/repository/mysql"
	_userUseCase "go-gorm-echo/user/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "echo"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	errMigration := db.AutoMigrate(&domain.User{})

	if errMigration != nil {
		log.Error(errMigration)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Error(err)
		}
		sqlDB.Close()
	}()

	userRepo := _userRepository.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUseCase(userRepo)

	e := echo.New()
	_userHttpDelivery.NewUserHandler(e, userUsecase)

	e.Start(":3000")
}
