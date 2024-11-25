package user

import (
	"net/http"

	models "github.com/Orken1119/Register/internal/models"
	"github.com/gin-gonic/gin"
)

// @Tags user
// @Accept json
// @Produce json
// @Param request body models.VolunteerPersonalData true "query params"
// @Security Bearer
// @Success 200 {object} map[string]string
// @Failure default {object} models.ErrorResponse
// @Router /user/edit-profile [put]
func (sc *UserController) EditPersonalData(c *gin.Context) {
	userID := c.GetUint("userID")
	var request models.VolunteerPersonalData

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of changePassword",
				},
			},
		})
		return
	}

	err = sc.UserRepository.EditPersonalData(c, int(userID), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_USER_PROFILE",
					Message: "Can't get profile from db",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Personal data was changed"})

}
