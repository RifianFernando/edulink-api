package models

import (
	"time"
	"gorm.io/gorm"
)

type Student struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	Name              string         `json:"name" binding:"required"`
	Gender            string         `json:"gender" binding:"required,oneof=Male Female other"`
	PlaceOfBirth      string         `json:"place_of_birth" binding:"required"`
	DateOfBirth       time.Time      `json:"date_of_birth" binding:"required,datetime=2006-01-02"`
	Religion          string         `json:"religion" binding:"required"`
	Address           string         `json:"address" binding:"required"`
	NumberPhone       string         `json:"number_phone" binding:"required,e164"`
	Email             string         `json:"email" binding:"required,email"`
	AcceptedDate      time.Time      `json:"accepted_date" binding:"required,datetime=2006-01-02"`
	SchoolOrigin      string         `json:"school_origin" binding:"required"`
	IDClass           uint           `json:"id_class" binding:"required"`
	FatherName        string         `json:"father_name" binding:"required"`
	FatherJob         string         `json:"father_job" binding:"required"`
	FatherNumberPhone string         `json:"father_number_phone" binding:"required,e164"`
	MotherName        string         `json:"mother_name" binding:"required"`
	MotherJob         string         `json:"mother_job" binding:"required"`
	MotherNumberPhone string         `json:"mother_number_phone" binding:"required,e164"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
