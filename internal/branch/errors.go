package branch

import "errors"

var (
	// ErrBranchInUse trả về khi branch đang được sử dụng bởi department
	ErrBranchInUse = errors.New("branch is being used by departments, cannot delete")
)
