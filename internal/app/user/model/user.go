package model

import (
	"time"
	"github.com/go-openapi/strfmt"
)

type AuthUserModel struct {
	ID        strfmt.UUID4 `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username  string       `json:"username" gorm:"column:username"` // <--- WAJIB "username" (kecil semua)
	Password  string       `json:"password" gorm:"column:password"` // <--- WAJIB "password" (kecil semua)
	IsActive  bool         `json:"is_active" gorm:"column:is_active;default:true"`
	LastLogin time.Time    `json:"last_login" gorm:"column:last_login"`
	CreatedBy string       `json:"created_by" gorm:"column:created_by"`
	UpdatedBy string       `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt *time.Time   `json:"deleted_at" gorm:"column:deleted_at"`
}

func (u AuthUserModel) TableName() string {
	return "users"
}