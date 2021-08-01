package baseerror

var (
	ErrSystem       = NewCodeError(10001, "服务器错误")
	ErrInvalidParam = NewCodeError(10002, "参数错误")
	ErrUserNotLogin = NewCodeError(10003, "用户未登录")
	ErrAccess       = NewCodeError(10004, "权限不足")

	ErrDefault = ErrSystem
)

type CodeError struct {
	code int64
	desc string
	data interface{}
}

func (err *CodeError) Error() string {
	return err.desc
}

func (err *CodeError) Code() int64 {
	return err.code
}

func (err *CodeError) Desc() string {
	return err.desc
}

func (err *CodeError) Data() interface{} {
	return err.data
}

func NewCodeError(code int64, desc string) *CodeError {
	return &CodeError{
		code: code,
		desc: desc,
	}
}

func NewError(code int64, desc string, data interface{}) *CodeError {
	return &CodeError{
		code: code,
		desc: desc,
		data: data,
	}
}

func IsCodeError(err error) bool {
	switch err.(type) {
	case *CodeError:
		return true
	}
	return false
}

func FromError(err error) (codeErr *CodeError, ok bool) {
	if se, ok := err.(*CodeError); ok {
		return se, true
	}
	return nil, false
}

func ToCodeError(err error) *CodeError {
	if IsCodeError(err) {
		return err.(*CodeError)
	}
	return ErrDefault
}
