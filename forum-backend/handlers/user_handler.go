package handlers

import (
	"net/http"
	"forum/dto"
	"forum/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
	jwtService *services.JWTService
}

func NewUserHandler(userService *services.UserService , jwtService *services.JWTService) *UserHandler {
	return &UserHandler{
		userService: userService,
		jwtService: jwtService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user":    user,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"user":    user,
		"token":   token,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}