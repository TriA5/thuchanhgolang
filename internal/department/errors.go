package department

import "errors"

var (
	// ErrDepartmentInUse trả về khi department đang được sử dụng bởi users
	ErrDepartmentInUse = errors.New("department is being used by users, cannot delete")
)
