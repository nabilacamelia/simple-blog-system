package payload

import (
	"simple-blog-system/internal/app/user/model"
)

type User struct {
	User model.AuthUserModel `json:"auth_user"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"` // n kecil semua!
	Password string `json:"password" binding:"required"` // n kecil semua!
}
