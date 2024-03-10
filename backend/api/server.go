package api

import (
	"chatbot-ai/util"

	"github.com/jackc/pgx/v5"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our chatbot service.
type Server struct {
	config util.Config
	store  *pgx.Conn
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store *pgx.Conn) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	// Setup Router
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Use CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Update with your allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	// add routes to router
	router.GET("/chatbot/initialize_data", server.initializeData)
	router.POST("/chatbot/prompt", server.prompt)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
