package models

import (
	"time"

	"go-admin/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"unique;not null;size:50"`
	Password  string         `json:"-" gorm:"not null"` // json:"-" indicates not displayed in JSON
	Email     *string        `json:"email" gorm:"unique;size:100"`
	Roles     []Role         `json:"roles" gorm:"many2many:user_roles;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Migrate automatically migrates table structure
func Migrate() error {
	// First migrate Role table
	if err := database.DB.AutoMigrate(&Role{}); err != nil {
		return err
	}

	// Initialize default roles
	if err := InitRoles(); err != nil {
		return err
	}

	// Then migrate User table (GORM will auto-create user_roles join table)
	return database.DB.AutoMigrate(&User{})
}

// HashPassword encrypts the password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies the password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// Create creates a user
func (u *User) Create() error {
	return database.DB.Create(u).Error
}

// AddRole adds a role to the user
func (u *User) AddRole(role *Role) error {
	return database.DB.Model(u).Association("Roles").Append(role)
}

// RemoveRole removes a role from the user
func (u *User) RemoveRole(role *Role) error {
	return database.DB.Model(u).Association("Roles").Delete(role)
}

// HasRole checks if the user has a specific role
func (u *User) HasRole(roleName string) bool {
	for _, role := range u.Roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}

// GetRoleNames returns a list of role names for the user
func (u *User) GetRoleNames() []string {
	names := make([]string, len(u.Roles))
	for i, role := range u.Roles {
		names[i] = role.Name
	}
	return names
}

// FindByUsername finds a user by username with roles preloaded
func FindByUsername(username string) (*User, error) {
	var user User
	err := database.DB.Preload("Roles").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID finds a user by ID with roles preloaded
func FindByID(id uint) (*User, error) {
	var user User
	err := database.DB.Preload("Roles").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
