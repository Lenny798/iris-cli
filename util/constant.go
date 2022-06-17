package util

// 定义错误编码
const (
	RecodeOk        = "0"
	RecodeDbErr     = "4001"
	RecodeLoginErr  = "4002"
	RecodeParamErr  = "4003"
	RecodeSysErr    = "4004"
	RecodeEthErr    = "4005"
	RecodeUnknowErr = "4006"
)

var recodeText = map[string]string{
	RecodeOk:        "成功",
	RecodeDbErr:     "数据库操作错误",
	RecodeLoginErr:  "用户登录失败",
	RecodeParamErr:  "参数错误",
	RecodeSysErr:    "系统错误",
	RecodeEthErr:    "与以太坊交互失败",
	RecodeUnknowErr: "未知错误",
}

// RecodeText 通过编码获取对应信息
func RecodeText(code string) string {
	str, ok := recodeText[code]
	if !ok {
		return recodeText[RecodeUnknowErr]
	}
	return str
}
