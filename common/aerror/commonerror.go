package aerror

const ErrCodeSystemSuccess = "0"
const ErrCodeSystemError = "10001"
const ErrCodeParamError = "10002"

var CommonError = map[string]string{
	ErrCodeSystemSuccess: "处理成功",
	ErrCodeSystemError:   "系统错误",
	ErrCodeParamError:    "参数错误",
}
