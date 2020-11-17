package models

import (
	//"errors"

	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	fmt.Println("Insert boi Nhi %s", err)
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) SelectSaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Select("Username", "Email", "Password").Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BatchInsert(db *gorm.DB) (*User, error) {
	var err error
	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// func (u *User) PostUser(db *gorm.DB) (*User, error) {
// 	var err error
// 	err = db.Create(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }

// Get user
func (u *User) SingleObject(db *gorm.DB) (*User, error) {
	var err error
	err = db.First(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// func (u *User) FindByID(db *gorm.DB, uid uint64) (*User, error) {
// 	var err error
// 	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }

// func (u *User) FindAllUser(db *gorm.DB) (*[]User, error) {
// 	var err error
// 	users := []User{}
// 	err = db.Debug().Model(&User{}).Limit(100).Find(&u).Error
// 	if err != nil {
// 		return &[]User{}, err
// 	}
// 	return &users, nil
// }

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Updates(User{Username: u.Username, Email: u.Email, Password: u.Password}).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
