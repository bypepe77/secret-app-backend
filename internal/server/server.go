package server

import (
	"fmt"
	"net/http"

	"github.com/bypepe77/secret-app-backend/internal/modules/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	config *Config
	db     *gorm.DB
	engine *gin.Engine
}

func NewServer(config *Config, db *gorm.DB) *server {
	return &server{
		config: config,
		db:     db,
		engine: gin.New(),
	}
}

func (s *server) buildConnectionString() string {
	if s.config.Port == "" {
		s.config.Port = "8080"
	}

	return fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
}

func (s *server) Run() error {
	s.engine.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Authorization", "content-type"},
		AllowHeaders:     []string{"Authorization", "content-type "},
	}))
	s.registerRoutes()
	return s.engine.Run()
}

func (s *server) registerRoutes() {
	s.engine.GET("/", healthCheck)

	// Register auth routes
	authRoutes := auth.NewUserRoute(s.db, *s.engine.Group("/auth"))
	authRoutes.RegisterUserRoutes()
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong", "status": true})
}
