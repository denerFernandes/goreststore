package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User - Models the user
type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// HashPassword with bcrypt
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword with encrypted password
func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeSave - Before Start Saving User we encrypt password to store
func (user *User) BeforeSave() error {
	hashedPassword, error := HashPassword(user.Password)
	if error != nil {
		return error
	}
	user.Password = string(hashedPassword)
	return nil
}

// Prepare - Prepare
func (user *User) Prepare() {
	user.ID = 0
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

// Validate - Validate data
func (user *User) Validate(action string) error {

	// This function used for
	// Validate request body from
	// Android Developer or WEB DEV

	switch strings.ToLower(action) {
	case "update":
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if user.Name == "" {
			return errors.New("Required Full Name")
		}
		if error := checkmail.ValidateFormat(user.Email); error != nil {
			return errors.New("Invalid Email format")
		}

		return nil

	case "login":
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Number Phome")
		}
		return nil

	default:
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if error := checkmail.ValidateFormat(user.Email); error != nil {
			return errors.New("Invalid Email format")
		}

		return nil
	}
}

// SaveUser - Create User
func (user *User) SaveUser(db *gorm.DB) (*User, error) {

	// Inserting to Database
	var err error
	err = db.Debug().Model(&User{}).Create(&user).Error

	// Check Error
	if err != nil {
		return &User{}, err
	}

	// If there is no error
	return user, err

}

// FindAllUsers - Find All User
func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {

	// Retrieving users from Database
	var err error
	var users []User
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error

	// Check Error
	if err != nil {
		return &[]User{}, err
	}

	// If there is no error
	return &users, nil

}

// FindUserByID - Find User By ID
func (user *User) FindUserByID(db *gorm.DB, userID uint64) (*User, error) {

	// Retrieving position by User(Driver) ID from Database
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", userID).Take(&user).Error

	// Check Error
	if err != nil {
		return &User{}, err
	}

	// Check If Position record not found
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}

	// If there is no error
	return user, nil
}

// UpdateUser - Update user
func (user *User) UpdateUser(db *gorm.DB, userID uint64) (*User, error) {

	// Hashing password before update
	err := user.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	// Updating User(Driver) in Database
	db = db.Debug().Model(&User{}).Where("id = ?", userID).Take(&User{}).Updates(&User{
		Password:  user.Password,
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: time.Now(),
	})

	// Check user updating Error
	if db.Error != nil {
		return &User{}, db.Error
	}

	// Retrieve Updated User
	err = db.Debug().Model(&User{}).Where("id = ?", userID).Take(&user).Error

	// Check Position Retrieval Error
	if err != nil {
		return &User{}, err
	}

	// If there is no error
	return user, nil

}

// DeleteUser - Delete A User
func (user *User) DeleteUser(db *gorm.DB, userID uint64) (int64, error) {

	// Delete driver position
	db = db.Debug().Model(&User{}).Where("id = ?", userID).Take(&User{}).Delete(&User{})

	// Check User Position removal Error
	if db.Error != nil {
		return 0, db.Error
	}

	// If there is no error
	return db.RowsAffected, nil

}

// FindUserByEmail - find user by email
func (user *User) FindUserByEmail(db *gorm.DB, Email string) (*User, error) {

	var err error

	err = db.Debug().Model(User{}).Where("email = ?", Email).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}
