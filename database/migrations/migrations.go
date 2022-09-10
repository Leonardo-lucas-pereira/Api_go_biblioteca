package migrations

import (
	"github.com/Leonardo-lucas-pereira/Api_go_biblioteca/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
