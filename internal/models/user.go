package models

import "context"

type User struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phoneNumber"`
	Information string `json:"information"`
	RoleID      uint   `json:"roleId"`
	CreatedAt   string `json:"createdAt"`
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
	Password    Password `json:"password"`
	PhoneNumber string   `json:"phoneNumber"`
	Information string   `json:"information"`
	Name        string   `json:"name"`
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
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, userID int) (User, error)
	GetUserProfile(c context.Context, userID int) (User, error)
	GetVolunteerProfile(c context.Context, userID int) (User, error)
	GetCodeByEmail(c context.Context, email string) (string, error)
	CreateUser(c context.Context, user UserRequest) (int, error)
	CreateVolunteer(c context.Context, user VolunteerRequest) (int, error)
	ChangeForgottenPassword(c context.Context, code string, email string, newPassword string) error
	CreatePasswordResetCode(c context.Context, email string, code string) error
	GetVolunteerByEmail(c context.Context, email string) (User, error)
	ChangeForgottenVolunteersPassword(c context.Context, code string, email string, newPassword string) error
}
