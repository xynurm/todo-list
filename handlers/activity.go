package handlers

import (
	"net/http"
	"strconv"
	"test/dto"
	activitydto "test/dto/activity"
	"test/models"
	"test/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerActivity struct {
	ActivityRepository repositories.ActivityRepository
}

func HandlerActivity(ActivityRepository repositories.ActivityRepository) *handlerActivity {
	return &handlerActivity{ActivityRepository}
}

// get  ALLusers
// @Summary CRUD  GET ALLUSER
// @Produce json
// @Router /users [get]
func (h *handlerActivity) GetAllActivity(c *gin.Context) {
	activities, err := h.ActivityRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: activities}
	c.JSON(http.StatusOK, response)
}

func (h *handlerActivity) CreateActivity(c *gin.Context) {
	var request activitydto.ActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validator
	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		response := dto.ErrorResult{Status: "Bad request", Message: "Title cannot be null"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	activityData := models.Activity{
		Title: request.Title,
		Email: request.Email,
	}

	activity, err := h.ActivityRepository.Save(activityData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: activity}
	c.JSON(http.StatusCreated, response)
}

// getuser by id
// @Summary get a user item by ID
// @ID get-user-by-id
// @Produce json
// @Param id path string true "user ID"
// @Router /user/{id} [get]
func (h *handlerActivity) GetActivityID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	activity, err := h.ActivityRepository.FindByID(id)
	if err != nil {
		response := dto.ErrorResult{Status: "Not found", Message: "Activity with ID " + strconv.Itoa(id) + " not found"}
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: activity}
	c.JSON(http.StatusOK, response)
}

// @Summary delete a user by ID
// @ID delete-user-by-id
// @Produce json
// @Param id path string true "user ID"
// @Router /user/{id} [delete]
func (h *handlerActivity) DeleteActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	activity, err := h.ActivityRepository.FindByID(id)
	if err != nil {
		response := dto.ErrorResult{Status: "Not found", Message: "Activity with ID " + strconv.Itoa(id) + " not found"}
		c.JSON(http.StatusNotFound, response)
		return
	}

	data, _ := h.ActivityRepository.Delete(activity)

	deleteResponse := activitydto.DeleteResponse{
		ActivityID: data.ActivityID,
	}

	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: deleteResponse}
	c.JSON(http.StatusOK, response)
}

func (h *handlerActivity) UpdateActivity(c *gin.Context) {
	var request activitydto.UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validator
	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		response := dto.ErrorResult{Status: "Bad request", Message: "Title cannot be null"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	activity, err := h.ActivityRepository.FindByID(id)
	if err != nil {
		response := dto.ErrorResult{Status: "Not found", Message: "Activity with ID " + strconv.Itoa(id) + " not found"}
		c.JSON(http.StatusNotFound, response)
		return
	}

	activity.Title = request.Title

	activityUpdate, err := h.ActivityRepository.Update(activity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: activityUpdate}
	c.JSON(http.StatusOK, response)
}
