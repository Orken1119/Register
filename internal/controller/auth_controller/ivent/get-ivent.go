package ivent

import (
	"net/http"

	models "github.com/Orken1119/Register/internal/models"
	"github.com/gin-gonic/gin"
)

// @Tags		ivent
// @Accept		json
// @Produce	json
// @Security Bearer
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/ivents/get-ivents [get]
func (av *IventController) GetAllIvent(c *gin.Context) {
	ivents, err := av.IventRepository.GetAllIvent(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "IVENT_ERROR",
					Message: "Error to get all ivents",
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: ivents})
}
