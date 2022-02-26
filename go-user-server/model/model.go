package model

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
}

/*
type ModelWithOperator struct {
	gorm.Model
	ID UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	CreatorID UUID `gorm:"not null"`
	Creator   User
	UpdaterID UUID `gorm:"not null"`
	Updater   User
	DeleterID NullUUID
	Deleter   *User
}
*/

type User struct {
	Model

	Username string `gorm:"not null;uniqueIndex"`
	Password []byte `gorm:"not null;default:''"`
	FullName string `gorm:"not null;index"`
	IDNumber string `gorm:"not null;uniqueIndex"`
	Email    string `gorm:"not null"`
	Role     string `gorm:"not null"`
}
