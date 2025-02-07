package app

import (
	"AngelicaRG/encuestasGo/models"
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
