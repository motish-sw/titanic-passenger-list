package core

import (
	"github.com/motish-sw/titanic-passenger-list/config"

	"github.com/jinzhu/gorm"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo   *echo.Echo            // HTTP middleware
	config *config.Configuration // Configuration
	db     *gorm.DB              // Database connection
}

// NewServer will create a new instance of the application
func NewServer(config *config.Configuration) *Server {
	server := &Server{}
	server.config = config

	server.Echo = NewRouter(server)

	// Middleware
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())

	return server
}

// GetDB returns gorm (ORM)
func (s *Server) GetDB() *gorm.DB {
	return s.db
}

// // GetCache returns the current redis client
// func (s *Server) GetCache() *redis.Client {
// 	return s.cache
// }

// // GetConfig return the current app configuration
// func (s *Server) GetConfig() *config.Configuration {
// 	return s.config
// }

// // GetModelRegistry returns the model registry
// func (s *Server) GetModelRegistry() *models.Model {
// 	return s.modelRegistry
// }

// Start the http server
func (s *Server) Start(addr string) error {
	return s.Echo.Start(addr)
}

// // ServeStaticFiles serve static files for development purpose
// func (s *Server) ServeStaticFiles() {
// 	s.Echo.Static("/assets", s.config.AssetsBuildDir)
// }

// // GracefulShutdown Wait for interrupt signal
// // to gracefully shutdown the server with a timeout of 5 seconds.
// func (s *Server) GracefulShutdown() {
// 	quit := make(chan os.Signal, 1)

// 	signal.Notify(quit, os.Interrupt)
// 	<-quit
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// // close cache
// 	// if s.cache != nil {
// 	// 	cErr := s.cache.Close()
// 	// 	if cErr != nil {
// 	// 		s.Echo.Logger.Fatal(cErr)
// 	// 	}
// 	// }

// 	// close database connection
// 	if s.db != nil {
// 		dErr := s.db.Close()
// 		if dErr != nil {
// 			s.Echo.Logger.Fatal(dErr)
// 		}
// 	}

// 	// shutdown http server
// 	if err := s.Echo.Shutdown(ctx); err != nil {
// 		s.Echo.Logger.Fatal(err)
// 	}
// }
