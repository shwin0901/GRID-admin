package models

import (
	"time"

	"go-admin/database"
)

// Role name constants
const (
	RoleNameAdmin    = "admin"
	RoleNameInternal = "internal"
)

// Role represents a user role
type Role struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique;not null;size:50"`
	Description string    `json:"description" gorm:"size:200"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// InitRoles creates default roles if they don't exist
func InitRoles() error {
	roles := []Role{
		{ID: 1, Name: RoleNameAdmin, Description: "Administrator with full access"},
		{ID: 2, Name: RoleNameInternal, Description: "Internal user with limited access"},
	}

	for _, role := range roles {
		// Use FirstOrCreate to avoid duplicate entries
		if err := database.DB.FirstOrCreate(&role, Role{Name: role.Name}).Error; err != nil {
			return err
		}
	}
	return nil
}

// FindRoleByName finds a role by its name
func FindRoleByName(name string) (*Role, error) {
	var role Role
	err := database.DB.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// FindRoleByID finds a role by its ID
func FindRoleByID(id uint) (*Role, error) {
	var role Role
	err := database.DB.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
