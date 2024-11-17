package auth

import (
	"fmt"
	"net/http"
	"unicode"

	"github.com/Orken1119/Register/internal/controller/auth_controller/tokenutil"
	models "github.com/Orken1119/Register/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserRepository models.UserRepository
}

// @Accept json
// @Produce json
// @Param request body models.UserRequest true "query params"
// @Success 200 {object} models.SuccessResponse
// @Failure default {object} models.ErrorResponse
// @Router /signup [post]
func (uc AuthController) Signup(c *gin.Context) {
	var request models.UserRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of signup",
				},
			},
		})
		return
	}

	user, _ := uc.UserRepository.GetUserByEmail(c, request.Email)
	if user.ID > 0 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "USER_EXISTS",
					Message: "User with this email already exists",
				},
			},
		})
		return
	}

	err := ValidatePassword(request.Password.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_PASSWORD_FORMAT",
					Message: err.Error(),
				},
			},
		})
		return
	}
	// Подтверждение пароля
	err = ConfirmPassword(request.Password.Password, request.Password.ConfirmPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_PASSWORD_MISMATCH",
					Message: err.Error(),
				},
			},
		})
		return
	}
	//

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_ENCRYPTE_PASSWORD",
					Message: "Couldn't encrypte password",
				},
			},
		})
		return
	}
	request.Password.Password = string(encryptedPassword)
	request.Password.ConfirmPassword = ""

	_, err = uc.UserRepository.CreateUser(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_USERS",
					Message: "Couldn't create user",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	user, err = uc.UserRepository.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_USER",
					Message: "User with this email wasn't found",
				},
			},
		})
		return
	}

	accessToken, err := tokenutil.CreateAccessToken(user, `access-secret-key`, 50)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "TOKEN_ERROR",
					Message: "Error to create access token",
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: accessToken})

}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	var (
		hasUpper, hasLower, hasDigit bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit {
		return fmt.Errorf("password must contain at least one uppercase letter, one lowercase letter and one digit")
	}

	return nil
}

// Функция подтверждения
func ConfirmPassword(password string, confirm string) error {
	if password != confirm {
		return fmt.Errorf("doesn't match the password")
	}
	return nil
}
