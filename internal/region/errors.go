package region

import "errors"

var (
	// ErrRegionInUse trả về khi region đang được sử dụng bởi branch
	ErrRegionInUse = errors.New("region is being used by branches, cannot delete")
)
