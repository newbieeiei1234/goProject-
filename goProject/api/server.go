package api

import (
	db "goProject/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.PATCH("/accounts/update", server.partialUpdateAccount)
	router.DELETE("/accounts/delete", server.deleteAccount)
	router.POST("/accounts/transfer", server.createTransfer)

	server.router = router
	return server
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func messageResponse(m string) gin.H {
	return gin.H{"message": m}
}
