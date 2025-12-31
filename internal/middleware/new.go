package middleware

import (
	"thuchanhgolang/internal/models"
	"thuchanhgolang/pkg/encrypter"
	"thuchanhgolang/pkg/jwt"
	"thuchanhgolang/pkg/log"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Auth() gin.HandlerFunc
	RequireRole(allowedRoles ...models.Role) gin.HandlerFunc
	CheckShopAccess() gin.HandlerFunc
	CheckRegionAccess() gin.HandlerFunc
	CheckBranchAccess() gin.HandlerFunc
	CheckDepartmentAccess() gin.HandlerFunc
	CheckUserAccess() gin.HandlerFunc
	SetScopeFromPayload() gin.HandlerFunc
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
