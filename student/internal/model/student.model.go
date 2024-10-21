package model

import (
	"time"
)

type Student struct {
	ID             int        `gorm:"primaryKey;autoIncrement"`
	FullName       string     `gorm:"type:varchar(255);not null;default:''"`
	Sex            int        `gorm:"type:smallint;not null;default:0"`
	DOB            time.Time  `gorm:"type:date;not null"`
	UniversityID   int        `gorm:"not null"`
	EnrollmentYear int        `gorm:"not null"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
	University     University `gorm:"foreignKey:UniversityID"`
}
