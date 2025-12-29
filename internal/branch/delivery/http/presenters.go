package http

import (
	"strings"

	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/models"
	"thuchanhgolang/pkg/response"
)

type createReq struct {
	Name string `json:"name" binding:"required"`
}

func (r createReq) validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

func (r createReq) toInput() branch.CreateInput {
	return branch.CreateInput{
		Name: r.Name,
	}
}

type detailResp struct {
	ID        string            `json:"id"`
	Alias     string            `json:"alias"`
	Code      string            `json:"code"`
	CreatedAt response.DateTime `json:"created_at"`
}

func (h handler) newDetailResp(d models.Branch) detailResp {
	return detailResp{
		ID:        d.ID.Hex(),
		Alias:     d.Alias,
		Code:      d.Code,
		CreatedAt: response.DateTime(d.CreatedAt),
	}
}
