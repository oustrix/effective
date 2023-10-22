package v1

import (
	"net/http"
	"strconv"

	"effective/internal/entity"
	"effective/internal/usecase"

	"github.com/gin-gonic/gin"
)

type humansRoutes struct {
	uc *usecase.HumansUseCase
}

func setupHumansRoutes(handler *gin.RouterGroup, uc *usecase.HumansUseCase) {
	r := humansRoutes{
		uc: uc,
	}

	h := handler.Group("/humans")
	{
		h.GET("/", r.getHumans)
		h.POST("/", r.createHuman)

		h.DELETE("/:id", r.deleteHuman)
		h.PUT("/:id", r.updateHuman)
	}
}

// @Summary		Return humans
// @Description	Return all humans by given filters
// @ID				getHumans
// @Tags			humans
// @Accept			json
// @Produce		json
// @Param			gender		query		string	false	"filter by gender"
// @Param			ageMin		query		int		false	"filter by minimal age"
// @Param			ageMax		query		int		false	"filter by maximum age"
// @Param			nation		query		string	false	"filter by nation"
// @Param			page		query		int		false	"filter by page (default=1)"
// @Param			pageSize	query		int		false	"filter by page size (default=10)"
// @Success		200			{object}	entity.HumansList
// @Failure		400			{object}	response
// @Failure		500			{object}	response
// @Router			/humans/ [get]
// .
func (r *humansRoutes) getHumans(c *gin.Context) {
	var filter entity.HumanFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	humans, err := r.uc.Get(&filter)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, humans)
	}
}

// @Summary		Create human
// @Description	Create human with given info
// @ID				createHuman
// @Tags			humans
// @Accept			json
// @Produce		json
// @Param			json	body		entity.CreateHuman	true	"human data"
// @Success		200		{object}	entity.Human
// @Failure		400		{object}	response
// @Failure		500		{object}	response
// @Router			/humans/ [post]
// .
func (r *humansRoutes) createHuman(c *gin.Context) {
	var humanData entity.CreateHuman

	err := c.Bind(&humanData)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	human, err := r.uc.Create(&humanData)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, human)
	}
}

// @Summary		Delete human
// @Description	Delete human by give ID
// @ID				deleteHuman
// @Tags			humans
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Human ID"
// @Success		200
// @Failure		400	{object}	response
// @Failure		500	{object}	response
// @Router			/humans/{id} [delete]
// .
func (r *humansRoutes) deleteHuman(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = r.uc.Delete(id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary		Update human
// @Description	Update human by give ID and updated data
// @ID				updateHuman
// @Tags			humans
// @Accept			json
// @Produce		json
// @Param			json	body		entity.UpdateHuman	true	"new human data. no one field is required"
// @Param			id	path	int	true	"Human ID"
// @Success		200 {object} entity.Human
// @Failure		400	{object}	response
// @Failure		500	{object}	response
// @Router			/humans/{id} [put]
// .
func (r *humansRoutes) updateHuman(c *gin.Context) {
	var humanData entity.UpdateHuman

	err := c.Bind(&humanData)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	human, err := r.uc.Update(id, &humanData)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, human)
	}
}
