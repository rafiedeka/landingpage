package user

import (
	"landingpage/auth"
	"landingpage/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Service
	authService auth.Service
}

func NewHandler(userService Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input RegisterUserDto

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Account failed to register!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Account failed to register!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenereteToken(user.ID)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Account failed to register!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	newUser, err := h.userService.SaveToken(user.ID, token)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Account failed to register!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input LoginUserDto

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	loggedinUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenereteToken(loggedinUser.ID)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Account failed to register!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	newUser, err := h.userService.SaveToken(loggedinUser.ID, token)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := FormatUser(newUser, token)

	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input CheckEmailDto

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	isEmailAvailable, err := h.userService.isEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}

		response := helper.APIResponse("Email Checking Failed", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"

	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUserId := c.MustGet("currentUser").(User)

	path := "images/" + helper.RandomString(10) + "_" + file.Filename

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, err := h.userService.GetUserByID(currentUserId.ID)

	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID.ID, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
