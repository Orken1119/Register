package ivent

import (
	"net/http"
	"strconv"

	models "github.com/Orken1119/Register/internal/models"
	"github.com/gin-gonic/gin"
)

// @Tags		ivent
// @Accept		json
// @Produce	json
// @Param id path int true "id"
// @Security Bearer
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/ivents/delete-ivent/{id} [delete]
func (av *IventController) DeleteIvent(c *gin.Context) {
	idVal := c.Param("id")
	id, err := strconv.Atoi(idVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_IVENT",
					Message: "incorrect id format",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	err = av.IventRepository.DeleteIvent(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "DELETE_IVENT_ERROR",
					Message: "Error to delete ivent",
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Ivent was deleted successfully"})

}
