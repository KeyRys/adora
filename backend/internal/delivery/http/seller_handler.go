package http

import (
	"backend/internal/domain"
	"backend/internal/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	Usecase *usecase.SellerUsecase
}

func NewSellerHandler(u *usecase.SellerUsecase) *SellerHandler {
	return &SellerHandler{
		Usecase: u,
	}
}

type BecomeSellerRequest struct {
	Location string `json:"location"`
}

func (h *SellerHandler) BecomeSeller(c *gin.Context) {

	var req BecomeSellerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("id").(string)

	err := h.Usecase.BecomeSeller(userID, req.Location)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success become seller",
	})
}

func (h *SellerHandler) CreateRabbit(c *gin.Context) {

	var rabbit domain.Rabbit
	//fmt.Println("Received request to create rabbit with body:", c.Request.Body)
	if err := c.ShouldBindJSON(&rabbit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("id").(string)
	err := h.Usecase.CreateRabbit(userID, &rabbit)
	//fmt.Println("CreateRabbit handler called with userID:", userID, "and rabbit:", rabbit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "rabbit created",
	})
}

func (h *SellerHandler) GetSellerRabbits(c *gin.Context) {

	userID := c.MustGet("id").(string)

	rabbits, err := h.Usecase.GetSellerRabbits(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rabbits)
}

func (h *SellerHandler) UpdateRabbit(c *gin.Context) {

	rabbitID := c.Param("id")

	seller_id := c.MustGet("id").(string)

	var body struct {
		Name         string `json:"name"`
		Breed        string `json:"breed"`
		HealthStatus string `json:"health_status"`
		Price        int    `json:"price"`
		Description  string `json:"description"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.Usecase.UpdateRabbit(
		rabbitID,
		seller_id,
		body.Name,
		body.Breed,
		body.HealthStatus,
		body.Price,
		body.Description,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "rabbit updated",
	})
}

func (h *SellerHandler) DeleteRabbit(c *gin.Context) {

	rabbitID := c.Param("id")

	user_id := c.MustGet("id").(string)

	err := h.Usecase.DeleteRabbit(
		rabbitID,
		user_id,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "rabbit deleted",
	})
}
