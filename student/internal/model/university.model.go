package model

import "time"

type University struct {
	ID                int       `gorm:"primaryKey;autoIncrement"`
	Name              string    `gorm:"type:varchar(255);not null;default:''"`
	EstablishmentYear int       `gorm:"not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
	Students          []Student `gorm:"foreignKey:UniversityID"`
}
