package models

import "time"

// UserRole represents the many-to-many relationship between users and roles
type UserRole struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	RoleID    uint      `json:"role_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName specifies the table name for UserRole
func (UserRole) TableName() string {
	return "user_roles"
}
