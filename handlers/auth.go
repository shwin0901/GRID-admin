package handlers

import (
	"go-admin/middleware"
	"go-admin/models"
	"go-admin/utils"

	"github.com/gin-gonic/gin"
)

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"omitempty,email"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token    string       `json:"token"`
	UserInfo *models.User `json:"user_info"`
}

// Register 用户注册
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 检查用户名是否已存在
	existingUser, _ := models.FindByUsername(req.Username)
	if existingUser != nil {
		utils.BadRequest(c, "用户名已存在")
		return
	}

	// 创建用户
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	// 加密密码
	if err := user.HashPassword(); err != nil {
		utils.InternalError(c, "密码加密失败")
		return
	}

	// 保存到数据库
	if err := user.Create(); err != nil {
		utils.InternalError(c, "创建用户失败")
		return
	}

	utils.SuccessWithMessage(c, "注册成功", user)
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 查找用户
	user, err := models.FindByUsername(req.Username)
	if err != nil {
		utils.Unauthorized(c, "用户名或密码错误")
		return
	}

	// 验证密码
	if !user.CheckPassword(req.Password) {
		utils.Unauthorized(c, "用户名或密码错误")
		return
	}

	// 生成 Token
	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.InternalError(c, "生成 Token 失败")
		return
	}

	utils.Success(c, LoginResponse{
		Token:    token,
		UserInfo: user,
	})
}

// GetProfile 获取用户信息
func GetProfile(c *gin.Context) {
	// 从上下文获取用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		utils.Unauthorized(c, "请先登录")
		return
	}

	// 查找用户
	user, err := models.FindByID(userID.(uint))
	if err != nil {
		utils.NotFound(c, "用户不存在")
		return
	}

	utils.Success(c, user)
}
