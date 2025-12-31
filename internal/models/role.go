package models

// Role định nghĩa các vai trò trong hệ thống
type Role string

const (
	// RoleManager có quyền CRUD tất cả thông tin trong một Shop
	RoleManager Role = "manager"

	// RoleRegionManager có quyền CRUD tất cả thông tin trong một Region
	RoleRegionManager Role = "region_manager"

	// RoleBranchManager có quyền CRUD tất cả thông tin trong một Branch
	RoleBranchManager Role = "branch_manager"

	// RoleHeadOfDepartment có quyền CRUD tất cả thông tin trong một Department
	RoleHeadOfDepartment Role = "head_of_department"

	// RoleEmployee chỉ thấy nhân viên trong cùng chi nhánh
	RoleEmployee Role = "employee"
)

// IsValid kiểm tra role có hợp lệ không
func (r Role) IsValid() bool {
	switch r {
	case RoleManager, RoleRegionManager, RoleBranchManager, RoleHeadOfDepartment, RoleEmployee:
		return true
	}
	return false
}

// String returns the string representation of the role
func (r Role) String() string {
	return string(r)
}
