package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                uint    `gorm:"primaryKey;auto_increment" json:"id"`
	Name              string  `gorm:"size:255;not null" json:"name"`
	Email             string  `gorm:"size:100;not null;unique;unique_index" json:"email"`
	Password          string  `gorm:"size:100" json:"-"`
	Verified          bool    `gorm:"type:boolean" json:"verified"`
	VerificationToken *string `gorm:"size:50;" json:"-"`
	Image             *string `gorm:"size:500;" json:"image"`
	Admin             bool    `gorm:"type:boolean;not null;default:0" json:"admin"`

	// Time
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

/**
 * TableName
 *
 * @return string
 */
func (User) TableName() string {
	return "users"
}

/**
 * VerifyPassword
 *
 * @param string , string
 * @return error
 */
func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

/**
 * IsVerified
 *
 * @return bool
 */
func (u *User) IsVerified() bool {
	return u.Verified
}

/**
 * IsAdmin
 *
 * @return bool
 */
func (u *User) IsAdmin() bool {
	return u.Admin
}
