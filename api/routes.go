package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type probe struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (server *Server) homeProbe(ctx *gin.Context) {
	// Set status code to 200
	response := probe{
		Status:  "200",
		Message: "Welcome to Kodeboris Community",
	}
	ctx.JSON(http.StatusOK, response)

}
