package aerror

const ErrCodeUserAccountNotExists = "20001"
const ErrCodeUserAccountExists = "20002"
const ErrCodeUserPasswordIncorrect = "20003"

var UserError = map[string]string{
	ErrCodeUserAccountExists:     "该账号已被注册",
	ErrCodeUserAccountNotExists:  "账号不存在",
	ErrCodeUserPasswordIncorrect: "密码错误",
}
