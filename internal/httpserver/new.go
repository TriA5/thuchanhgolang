package httpserver

import (
	"time"

	pkgLog "thuchanhgolang/pkg/log"
	"thuchanhgolang/pkg/mongo"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	gin            *gin.Engine
	l              pkgLog.Logger
	port           int
	database       mongo.Database
	jwtSecretKey   string
	accessDuration time.Duration
	// encrypter    pkgCrt.Encrypter
	// secretConfig SecretConfig
}

type Config struct {
	Port           int
	Database       mongo.Database
	JWTSecretKey   string
	AccessDuration time.Duration
	// Encrypter    pkgCrt.Encrypter
	// SecretConfig SecretConfig
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	return &HTTPServer{
		l:              l,
		gin:            gin.Default(),
		port:           cfg.Port,
		database:       cfg.Database,
		jwtSecretKey:   cfg.JWTSecretKey,
		accessDuration: cfg.AccessDuration,
		// encrypter:    cfg.Encrypter,
		// secretConfig: cfg.SecretConfig,
	}
}
