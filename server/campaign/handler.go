package campaign

import (
	"fmt"
	"landingpage/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService Service
}

func NewHandler(campaignService Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) FindCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	fmt.Println(userId)
	campaigns, err := h.campaignService.FindCampaigns(userId)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Error Find Campaign!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := FormatCampaigns(campaigns)

	response := helper.APIResponse("Find Campaigns", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
