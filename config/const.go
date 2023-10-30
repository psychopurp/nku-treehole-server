package config

import "time"

const (
	UID string = "UID"
	// Login expiration duration: three days
	EXPIRE_DURATION time.Duration = 3 * 24 * 60 * time.Minute

	// Success status code
	SUCCESS_CODE int = 200

	// Login expired or not logged in
	LOGIN_EXPIRE int = 400
)
