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

func (h *campaignHandler)DetailCampaign(c *gin.Context){
	//input dari url
	//handler
	//service
	//repo
	//api/v1/campaign/1 -> id campaign
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"error":"ID Campaign Not Found"}
		response := helper.APIResponse("Campaign Failed",http.StatusBadRequest,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	data,err := h.service.GetCampaignById(input)
	if err != nil {
		errorMessage := gin.H{"error":"Campaign Not Found"}
		response := helper.APIResponse("Campaign Failed",http.StatusBadRequest,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	response := helper.APIResponse("Campaign  Detail success",http.StatusOK,"success",data)
	c.JSON(http.StatusOK,response)



}