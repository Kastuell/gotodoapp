package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kastuell/gotodoapp/internal/models"
)

func (h *Handler) register(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Auth.Register(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

}
