package middleware

import (
	"thuchanhgolang/pkg/encrypter"
	"thuchanhgolang/pkg/jwt"
	"thuchanhgolang/pkg/log"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Auth() gin.HandlerFunc
}

type implMiddleware struct {
	l         log.Logger
	jwtMgr    jwt.Manager
	encrypter encrypter.Encrypter
}

func New(l log.Logger, jwtMgr jwt.Manager, enc encrypter.Encrypter) Middleware {
	return &implMiddleware{
		l:         l,
		jwtMgr:    jwtMgr,
		encrypter: enc,
	}
}
