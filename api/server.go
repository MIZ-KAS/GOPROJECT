package api

import (
	"github.com/gin-gonic/gin"
	db "kodboris/db/sqlc"
)

type Server struct {
	router *gin.Engine
	db     *db.Store
}

func (server *Server) setupRouter() {
	router := gin.Default()
	//router.POST("api/v1/accounts", server.createAccount)
	router.GET("/", server.homeProbe)
	server.router = router
	return

}

// start the HTTP server
func (server *Server) Start(address string, store *db.Store) error {
	server.setupRouter()
	return server.router.Run(address)
}

// error handler
func errorResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}
