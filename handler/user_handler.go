package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nku-treehole-server/dto"
	"nku-treehole-server/model"
	"nku-treehole-server/pkg/id_generator"
	"nku-treehole-server/pkg/logger"
	"nku-treehole-server/pkg/md5"
	"nku-treehole-server/service"
	"strings"
)

func Login(c *gin.Context) {
	var reqParam dto.LoginRequest
	if err := c.ShouldBindJSON(&reqParam); err != nil {
		ErrorResponse(c, "参数错误")
		return
	}
	u := &model.User{}
	user, err := u.SearchUserByEmail(reqParam.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse(c, "用户不存在")
		} else {
			ErrorResponse(c, fmt.Sprintf("用户查询失败 %v", err))
		}
		return
	}
	if user.Password != md5.Md5(reqParam.Password) {
		ErrorResponse(c, "密码错误")
		return
	}
	userService := &service.UserService{}
	token, err := userService.AddSession(user)
	if err != nil {
		ErrorResponse(c, "登陆失败")
		return
	}
	resp := dto.NewUserProfile(user, token)
	SuccessResponse(c, resp)
}

func Register(c *gin.Context) {
	var reqParam dto.RegisterRequest
	if err := c.ShouldBindJSON(&reqParam); err != nil {
		ErrorResponse(c, "参数错误")
		return
	}
	u := &model.User{}
	u.Email = reqParam.Email
	u.Password = md5.Md5(reqParam.Password)
	u.Name = reqParam.Username
	u.ID = id_generator.GenerateID()
	if err := u.CreateUser(u); err != nil {
		logger.Errorf("err = %v ", err.Error())
		if strings.Contains(err.Error(), "Duplicate") {
			ErrorResponse(c, "该邮箱已存在")
			return
		}
		ErrorResponse(c, "用户创建失败")
		return
	}

	userService := &service.UserService{}
	token, err := userService.AddSession(u)
	if err != nil {
		ErrorResponse(c, "注册失败")
		return
	}
	resp := dto.NewUserProfile(u, token)
	SuccessResponse(c, resp)
}
