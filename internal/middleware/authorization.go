package middleware

import (
	"thuchanhgolang/internal/models"
	"thuchanhgolang/pkg/jwt"
	"thuchanhgolang/pkg/response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RequireRole middleware kiểm tra user có role yêu cầu không
func (mw *implMiddleware) RequireRole(allowedRoles ...models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, ok := jwt.GetPayloadFromContext(c.Request.Context())
		if !ok {
			mw.l.Warnf(c.Request.Context(), "middleware.RequireRole: payload not found")
			response.Unauthorized(c)
			c.Abort()
			return
		}

		userRole := models.Role(payload.Role)

		// Kiểm tra role có trong danh sách allowed không
		allowed := false
		for _, role := range allowedRoles {
			if userRole == role {
				allowed = true
				break
			}
		}

		if !allowed {
			mw.l.Warnf(c.Request.Context(), "middleware.RequireRole: user role %s not allowed", userRole)
			response.Forbidden(c)
			c.Abort()
			return
		}

		c.Next()
	}
}

// CheckShopAccess kiểm tra user có quyền truy cập Shop không
func (mw *implMiddleware) CheckShopAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, _ := jwt.GetPayloadFromContext(c.Request.Context())
		userRole := models.Role(payload.Role)

		// Manager có quyền CRUD tất cả trong Shop
		if userRole != models.RoleManager {
			mw.l.Warnf(c.Request.Context(), "middleware.CheckShopAccess: user role %s not allowed", userRole)
			response.Forbidden(c)
			c.Abort()
			return
		}

		// Kiểm tra ShopID từ payload
		shopID := c.Param("id")
		if shopID != "" && payload.ShopID != shopID {
			mw.l.Warnf(c.Request.Context(), "middleware.CheckShopAccess: shop_id mismatch")
			response.Forbidden(c)
			c.Abort()
			return
		}

		c.Next()
	}
}

// CheckRegionAccess kiểm tra user có quyền truy cập Region không
func (mw *implMiddleware) CheckRegionAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, _ := jwt.GetPayloadFromContext(c.Request.Context())
		userRole := models.Role(payload.Role)

		// Manager hoặc RegionManager có quyền
		if userRole != models.RoleManager && userRole != models.RoleRegionManager {
			mw.l.Warnf(c.Request.Context(), "middleware.CheckRegionAccess: user role %s not allowed", userRole)
			response.Forbidden(c)
			c.Abort()
			return
		}

		// RegionManager chỉ được truy cập region của mình
		if userRole == models.RoleRegionManager {
			regionID := c.Param("id")
			if regionID != "" && payload.RegionID != regionID {
				mw.l.Warnf(c.Request.Context(), "middleware.CheckRegionAccess: region_id mismatch")
				response.Forbidden(c)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// CheckBranchAccess kiểm tra user có quyền truy cập Branch không
func (mw *implMiddleware) CheckBranchAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, _ := jwt.GetPayloadFromContext(c.Request.Context())
		userRole := models.Role(payload.Role)

		// Manager, RegionManager, BranchManager có quyền
		allowedRoles := []models.Role{models.RoleManager, models.RoleRegionManager, models.RoleBranchManager}
		allowed := false
		for _, role := range allowedRoles {
			if userRole == role {
				allowed = true
				break
			}
		}

		if !allowed {
			mw.l.Warnf(c.Request.Context(), "middleware.CheckBranchAccess: user role %s not allowed", userRole)
			response.Forbidden(c)
			c.Abort()
			return
		}

		// BranchManager chỉ được truy cập branch của mình
		if userRole == models.RoleBranchManager {
			branchID := c.Param("id")
			if branchID != "" && payload.BranchID != branchID {
				mw.l.Warnf(c.Request.Context(), "middleware.CheckBranchAccess: branch_id mismatch")
				response.Forbidden(c)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// CheckDepartmentAccess kiểm tra user có quyền truy cập Department không
func (mw *implMiddleware) CheckDepartmentAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, _ := jwt.GetPayloadFromContext(c.Request.Context())
		userRole := models.Role(payload.Role)

		// Manager, RegionManager, BranchManager, HeadOfDepartment có quyền
		allowedRoles := []models.Role{
			models.RoleManager,
			models.RoleRegionManager,
			models.RoleBranchManager,
			models.RoleHeadOfDepartment,
		}
		allowed := false
		for _, role := range allowedRoles {
			if userRole == role {
				allowed = true
				break
			}
		}

		if !allowed {
			mw.l.Warnf(c.Request.Context(), "middleware.CheckDepartmentAccess: user role %s not allowed", userRole)
			response.Forbidden(c)
			c.Abort()
			return
		}

		c.Next()
	}
}

// CheckUserAccess kiểm tra quyền truy cập User
// Employee chỉ xem được users trong cùng branch
func (mw *implMiddleware) CheckUserAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, _ := jwt.GetPayloadFromContext(c.Request.Context())
		userRole := models.Role(payload.Role)

		// Employee chỉ được xem users trong cùng branch (không được CRUD shop/region/branch/dept)
		if userRole == models.RoleEmployee {
			// Employee chỉ được GET users, không được POST/PUT/DELETE
			if c.Request.Method != "GET" {
				mw.l.Warnf(c.Request.Context(), "middleware.CheckUserAccess: employee cannot modify users")
				response.Forbidden(c)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// SetScopeFromPayload tạo scope từ JWT payload và set vào context
func (mw *implMiddleware) SetScopeFromPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, ok := jwt.GetPayloadFromContext(c.Request.Context())
		if !ok {
			c.Next()
			return
		}

		scope := models.Scope{
			UserID: payload.UserID,
			Role:   models.Role(payload.Role),
		}

		// Parse IDs từ payload
		if payload.ShopID != "" {
			if shopID, err := primitive.ObjectIDFromHex(payload.ShopID); err == nil {
				scope.ShopID = &shopID
			}
		}
		if payload.RegionID != "" {
			if regionID, err := primitive.ObjectIDFromHex(payload.RegionID); err == nil {
				scope.RegionID = &regionID
			}
		}
		if payload.BranchID != "" {
			if branchID, err := primitive.ObjectIDFromHex(payload.BranchID); err == nil {
				scope.BranchID = &branchID
			}
		}

		c.Set("scope", scope)
		c.Next()
	}
}
