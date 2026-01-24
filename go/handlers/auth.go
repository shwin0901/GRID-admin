package handlers

import (
	"go-admin/middleware"
	"go-admin/models"
	"go-admin/utils"

	"github.com/gin-gonic/gin"
)

// RegisterRequest represents a registration request
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"omitempty,email"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	Token    string       `json:"token"`
	UserInfo *models.User `json:"user_info"`
}

// Register handles user registration
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid parameters: "+err.Error())
		return
	}

	// Check if the username already exists
	existingUser, _ := models.FindByUsername(req.Username)
	if existingUser != nil {
		utils.BadRequest(c, "Username already exists")
		return
	}

	// Find internal role
	role, err := models.FindRoleByName(models.RoleNameInternal)
	if err != nil {
		utils.InternalError(c, "Failed to find role")
		return
	}

	// Handle empty email (store as NULL instead of empty string)
	var email *string
	if req.Email != "" {
		email = &req.Email
	}

	// Create a user with initial role
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    email,
		Roles:    []models.Role{*role},
	}

	// Encrypt the password
	if err := user.HashPassword(); err != nil {
		utils.InternalError(c, "Failed to encrypt password")
		return
	}

	// Save to the database
	if err := user.Create(); err != nil {
		utils.InternalError(c, "Failed to create user")
		return
	}

	utils.SuccessWithMessage(c, "Registration successful", user)
}

// Login handles user login
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid parameters: "+err.Error())
		return
	}

	// Find the user
	user, err := models.FindByUsername(req.Username)
	if err != nil {
		utils.BadRequest(c, "Invalid username or password")
		return
	}

	// Verify the password
	if !user.CheckPassword(req.Password) {
		utils.BadRequest(c, "Invalid username or password")
		return
	}

	// Generate a token with all user roles
	token, err := middleware.GenerateToken(user.ID, user.Username, user.GetRoleNames())
	if err != nil {
		utils.InternalError(c, "Failed to generate token")
		return
	}

	utils.Success(c, LoginResponse{
		Token:    token,
		UserInfo: user,
	})
}

// GetProfile retrieves user information
func GetProfile(c *gin.Context) {
	// Get the user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		utils.Unauthorized(c, "Please log in first")
		return
	}

	// Find the user
	user, err := models.FindByID(userID.(uint))
	if err != nil {
		utils.NotFound(c, "User not found")
		return
	}

	utils.Success(c, user)
}
