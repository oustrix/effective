package entity

// Human is main human structure.
type Human struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"size:32,not null"`
	Surname    string `gorm:"size:32,not null"`
	Patronymic string `gorm:"size:32"`
	Age        int    `gorm:"not null"`
	Gender     string `gorm:"size:32,not null"`
	Nation     string `gorm:"size:32,not null"`
}

// HumansList -.
type HumansList []Human

// CreateHuman is a DTO for create human request.
type CreateHuman struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic,omitempty"`
}

// UpdateHuman is a DTO for update human request.
type UpdateHuman struct {
	Name       *string `json:"name,omitempty"`
	Surname    *string `json:"surname,omitempty"`
	Patronymic *string `json:"patronymic,omitempty"`
}

// HumanFilter contains all filter which could be used for filtering GET-request.
type HumanFilter struct {
	Gender   string `form:"gender"`
	AgeMin   int    `form:"ageMin"`
	AgeMax   int    `form:"ageMax"`
	Nation   string `form:"nation"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}
