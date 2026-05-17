package http

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Usecase *usecase.ProfileUsecase
}

func NewProfileHandler(
	u *usecase.ProfileUsecase,
) *ProfileHandler {

	return &ProfileHandler{
		Usecase: u,
	}
}

func (h *ProfileHandler) GetMyProfile(
	c *gin.Context,
) {

	userID := c.GetString("id")
	//fmt.Println("userID:", userID)

	profile, err := h.Usecase.GetProfile(userID)
	//fmt.Println("GET PROFILE ERROR:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}
