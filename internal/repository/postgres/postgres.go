package postgres

import (
	"effective/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to PostgreSQL.
func Connect(path string, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(path), cfg)
}

// Migrate PostgreSQL to HEAD.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Human{})
}
