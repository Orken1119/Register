package models

import (
	"context"
	"time"
)

type User struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name,omitempty"`
	Password    string    `json:"password,omitempty"`
	PhoneNumber string    `json:"phoneNumber"`
	RoleID      uint      `json:"roleId"`
	CreatedAt   time.Time `json:"created_at"`
	Skills      string    `json:"skills,omitempty"`
	City        string    `json:"city,omitempty"`
	Age         int       `json:"age,omitempty"`
}

type ChangePasswordRequest struct {
	Email    string   `json:"email"`
	Code     string   `json:"code"`
	Password Password `json:"password"`
}

type UserRequest struct {
	Email       string   `json:"email"`
	Password    Password `json:"password"`
	PhoneNumber string   `json:"phoneNumber"`
}

type VolunteerRequest struct {
	Email       string   `json:"email"`
	Password    Password `json:"password,omitempty"`
	PhoneNumber string   `json:"phoneNumber"`
	Name        string   `json:"name"`
	Skills      string   `json:"skills"`
	City        string   `json:"city"`
	Age         int      `json:"age"`
}

type VolunteerPersonalData struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Skills      string `json:"skills"`
	City        string `json:"city"`
	Age         int    `json:"age"`
	Direction   string `json:"direction"`
}

type VolunteerProfile struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Skills      string `json:"skills"`
	City        string `json:"city"`
	Age         int    `json:"age"`
	Direction   string `json:"direction"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByID(c context.Context, userID int) (User, error)
	GetUserProfile(c context.Context, userID int) (User, error)
	GetVolunteerProfile(c context.Context, userID int) (VolunteerProfile, error)
	GetCodeByEmail(c context.Context, email string) (string, error)
	CreateUser(c context.Context, user UserRequest) (int, error)
	CreateVolunteer(c context.Context, user VolunteerRequest) (int, error)
	ChangeForgottenPassword(c context.Context, code string, email string, newPassword string) error
	CreatePasswordResetCode(c context.Context, email string, code string) error
	GetVolunteerByEmail(c context.Context, email string) (User, error)
	ChangeForgottenVolunteersPassword(c context.Context, code string, email string, newPassword string) error
	ChangePassword(c context.Context, userID int, password string) error
	EditPersonalData(c context.Context, userID int, volunteer VolunteerPersonalData) error
}

type RequestRepository interface {
	MakeRequest(c context.Context)
}
