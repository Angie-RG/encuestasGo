package app

import (
	"AngelicaRG/encuestasGo/models"
	"AngelicaRG/encuestasGo/utils"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db     *gorm.DB
	onceDB sync.Once
)

type UserRole struct {
	User models.User
	Role models.Role
}

func DB() *gorm.DB {
	onceDB.Do(func() {
		port, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			port = 5432
		}

		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			port,
			os.Getenv("DB_NAME"),
		)

		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err != nil {
			log.Fatal("Could'nt connection data base")
		}

		if err := database.AutoMigrate(
			&models.Role{},
			&models.User{},
			&models.UserRole{},
			&models.Pregunta{},
			&models.Respuesta{},
			&models.PreguntaRespuesta{},
		); err != nil {
			log.Fatal("Could'nt migrate models")
		}

		db = database
	})

	return db
}

func Seeders() {
	configRoles()
	setupUsers()
}

func configRoles() {
	roles := []models.Role{
		{Name: "superAdmin"},
		{Name: "admin"},
		{Name: "user"},
		{Name: "viewer"},
	}

	for _, rol := range roles {
		role := &models.Role{}
		if err := DB().Where(&models.Role{Name: rol.Name}).FirstOrCreate(&role).Error; err != nil {
			slog.Error(fmt.Sprintf("Could'nt create role: %v", rol.Name))
			continue
		}
	}
}

func setupUsers() {
	active := true
	pass := os.Getenv("GENERATE_PASSWORD")

	userRoles := []UserRole{
		{
			User: models.User{
				FirstName: utils.ToString("Melinda"),
				LastName:  utils.ToString("Gordon"),
				Email:     "melinda.gordon@outlook.com",
				Active:    &active,
				Password:  pass,
			},
			Role: models.Role{
				Name: "admin",
			},
		},
		{
			User: models.User{
				FirstName: utils.ToString("Sebastian Alan"),
				LastName:  utils.ToString("Rodriguez Hernandez"),
				Email:     "sebastian.hernandez@outlook.com",
				Active:    &active,
				Password:  pass,
			},
			Role: models.Role{
				Name: "superAdmin",
			},
		},
		{
			User: models.User{
				FirstName: utils.ToString("Caroline Diana"),
				LastName:  utils.ToString("Lopez"),
				Email:     "caroline.lopez@outlook",
				Active:    &active,
				Password:  pass,
			},
			Role: models.Role{
				Name: "user",
			},
		},
	}

	if err := DB().Transaction(func(tx *gorm.DB) error {
		for _, data := range userRoles {
			user := data.User
			if err := tx.Where(&models.User{Email: user.Email}).FirstOrCreate(&user).Error; err != nil {
				slog.Error(fmt.Sprintf("The user could not be created: %v", err))
				return errors.New("Error occurred while creating the user.")
			}

			role := data.Role
			if err := tx.Where(&role).First(&role).Error; err != nil {
				slog.Error(fmt.Sprintf("Could not be found the rol: %v", err))
				return errors.New("Could not be found rol")
			}

			userRole := &models.UserRole{UserID: user.ID, RoleID: role.ID, CreatedByID: user.ID, UpdatedByID: user.ID}
			if err := tx.Where(&userRole).FirstOrCreate(&userRole).Error; err != nil {
				slog.Error(fmt.Sprintf("Could not assign role of user: %v", err))
				return errors.New("Could not assign role of user")
			}
		}

		return nil
	}); err != nil {
		slog.Error(fmt.Sprintf("Error could not create roles of user: %v", err))
	}

}
