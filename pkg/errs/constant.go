package errs

import "fmt"

var (
	ErrAccountIllegal      = fmt.Errorf("账号非法")
	ErrAccountBlack        = fmt.Errorf("账号在黑名单中")
	ErrAccountNonWhitelist = fmt.Errorf("账号不在白名单中")
)
