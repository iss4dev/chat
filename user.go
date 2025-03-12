package todo

import "time"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password_hash"`
	Role     string `json:"role" db:"role"`
}

type UserProfile struct {
	AvatarURL string    `json:"avatar_url" db:"avatar_url"`
	Bio       string    `json:"bio" db:"bio"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
