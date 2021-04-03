package handler

import (
	"github.com/gin-gonic/gin"
	"golangbwa/campaign"
	"golangbwa/helper"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler{
	return &campaignHandler{service}
}

func (h *campaignHandler)GetCampaign(c *gin.Context){
	//strconv.Atoi digunakan untuk convert string into integer
	userID,_ := strconv.Atoi(c.Query("user_id"))
	campaigns,err := h.service.GetCampaigns(userID)

	if err != nil{
		errorMessage := gin.H{"error":"Data Campaign Not Found"}
		response := helper.APIResponse("Campaign Failed",http.StatusBadRequest,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	//if len(campaigns) == 0{
	//	errorMessage := gin.H{"error":"Data Campaign Not Found"}
	//	response := helper.APIResponse("Campaign Failed",http.StatusBadRequest,"error",errorMessage)
	//	c.JSON(http.StatusBadRequest,response)
	//	return
	//}

	response := helper.APIResponse("Campaign  List success",http.StatusOK,"success",campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK,response)
}