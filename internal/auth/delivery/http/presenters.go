package http

import (
	"strings"
	"thuchanhgolang/internal/auth"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// registerReq là cấu trúc nhận dữ liệu đăng ký từ HTTP request
type registerReq struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	ShopID   string `json:"shop_id" binding:"required"` // Shop ID
}

// validate kiểm tra dữ liệu đầu vào
func (r registerReq) validate() error {
	// Kiểm tra các fields không rỗng
	if strings.TrimSpace(r.Username) == "" || strings.TrimSpace(r.Password) == "" || strings.TrimSpace(r.Email) == "" {
		return errWrongBody
	}

	// Kiểm tra ShopID hợp lệ
	if _, err := primitive.ObjectIDFromHex(r.ShopID); err != nil {
		return errWrongBody
	}

	return nil
}

// toInput chuyển đổi request thành input cho usecase
func (r registerReq) toInput() auth.RegisterInput {
	shopID, _ := primitive.ObjectIDFromHex(r.ShopID)
	return auth.RegisterInput{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
		ShopID:   shopID,
	}
}

// registerResp là cấu trúc response sau khi đăng ký thành công
type registerResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ShopID   string `json:"shop_id"`
	Token    string `json:"token"`
}

// newRegisterResp tạo response từ RegisterOutput
func (h handler) newRegisterResp(output auth.RegisterOutput) registerResp {
	return registerResp{
		ID:       output.ID.Hex(),
		Username: output.Username,
		Email:    output.Email,
		ShopID:   output.ShopID.Hex(),
		Token:    output.Token,
	}
}

// loginReq là cấu trúc nhận dữ liệu đăng nhập từ HTTP request
type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// validate kiểm tra dữ liệu đầu vào
func (r loginReq) validate() error {
	// Kiểm tra các fields không rỗng
	if strings.TrimSpace(r.Username) == "" || strings.TrimSpace(r.Password) == "" {
		return errWrongBody
	}
	return nil
}

// toInput chuyển đổi request thành input cho usecase
func (r loginReq) toInput() auth.LoginInput {
	return auth.LoginInput{
		Username: r.Username,
		Password: r.Password,
	}
}

// loginResp là cấu trúc response sau khi đăng nhập thành công
type loginResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

// newLoginResp tạo response từ LoginOutput
func (h handler) newLoginResp(output auth.LoginOutput) loginResp {
	return loginResp{
		ID:       output.ID.Hex(),
		Username: output.Username,
		Email:    output.Email,
		Token:    output.Token,
	}
}

// emptyScope trả về scope rỗng
func (h handler) emptyScope() models.Scope {
	return models.Scope{}
}
