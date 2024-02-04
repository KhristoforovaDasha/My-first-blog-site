package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserRegister
	IsAdmin bool
}

type UserRegister struct {
	Login    string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
}

type UserRepository interface {
	Create(*User) error
	GetAll() (*[]User, error)
	Get(id uint) (*User, error)
	Update(*User) error
	Delete(id uint) error
	GetByLogin(login string) (*User, error)
}

type UserService interface {
	Get(id uint) (*User, error)
	Update(user *User) error
	Delete(user *User) error

	Register(userReg *UserRegister) error
	Login(userLog *UserRegister) (uint, error)
}
