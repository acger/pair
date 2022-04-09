package aerror

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

type RspErrMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler(err error) (int, interface{}) {
	var code string
	var msg string

	e := status.Convert(err)

	if e.Code() == 2 {
		code = ErrCodeSystemError
		msg = GetErrCodeMessage(code)
	} else {
		code = strconv.Itoa(int(e.Code()))
		msg = e.Message()
	}

	return http.StatusOK, RspErrMessage{
		Code:    code,
		Message: msg,
	}
}

func Err(code string, message ...string) error {
	var codeInt int
	var err error
	var msg string

	if len(message) > 0 {
		for _, m := range message {
			msg = fmt.Sprint(msg, m)
		}
	} else {
		msg = GetErrCodeMessage(code)
	}

	if codeInt, err = strconv.Atoi(code); err != nil {
		codeInt = 10001
	}

	cc := codes.Code(codeInt)
	return status.Error(cc, msg)
}

func ErrLog(err error, arg ...interface{}) error {
	logx.ErrorCaller(2, err)

	for _, a := range arg {
		logx.Error(a)
	}

	logx.Error("-end-")

	return Err(ErrCodeSystemError)
}

func GetErrCodeMessage(code string) string {
	var m map[string]string

	switch code[:1] {
	case "1":
		m = CommonError
	case "2":
		m = UserError
	case "3":
		m = PairError
	case "4":
		m = ChatError
	default:
		m = make(map[string]string)
	}

	if message, ok := m[code]; ok {
		return message
	}

	return CommonError[ErrCodeSystemError]
}
