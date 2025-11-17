package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kastuell/gotodoapp/internal/models"
)

func (h *Handler) getMe(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.User.GetMe(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]models.User{
		"user": user,
	})

}
