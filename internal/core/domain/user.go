package domain

import "time"

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
)

type User struct {
	AuditFields
	Email        string     `json:"username" db:"email"`
	FullName     string     `json:"full_name" db:"full_name"`
	PasswordHash string     `json:"-" db:"password_hash"`
	RoleID       string     `json:"role_id" db:"role_id"`
	Status       UserStatus `json:"status" db:"status"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
}
