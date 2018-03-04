package utils

// retCodeDict return code directory
var retCodeDict = map[string]string{
	"0000": "成功",
	"1000": "系统错误",
	"2000": "Ajax Autocomplete 错误",
	"2001": "Ajax Data 应答错误",
}

// GetRetMsg func(retcode string) string
// Get Return Message By Return Code
func GetRetMsg(retcode string) string {
	return retCodeDict[retcode]
}
