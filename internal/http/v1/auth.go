package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kastuell/gotodoapp/internal/domain"
)

type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
		// auth.POST("/refresh", h.userRefresh)
	}
}

func (h *Handler) register(c *gin.Context) {
	var input domain.CreateUserInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.services.Auth.Register(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})

}

func (h *Handler) login(c *gin.Context) {
	var input domain.GetIdByCreditsInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.services.Auth.Login(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}
