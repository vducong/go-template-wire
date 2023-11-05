package failure

import "net/http"

type ErrCode int

const (
	ErrReusableCodeGetByCodeBinding ErrCode = 991001
	ErrReusableCodeNotFound         ErrCode = 991002
	ErrReusableCodeFailed           ErrCode = 991003
)

var errMsgMap = map[ErrCode]string{
	ErrReusableCodeNotFound: "Mã quà tặng không tồn tại",
}

var errCodeMap = map[ErrCode]int{
	ErrReusableCodeFailed: http.StatusInternalServerError,
}
