package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"nku-treehole-server/dto"
	"nku-treehole-server/model"
	"nku-treehole-server/pkg/id_generator"
	"nku-treehole-server/pkg/logger"
	"nku-treehole-server/pkg/md5"
	"nku-treehole-server/service"
)

func Login(c *gin.Context) {
	var reqParam dto.LoginRequest
	if err := c.ShouldBindJSON(&reqParam); err != nil {
		ErrorResponse(c, "Parameter error")
		return
	}
	u := &model.User{}
	user, err := u.SearchUserByEmail(reqParam.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse(c, "User does not exist")
		} else {
			ErrorResponse(c, fmt.Sprintf("User query failed: %v", err))
		}
		return
	}
	if user.Password != md5.Md5(reqParam.Password) {
		ErrorResponse(c, "Incorrect password")
		return
	}
	userService := &service.UserService{}
	token, err := userService.AddSession(user)
	if err != nil {
		ErrorResponse(c, "Login failed")
		return
	}
	resp := dto.NewUserProfile(user, token)
	SuccessResponse(c, resp)
}

func Register(c *gin.Context) {
	var reqParam dto.RegisterRequest
	if err := c.ShouldBindJSON(&reqParam); err != nil {
		ErrorResponse(c, "Parameter error")
		return
	}
	u := &model.User{}
	u.Email = reqParam.Email
	u.Password = md5.Md5(reqParam.Password)
	u.Name = reqParam.Username
	u.ID = id_generator.GenerateID()
	if err := u.CreateUser(u); err != nil {
		logger.Errorf("Error = %v", err.Error())
		if strings.Contains(err.Error(), "Duplicate") {
			ErrorResponse(c, "Email already exists")
			return
		}
		ErrorResponse(c, "User creation failed")
		return
	}

	userService := &service.UserService{}
	token, err := userService.AddSession(u)
	if err != nil {
		ErrorResponse(c, "Registration failed")
		return
	}
	resp := dto.NewUserProfile(u, token)
	SuccessResponse(c, resp)
}
