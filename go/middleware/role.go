package middleware

import (
	"go-admin/models"
	"go-admin/utils"

	"github.com/gin-gonic/gin"
)

// RequireRole creates a middleware that checks if the user has at least one of the required roles
func RequireRole(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get roles from context (set by JWTAuth middleware)
		rolesValue, exists := c.Get("roles")
		if !exists {
			utils.Unauthorized(c, "Role information not found")
			c.Abort()
			return
		}

		userRoles := rolesValue.([]string)

		// Check if user has at least one of the required roles
		for _, userRole := range userRoles {
			for _, requiredRole := range requiredRoles {
				if userRole == requiredRole {
					c.Next()
					return
				}
			}
		}

		utils.Forbidden(c, "Insufficient permissions")
		c.Abort()
	}
}

// RequireAdmin is a convenience middleware that only allows admin users
func RequireAdmin() gin.HandlerFunc {
	return RequireRole(models.RoleNameAdmin)
}

// RequireAllRoles creates a middleware that checks if the user has ALL of the required roles
func RequireAllRoles(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get roles from context (set by JWTAuth middleware)
		rolesValue, exists := c.Get("roles")
		if !exists {
			utils.Unauthorized(c, "Role information not found")
			c.Abort()
			return
		}

		userRoles := rolesValue.([]string)
		userRoleMap := make(map[string]bool)
		for _, role := range userRoles {
			userRoleMap[role] = true
		}

		// Check if user has ALL required roles
		for _, requiredRole := range requiredRoles {
			if !userRoleMap[requiredRole] {
				utils.Forbidden(c, "Insufficient permissions")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
