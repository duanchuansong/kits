package errcode

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// 自定义错误类型
var (
	Succ                = errors.New(1, "success", "成功")
	ErrUserNotFound     = errors.New(404, "user", "用户不存在")
	ErrPermissionDenied = errors.New(403, "permission", "权限不足")
	// 可以继续添加更多自定义错误...
)

// ToError 转换为 errors.Error 类型
func ToError(errCode int, errMsg string) error {
	return errors.New(errCode, "app", errMsg)
}
