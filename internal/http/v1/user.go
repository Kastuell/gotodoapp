package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kastuell/gotodoapp/internal/domain"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user", h.userIdentity)
	{
		user.GET("", h.getMe)
	}
}

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

	c.JSON(http.StatusOK, map[string]domain.User{
		"user": user,
	})

}
