package models

import (
	"gorm.io/gorm"
)

type Users struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap  string `json:"nama_lengkap" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

func (u *Users) CreateUser(db *gorm.DB) int64 {
	cek_user := db.Where("username = ?", u.Username).First(u)
	if cek_user.RowsAffected == 0 {
		return db.Create(u).RowsAffected
	}
	return 0

}

func FindUserByUsername(db *gorm.DB, username string, user *Users) *gorm.DB {
	return db.First(user, username)
}
