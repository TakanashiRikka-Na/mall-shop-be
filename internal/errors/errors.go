package errors

import (
	"github.com/go-kratos/kratos/v2/errors"
	v1 "mall-shop-be/api/helloworld/v1"
	"time"
)

type ErrorController interface {
	UnKnownError() *errors.Error
	UserExist() *errors.Error
	UserNotFound() *errors.Error
	UserNameValid() *errors.Error
	PasswdValid() *errors.Error
	PasswdError() *errors.Error
}

func NewErrorController() ErrorController {
	return &ERR{}
}

type ERR struct{}

func (*ERR) UserNameValid() *errors.Error {
	return UserNameVaild
}

func (*ERR) PasswdError() *errors.Error {
	return PasswdError
}

var (
	UserNotFound  = errors.BadRequest(v1.ErrorReason_USER_NOT_FOUND.String(), "未找到该用户")
	UserExisted   = errors.BadRequest(v1.ErrorReason_USER_EXISTED.String(), "用户已经存在")
	PasswdValid   = errors.BadRequest(v1.ErrorReason_PASSWD_NOT_VALID.String(), "密码不符合规范")
	PasswdError   = errors.BadRequest(v1.ErrorReason_PASSWD_ERROR.String(), "密码错误")
	UserNameVaild = errors.BadRequest(v1.ErrorReason_USERNAME_NOT_VALID.String(), "用户名不是纯数字")
)

func (*ERR) UnKnownError() *errors.Error {
	return errors.InternalServer(v1.ErrorReason_UNKNOWN_ERROR.String(), "未知错误，时间: "+time.Now().Format("2006-01-02 15:04:05"))
}

func (*ERR) UserExist() *errors.Error {
	return UserExisted
}

func (*ERR) UserNotFound() *errors.Error {
	return UserNotFound
}

func (*ERR) PasswdValid() *errors.Error {
	return PasswdValid
}
