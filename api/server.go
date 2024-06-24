package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vivekcode101/simplebanks-golang/db/sqlc"
)

// Server serves all HTTP request for banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// New Server creates a new HTTP server and routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start Runs the http server on specific server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}