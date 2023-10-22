package postgres

import (
	"effective/internal/entity"

	"gorm.io/gorm"
)

// HumansRepository contains a DB.
type HumansRepository struct {
	db *gorm.DB
}

// NewHumansRepository creates new HumansRepository.
func NewHumansRepository(db *gorm.DB) *HumansRepository {
	return &HumansRepository{
		db: db,
	}
}

// CreateHuman creates new human.
func (r *HumansRepository) CreateHuman(human *entity.Human) (*entity.Human, error) {
	err := r.db.Create(human).Error

	return human, err
}

// DeleteHuman deletes existing human.
func (r *HumansRepository) DeleteHuman(id int) error {
	return r.db.Delete(&entity.Human{}, id).Error
}

// GetHumans return humans by given filters.
func (r *HumansRepository) GetHumans(filter *entity.HumanFilter) (*entity.HumansList, error) {
	var humans entity.HumansList

	query := r.db.Model(&entity.Human{})

	if filter.Gender != "" {
		query = query.Where("gender = ?", filter.Gender)
	}
	if filter.AgeMin > 0 {
		query = query.Where("age >= ?", filter.AgeMin)
	}
	if filter.AgeMax > 0 {
		query = query.Where("age <= ?", filter.AgeMax)
	}
	if filter.Nation != "" {
		query = query.Where("nation = ?", filter.Nation)
	}

	offset := (filter.Page - 1) * filter.PageSize

	err := query.Offset(offset).Limit(filter.PageSize).Find(&humans).Error
	if err != nil {
		return nil, err
	}

	return &humans, nil
}

// GetHuman return human by given id.
func (r *HumansRepository) GetHuman(id int) (*entity.Human, error) {
	var human entity.Human
	err := r.db.First(&human, id).Error

	return &human, err
}

// UpdateHuman updates existing human.
func (r *HumansRepository) UpdateHuman(human *entity.Human) error {
	return r.db.Save(human).Error
}
