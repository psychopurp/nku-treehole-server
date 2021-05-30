package config

import "time"

const (
	UID string = "UID"
	// 登陆过期时效 三天
	EXPIRE_DURATION time.Duration = 3 * 24 * 60 * time.Minute

	// 请求成功状态码
	SUCCESS_CODE int = 200

	// 登录过期，或者未登录
	LOGIN_EXPIRE int = 400
)
