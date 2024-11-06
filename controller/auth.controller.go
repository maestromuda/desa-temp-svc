package controller

import (
	"desa-temp-svc/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username      string `json:"username"`
	DesaName      string `json:"desa_name"`
	KecamatanName string `json:"kecamatan_name"`
	KabupatenName string `json:"kabupaten_name"`
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body LoginRequest true "Login Request"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := usecase.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
