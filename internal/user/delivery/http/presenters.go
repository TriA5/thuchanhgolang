package http

import (
	"strings"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createReq là cấu trúc nhận dữ liệu từ HTTP request
// CHỈ CẦN: department_id (nếu user thuộc dept) HOẶC branch_id (nếu không thuộc dept)
type createReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	// Các ID sau là OPTIONAL - hệ thống tự động lấy từ department/branch
	ShopID       *string `json:"shop_id"`       // Optional - auto-fetched
	RegionID     *string `json:"region_id"`     // Optional - auto-fetched
	BranchID     *string `json:"branch_id"`     // Optional nếu có department_id
	DepartmentID *string `json:"department_id"` // Optional - nhưng phải có branch_id HOẶC department_id
}

// validate kiểm tra dữ liệu đầu vào
func (r createReq) validate() error {
	// Kiểm tra phải có ít nhất department_id HOẶC branch_id
	if r.DepartmentID == nil && r.BranchID == nil {
		return errInvalidDeptID // Reuse error (hoặc tạo error mới)
	}

	// Kiểm tra DepartmentID nếu có
	if r.DepartmentID != nil {
		if _, err := primitive.ObjectIDFromHex(*r.DepartmentID); err != nil {
			return errInvalidDeptID
		}
	}

	// Kiểm tra BranchID nếu có
	if r.BranchID != nil {
		if _, err := primitive.ObjectIDFromHex(*r.BranchID); err != nil {
			return errInvalidBranchID
		}
	}

	// Kiểm tra ShopID nếu có (backward compatibility)
	if r.ShopID != nil {
		if _, err := primitive.ObjectIDFromHex(*r.ShopID); err != nil {
			return errInvalidShopID
		}
	}

	// Kiểm tra RegionID nếu có (backward compatibility)
	if r.RegionID != nil {
		if _, err := primitive.ObjectIDFromHex(*r.RegionID); err != nil {
			return errInvalidRegionID
		}
	}

	// Kiểm tra các fields không rỗng
	if strings.TrimSpace(r.Username) == "" || strings.TrimSpace(r.Password) == "" || strings.TrimSpace(r.Email) == "" {
		return errWrongBody
	}
	return nil
}

// toInput chuyển đổi request thành input cho usecase
func (r createReq) toInput() user.CreateInput {
	input := user.CreateInput{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}

	// Parse DepartmentID nếu có
	if r.DepartmentID != nil {
		id, _ := primitive.ObjectIDFromHex(*r.DepartmentID)
		input.DepartmentID = &id
	}

	// Parse BranchID nếu có
	if r.BranchID != nil {
		id, _ := primitive.ObjectIDFromHex(*r.BranchID)
		input.BranchID = id
	}

	// Parse ShopID nếu có (backward compatibility)
	if r.ShopID != nil {
		id, _ := primitive.ObjectIDFromHex(*r.ShopID)
		input.ShopID = id
	}

	// Parse RegionID nếu có (backward compatibility)
	if r.RegionID != nil {
		id, _ := primitive.ObjectIDFromHex(*r.RegionID)
		input.RegionID = id
	}

	return input
}

// updateReq là cấu trúc nhận dữ liệu update từ HTTP request
type updateReq struct {
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	Email        *string `json:"email"`
	ShopID       *string `json:"shop_id"`
	RegionID     *string `json:"region_id"`
	BranchID     *string `json:"branch_id"`
	DepartmentID *string `json:"department_id"`
}

// validate kiểm tra dữ liệu update
func (r updateReq) validate() error {
	// Kiểm tra ShopID nếu có
	if r.ShopID != nil {
		if _, err := primitive.ObjectIDFromHex(*r.ShopID); err != nil {
			return errInvalidShopID
		}
	}
	// Kiểm tra RegionID nếu có
	if r.RegionID != nil {
		if _, err := primitive.ObjectIDFromHex(*r.RegionID); err != nil {
			return errInvalidRegionID
		}
	}
	// Kiểm tra BranchID nếu có
	if r.BranchID != nil {
		if _, err := primitive.ObjectIDFromHex(*r.BranchID); err != nil {
			return errInvalidBranchID
		}
	}
	// Kiểm tra DepartmentID nếu có
	if r.DepartmentID != nil {
		if _, err := primitive.ObjectIDFromHex(*r.DepartmentID); err != nil {
			return errInvalidDeptID
		}
	}
	return nil
}

// toInput chuyển đổi update request thành input cho usecase
func (r updateReq) toInput(id primitive.ObjectID) user.UpdateInput {
	input := user.UpdateInput{
		ID:       id,
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}

	if r.ShopID != nil {
		shopID, _ := primitive.ObjectIDFromHex(*r.ShopID)
		input.ShopID = &shopID
	}
	if r.RegionID != nil {
		regionID, _ := primitive.ObjectIDFromHex(*r.RegionID)
		input.RegionID = &regionID
	}
	if r.BranchID != nil {
		branchID, _ := primitive.ObjectIDFromHex(*r.BranchID)
		input.BranchID = &branchID
	}
	if r.DepartmentID != nil {
		deptID, _ := primitive.ObjectIDFromHex(*r.DepartmentID)
		input.DepartmentID = &deptID
	}

	return input
}

// detailResp là cấu trúc trả về cho client
type detailResp struct {
	ID           string  `json:"id"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	ShopID       string  `json:"shop_id,omitempty"`
	RegionID     string  `json:"region_id,omitempty"`
	BranchID     string  `json:"branch_id,omitempty"`
	DepartmentID *string `json:"department_id,omitempty"`
}

// newDetailResp tạo response từ user model
func (h handler) newDetailResp(d models.User) detailResp {
	resp := detailResp{
		ID:       d.ID.Hex(),
		Username: d.Username,
		Email:    d.Email,
	}

	// Chỉ thêm các field nếu có giá trị (không phải zero value)
	if d.ShopID != primitive.NilObjectID {
		resp.ShopID = d.ShopID.Hex()
	}
	if d.RegionID != primitive.NilObjectID {
		resp.RegionID = d.RegionID.Hex()
	}
	if d.BranchID != primitive.NilObjectID {
		resp.BranchID = d.BranchID.Hex()
	}
	if d.DepartmentID != nil && *d.DepartmentID != primitive.NilObjectID {
		deptID := d.DepartmentID.Hex()
		resp.DepartmentID = &deptID
	}

	return resp
}
