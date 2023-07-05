package models

import (
	"mime/multipart"
	"time"

	"github.com/lib/pq"
)

type SignUpBody struct {
	Name     string                `form:"name" binding:"required"`
	Email    string                `form:"email" binding:"required"`
	Password string                `form:"password" binding:"required"`
	Image    *multipart.FileHeader `form:"image"`
	ImageURL string
}

type User struct {
	Id              string    `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Email           string    `json:"email" db:"email"`
	Password        string    `json:"password" db:"password"`
	ProfileImageURL *string   `json:"profileImageURL,omitempty" db:"profile_image_url"` // nullable
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

type LoginResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"accessToken"`
}

type UserResponse struct {
	Id              string    `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Email           string    `json:"email" db:"email"`
	ProfileImageURL *string   `json:"profileImageURL,omitempty" db:"profile_image_url"` // nullable
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

type UserUpdateBody struct {
	Name     string                `form:"name"`
	Email    string                `form:"email"`
	Password string                `form:"password"`
	Image    *multipart.FileHeader `form:"image"` // used only on controller level to upload the file
	ImageURL string                // used on service level to save URL in db
}

type TodoCreateBody struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	ActiveDays  []string `json:"activeDays"`
	Priority    string   `json:"priority"`
	IsDaily     bool     `json:"isDaily"`
}

type TodoResponseBody struct {
	Id          int            `db:"id" json:"id"`
	Title       string         `db:"title" json:"title"`
	Description string         `db:"description" json:"description"`
	Status      string         `db:"status" json:"status"`
	ActiveDays  pq.StringArray `db:"active_days" json:"activeDays"`
	Priority    string         `db:"priority" json:"priority"`
	IsDaily     bool           `db:"is_daily" json:"isDaily"`
	UserId      string         `db:"user_id" json:"userId"`
	CreatedAt   time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updatedAt"`
}

type TodoUpdateBody struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	ActiveDays  []string `json:"activeDays"`
	Priority    string   `json:"priority"`
	IsDaily     bool     `json:"isDaily"`
}
