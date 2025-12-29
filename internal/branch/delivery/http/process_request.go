package http

import (
	"thuchanhgolang/internal/models"
	pkgErrors "thuchanhgolang/pkg/errors"
	"thuchanhgolang/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func (h handler) processCreateRequest(c *gin.Context) (createReq, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "branch.http.processCreateRequest.GetPayloadFromContext: unauthorized")
		return createReq{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "branch.http.processCreateRequest.ShouldBindJSON: %v", err)
		return createReq{}, models.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "branch.http.processCreateRequest.validate: %v", err)
		return createReq{}, models.Scope{}, err
	}

	sc := jwt.NewScope(payload)

	return req, sc, nil
}
