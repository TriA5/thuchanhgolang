package shop

import "errors"

var (
	// ErrShopInUse trả về khi shop đang được sử dụng bởi region
	ErrShopInUse = errors.New("shop is being used by regions, cannot delete")
)
